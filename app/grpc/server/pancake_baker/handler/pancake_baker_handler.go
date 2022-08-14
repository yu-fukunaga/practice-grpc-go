package handler

import (
	"context"
	"math/rand"
	"sync"
	"time"

	"practice/gen/proto"

	"github.com/golang/protobuf/ptypes/timestamp"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

type BakerHandler struct {
	proto.UnimplementedPancakeBakerServiceServer
	report *report
}

type report struct {
	sync.Mutex                            // 複数人が同時に焼いても大丈夫にしておきます
	data       map[proto.Pancake_Menu]int // 焼いた数
}

func NewBakerHandler() *BakerHandler {
	return &BakerHandler{
		report: &report{
			data: make(map[proto.Pancake_Menu]int),
		},
	}
}

func (h *BakerHandler) Bake(ctx context.Context, req *proto.BakeRequest) (*proto.BakeResponse, error) {
	// validation
	if req.Menu == proto.Pancake_UNKNOWN || req.Menu > proto.Pancake_SPICY_CURRY {
		return nil, status.Errorf(codes.InvalidArgument, "Choose your pancakes!")
	}

	// Cook pancakes and keep track of the number
	now := time.Now()
	h.report.Lock()
	h.report.data[req.Menu] = h.report.data[req.Menu] + 1
	h.report.Unlock()

	// response
	return &proto.BakeResponse{
		Pancake: &proto.Pancake{
			Menu:           req.Menu,
			ChefName:       "uyu",
			TechnicalScore: rand.Float32(),
			CreatedAt: &timestamp.Timestamp{
				Seconds: now.Unix(),
				Nanos:   int32(now.Nanosecond()),
			},
		},
	}, nil

}

func (h *BakerHandler) Report(ctx context.Context, req *proto.ReportRequest) (*proto.ReportResponse, error) {

	counts := make([]*proto.Report_BakeCount, 0)

	// create Report
	h.report.Lock()
	for k, v := range h.report.data {
		counts = append(counts, &proto.Report_BakeCount{
			Menu:  k,
			Count: int32(v),
		})
	}
	h.report.Unlock()

	// response
	return &proto.ReportResponse{
		Report: &proto.Report{
			BakeCounts: counts,
		},
	}, nil
}
