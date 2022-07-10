package handler

import (
	"bytes"
	"io"
	"net/http"
	"sync"

	"practice/gen/proto"

	"github.com/google/uuid"
)

type ImageUploaderHandler struct {
	proto.UnimplementedImageUploadServiceServer
	sync.Mutex
	files map[string][]byte
}

func NewImageUploaderHandler() *ImageUploaderHandler {
	return &ImageUploaderHandler{
		files: make(map[string][]byte),
	}
}

// Image Upload
func (h *ImageUploaderHandler) Upload(stream proto.ImageUploadService_UploadServer) error {
	// 最初のリクエストを受け取る
	req, err := stream.Recv()
	if err != nil {
		return err
	}

	// 初回は必ずメタ情報が送られる
	meta := req.GetFileMeta()
	filename := meta.Filename

	// UUIDの生成
	u, err := uuid.NewRandom()
	if err != nil {
		return err
	}
	uuid := u.String()

	// 画像データ格納用バッファ
	buf := &bytes.Buffer{}

	// 境ごとにアップロードされたバイナリをループしながらすべて受け取る
	for {
		r, err := stream.Recv()
		if err == io.EOF {
			break
		} else if err != nil {
			return err
		}

		chunk := r.GetData()
		_, err = buf.Write(chunk)
		if err != nil {
			return err
		}
	}

	data := buf.Bytes()
	mimeType := http.DetectContentType(data)

	h.files[filename] = data

	err = stream.SendAndClose(&proto.ImageUploadResponse{
		Uuid:        uuid,
		Size:        int32(len(data)),
		Filename:    filename,
		ContentType: mimeType,
	})

	return err
}
