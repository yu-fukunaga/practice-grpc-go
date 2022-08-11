// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.19.4
// source: proto/reversi_game.proto

package proto

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type PlayRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RoomId int32   `protobuf:"varint,1,opt,name=room_id,json=roomId,proto3" json:"room_id,omitempty"`
	Player *Player `protobuf:"bytes,2,opt,name=player,proto3" json:"player,omitempty"`
	// Types that are assignable to Action:
	//	*PlayRequest_Start
	Action isPlayRequest_Action `protobuf_oneof:"action"`
}

func (x *PlayRequest) Reset() {
	*x = PlayRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_reversi_game_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PlayRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PlayRequest) ProtoMessage() {}

func (x *PlayRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_reversi_game_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PlayRequest.ProtoReflect.Descriptor instead.
func (*PlayRequest) Descriptor() ([]byte, []int) {
	return file_proto_reversi_game_proto_rawDescGZIP(), []int{0}
}

func (x *PlayRequest) GetRoomId() int32 {
	if x != nil {
		return x.RoomId
	}
	return 0
}

func (x *PlayRequest) GetPlayer() *Player {
	if x != nil {
		return x.Player
	}
	return nil
}

func (m *PlayRequest) GetAction() isPlayRequest_Action {
	if m != nil {
		return m.Action
	}
	return nil
}

func (x *PlayRequest) GetStart() *PlayRequest_StartAction {
	if x, ok := x.GetAction().(*PlayRequest_Start); ok {
		return x.Start
	}
	return nil
}

type isPlayRequest_Action interface {
	isPlayRequest_Action()
}

type PlayRequest_Start struct {
	Start *PlayRequest_StartAction `protobuf:"bytes,3,opt,name=start,proto3,oneof"`
}

func (*PlayRequest_Start) isPlayRequest_Action() {}

type PlayResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Event:
	//	*PlayResponse_Waiting
	Event isPlayResponse_Event `protobuf_oneof:"event"`
}

func (x *PlayResponse) Reset() {
	*x = PlayResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_reversi_game_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PlayResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PlayResponse) ProtoMessage() {}

func (x *PlayResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_reversi_game_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PlayResponse.ProtoReflect.Descriptor instead.
func (*PlayResponse) Descriptor() ([]byte, []int) {
	return file_proto_reversi_game_proto_rawDescGZIP(), []int{1}
}

func (m *PlayResponse) GetEvent() isPlayResponse_Event {
	if m != nil {
		return m.Event
	}
	return nil
}

func (x *PlayResponse) GetWaiting() *PlayResponse_WaitingEvent {
	if x, ok := x.GetEvent().(*PlayResponse_Waiting); ok {
		return x.Waiting
	}
	return nil
}

type isPlayResponse_Event interface {
	isPlayResponse_Event()
}

type PlayResponse_Waiting struct {
	Waiting *PlayResponse_WaitingEvent `protobuf:"bytes,1,opt,name=waiting,proto3,oneof"`
}

func (*PlayResponse_Waiting) isPlayResponse_Event() {}

type Move struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	X int32 `protobuf:"varint,1,opt,name=x,proto3" json:"x,omitempty"`
	Y int32 `protobuf:"varint,2,opt,name=y,proto3" json:"y,omitempty"`
}

func (x *Move) Reset() {
	*x = Move{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_reversi_game_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Move) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Move) ProtoMessage() {}

func (x *Move) ProtoReflect() protoreflect.Message {
	mi := &file_proto_reversi_game_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Move.ProtoReflect.Descriptor instead.
func (*Move) Descriptor() ([]byte, []int) {
	return file_proto_reversi_game_proto_rawDescGZIP(), []int{2}
}

func (x *Move) GetX() int32 {
	if x != nil {
		return x.X
	}
	return 0
}

func (x *Move) GetY() int32 {
	if x != nil {
		return x.Y
	}
	return 0
}

type Board struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Cols []*Board_Col `protobuf:"bytes,1,rep,name=cols,proto3" json:"cols,omitempty"`
}

func (x *Board) Reset() {
	*x = Board{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_reversi_game_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Board) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Board) ProtoMessage() {}

func (x *Board) ProtoReflect() protoreflect.Message {
	mi := &file_proto_reversi_game_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Board.ProtoReflect.Descriptor instead.
func (*Board) Descriptor() ([]byte, []int) {
	return file_proto_reversi_game_proto_rawDescGZIP(), []int{3}
}

func (x *Board) GetCols() []*Board_Col {
	if x != nil {
		return x.Cols
	}
	return nil
}

type PlayRequest_StartAction struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *PlayRequest_StartAction) Reset() {
	*x = PlayRequest_StartAction{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_reversi_game_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PlayRequest_StartAction) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PlayRequest_StartAction) ProtoMessage() {}

func (x *PlayRequest_StartAction) ProtoReflect() protoreflect.Message {
	mi := &file_proto_reversi_game_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PlayRequest_StartAction.ProtoReflect.Descriptor instead.
func (*PlayRequest_StartAction) Descriptor() ([]byte, []int) {
	return file_proto_reversi_game_proto_rawDescGZIP(), []int{0, 0}
}

type PlayRequest_MoveAction struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Move *Move `protobuf:"bytes,1,opt,name=move,proto3" json:"move,omitempty"`
}

func (x *PlayRequest_MoveAction) Reset() {
	*x = PlayRequest_MoveAction{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_reversi_game_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PlayRequest_MoveAction) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PlayRequest_MoveAction) ProtoMessage() {}

func (x *PlayRequest_MoveAction) ProtoReflect() protoreflect.Message {
	mi := &file_proto_reversi_game_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PlayRequest_MoveAction.ProtoReflect.Descriptor instead.
func (*PlayRequest_MoveAction) Descriptor() ([]byte, []int) {
	return file_proto_reversi_game_proto_rawDescGZIP(), []int{0, 1}
}

func (x *PlayRequest_MoveAction) GetMove() *Move {
	if x != nil {
		return x.Move
	}
	return nil
}

type PlayResponse_WaitingEvent struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *PlayResponse_WaitingEvent) Reset() {
	*x = PlayResponse_WaitingEvent{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_reversi_game_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PlayResponse_WaitingEvent) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PlayResponse_WaitingEvent) ProtoMessage() {}

func (x *PlayResponse_WaitingEvent) ProtoReflect() protoreflect.Message {
	mi := &file_proto_reversi_game_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PlayResponse_WaitingEvent.ProtoReflect.Descriptor instead.
func (*PlayResponse_WaitingEvent) Descriptor() ([]byte, []int) {
	return file_proto_reversi_game_proto_rawDescGZIP(), []int{1, 0}
}

type PlayResponse_ReadyEvent struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *PlayResponse_ReadyEvent) Reset() {
	*x = PlayResponse_ReadyEvent{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_reversi_game_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PlayResponse_ReadyEvent) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PlayResponse_ReadyEvent) ProtoMessage() {}

func (x *PlayResponse_ReadyEvent) ProtoReflect() protoreflect.Message {
	mi := &file_proto_reversi_game_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PlayResponse_ReadyEvent.ProtoReflect.Descriptor instead.
func (*PlayResponse_ReadyEvent) Descriptor() ([]byte, []int) {
	return file_proto_reversi_game_proto_rawDescGZIP(), []int{1, 1}
}

type PlayResponse_MoveEvent struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Player *Player `protobuf:"bytes,1,opt,name=player,proto3" json:"player,omitempty"`
	Move   *Move   `protobuf:"bytes,2,opt,name=move,proto3" json:"move,omitempty"`
	Board  *Board  `protobuf:"bytes,3,opt,name=board,proto3" json:"board,omitempty"`
}

func (x *PlayResponse_MoveEvent) Reset() {
	*x = PlayResponse_MoveEvent{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_reversi_game_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PlayResponse_MoveEvent) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PlayResponse_MoveEvent) ProtoMessage() {}

func (x *PlayResponse_MoveEvent) ProtoReflect() protoreflect.Message {
	mi := &file_proto_reversi_game_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PlayResponse_MoveEvent.ProtoReflect.Descriptor instead.
func (*PlayResponse_MoveEvent) Descriptor() ([]byte, []int) {
	return file_proto_reversi_game_proto_rawDescGZIP(), []int{1, 2}
}

func (x *PlayResponse_MoveEvent) GetPlayer() *Player {
	if x != nil {
		return x.Player
	}
	return nil
}

func (x *PlayResponse_MoveEvent) GetMove() *Move {
	if x != nil {
		return x.Move
	}
	return nil
}

func (x *PlayResponse_MoveEvent) GetBoard() *Board {
	if x != nil {
		return x.Board
	}
	return nil
}

type PlayResponse_FinishedEvent struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Winner Color  `protobuf:"varint,1,opt,name=winner,proto3,enum=game.Color" json:"winner,omitempty"`
	Board  *Board `protobuf:"bytes,2,opt,name=board,proto3" json:"board,omitempty"`
}

func (x *PlayResponse_FinishedEvent) Reset() {
	*x = PlayResponse_FinishedEvent{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_reversi_game_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PlayResponse_FinishedEvent) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PlayResponse_FinishedEvent) ProtoMessage() {}

func (x *PlayResponse_FinishedEvent) ProtoReflect() protoreflect.Message {
	mi := &file_proto_reversi_game_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PlayResponse_FinishedEvent.ProtoReflect.Descriptor instead.
func (*PlayResponse_FinishedEvent) Descriptor() ([]byte, []int) {
	return file_proto_reversi_game_proto_rawDescGZIP(), []int{1, 3}
}

func (x *PlayResponse_FinishedEvent) GetWinner() Color {
	if x != nil {
		return x.Winner
	}
	return Color_UNKNOWN
}

func (x *PlayResponse_FinishedEvent) GetBoard() *Board {
	if x != nil {
		return x.Board
	}
	return nil
}

type Board_Col struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Cells []Color `protobuf:"varint,1,rep,packed,name=cells,proto3,enum=game.Color" json:"cells,omitempty"`
}

func (x *Board_Col) Reset() {
	*x = Board_Col{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_reversi_game_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Board_Col) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Board_Col) ProtoMessage() {}

func (x *Board_Col) ProtoReflect() protoreflect.Message {
	mi := &file_proto_reversi_game_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Board_Col.ProtoReflect.Descriptor instead.
func (*Board_Col) Descriptor() ([]byte, []int) {
	return file_proto_reversi_game_proto_rawDescGZIP(), []int{3, 0}
}

func (x *Board_Col) GetCells() []Color {
	if x != nil {
		return x.Cells
	}
	return nil
}

var File_proto_reversi_game_proto protoreflect.FileDescriptor

var file_proto_reversi_game_proto_rawDesc = []byte{
	0x0a, 0x18, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x72, 0x65, 0x76, 0x65, 0x72, 0x73, 0x69, 0x5f,
	0x67, 0x61, 0x6d, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x04, 0x67, 0x61, 0x6d, 0x65,
	0x1a, 0x1a, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x72, 0x65, 0x76, 0x65, 0x72, 0x73, 0x69, 0x5f,
	0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x19, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2f, 0x72, 0x65, 0x76, 0x65, 0x72, 0x73, 0x69, 0x5f, 0x63, 0x6f, 0x6c, 0x6f,
	0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xca, 0x01, 0x0a, 0x0b, 0x50, 0x6c, 0x61, 0x79,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x72, 0x6f, 0x6f, 0x6d, 0x5f,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x72, 0x6f, 0x6f, 0x6d, 0x49, 0x64,
	0x12, 0x24, 0x0a, 0x06, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x0c, 0x2e, 0x67, 0x61, 0x6d, 0x65, 0x2e, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x52, 0x06,
	0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x12, 0x35, 0x0a, 0x05, 0x73, 0x74, 0x61, 0x72, 0x74, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x67, 0x61, 0x6d, 0x65, 0x2e, 0x50, 0x6c, 0x61,
	0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x53, 0x74, 0x61, 0x72, 0x74, 0x41, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x48, 0x00, 0x52, 0x05, 0x73, 0x74, 0x61, 0x72, 0x74, 0x1a, 0x0d, 0x0a,
	0x0b, 0x53, 0x74, 0x61, 0x72, 0x74, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x1a, 0x2c, 0x0a, 0x0a,
	0x4d, 0x6f, 0x76, 0x65, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1e, 0x0a, 0x04, 0x6d, 0x6f,
	0x76, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0a, 0x2e, 0x67, 0x61, 0x6d, 0x65, 0x2e,
	0x4d, 0x6f, 0x76, 0x65, 0x52, 0x04, 0x6d, 0x6f, 0x76, 0x65, 0x42, 0x08, 0x0a, 0x06, 0x61, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x22, 0xc1, 0x02, 0x0a, 0x0c, 0x50, 0x6c, 0x61, 0x79, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3b, 0x0a, 0x07, 0x77, 0x61, 0x69, 0x74, 0x69, 0x6e, 0x67,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x67, 0x61, 0x6d, 0x65, 0x2e, 0x50, 0x6c,
	0x61, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x57, 0x61, 0x69, 0x74, 0x69,
	0x6e, 0x67, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x48, 0x00, 0x52, 0x07, 0x77, 0x61, 0x69, 0x74, 0x69,
	0x6e, 0x67, 0x1a, 0x0e, 0x0a, 0x0c, 0x57, 0x61, 0x69, 0x74, 0x69, 0x6e, 0x67, 0x45, 0x76, 0x65,
	0x6e, 0x74, 0x1a, 0x0c, 0x0a, 0x0a, 0x52, 0x65, 0x61, 0x64, 0x79, 0x45, 0x76, 0x65, 0x6e, 0x74,
	0x1a, 0x74, 0x0a, 0x09, 0x4d, 0x6f, 0x76, 0x65, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x12, 0x24, 0x0a,
	0x06, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e,
	0x67, 0x61, 0x6d, 0x65, 0x2e, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x52, 0x06, 0x70, 0x6c, 0x61,
	0x79, 0x65, 0x72, 0x12, 0x1e, 0x0a, 0x04, 0x6d, 0x6f, 0x76, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x0a, 0x2e, 0x67, 0x61, 0x6d, 0x65, 0x2e, 0x4d, 0x6f, 0x76, 0x65, 0x52, 0x04, 0x6d,
	0x6f, 0x76, 0x65, 0x12, 0x21, 0x0a, 0x05, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x0b, 0x2e, 0x67, 0x61, 0x6d, 0x65, 0x2e, 0x42, 0x6f, 0x61, 0x72, 0x64, 0x52,
	0x05, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x1a, 0x57, 0x0a, 0x0d, 0x46, 0x69, 0x6e, 0x69, 0x73, 0x68,
	0x65, 0x64, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x12, 0x23, 0x0a, 0x06, 0x77, 0x69, 0x6e, 0x6e, 0x65,
	0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x0b, 0x2e, 0x67, 0x61, 0x6d, 0x65, 0x2e, 0x43,
	0x6f, 0x6c, 0x6f, 0x72, 0x52, 0x06, 0x77, 0x69, 0x6e, 0x6e, 0x65, 0x72, 0x12, 0x21, 0x0a, 0x05,
	0x62, 0x6f, 0x61, 0x72, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0b, 0x2e, 0x67, 0x61,
	0x6d, 0x65, 0x2e, 0x42, 0x6f, 0x61, 0x72, 0x64, 0x52, 0x05, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x42,
	0x07, 0x0a, 0x05, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x22, 0x22, 0x0a, 0x04, 0x4d, 0x6f, 0x76, 0x65,
	0x12, 0x0c, 0x0a, 0x01, 0x78, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x01, 0x78, 0x12, 0x0c,
	0x0a, 0x01, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x01, 0x79, 0x22, 0x56, 0x0a, 0x05,
	0x42, 0x6f, 0x61, 0x72, 0x64, 0x12, 0x23, 0x0a, 0x04, 0x63, 0x6f, 0x6c, 0x73, 0x18, 0x01, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x67, 0x61, 0x6d, 0x65, 0x2e, 0x42, 0x6f, 0x61, 0x72, 0x64,
	0x2e, 0x43, 0x6f, 0x6c, 0x52, 0x04, 0x63, 0x6f, 0x6c, 0x73, 0x1a, 0x28, 0x0a, 0x03, 0x43, 0x6f,
	0x6c, 0x12, 0x21, 0x0a, 0x05, 0x63, 0x65, 0x6c, 0x6c, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0e,
	0x32, 0x0b, 0x2e, 0x67, 0x61, 0x6d, 0x65, 0x2e, 0x43, 0x6f, 0x6c, 0x6f, 0x72, 0x52, 0x05, 0x63,
	0x65, 0x6c, 0x6c, 0x73, 0x32, 0x42, 0x0a, 0x0b, 0x47, 0x61, 0x6d, 0x65, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x12, 0x33, 0x0a, 0x04, 0x50, 0x6c, 0x61, 0x79, 0x12, 0x11, 0x2e, 0x67, 0x61,
	0x6d, 0x65, 0x2e, 0x50, 0x6c, 0x61, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x12,
	0x2e, 0x67, 0x61, 0x6d, 0x65, 0x2e, 0x50, 0x6c, 0x61, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x00, 0x28, 0x01, 0x30, 0x01, 0x42, 0x0b, 0x5a, 0x09, 0x67, 0x65, 0x6e, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_reversi_game_proto_rawDescOnce sync.Once
	file_proto_reversi_game_proto_rawDescData = file_proto_reversi_game_proto_rawDesc
)

func file_proto_reversi_game_proto_rawDescGZIP() []byte {
	file_proto_reversi_game_proto_rawDescOnce.Do(func() {
		file_proto_reversi_game_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_reversi_game_proto_rawDescData)
	})
	return file_proto_reversi_game_proto_rawDescData
}

var file_proto_reversi_game_proto_msgTypes = make([]protoimpl.MessageInfo, 11)
var file_proto_reversi_game_proto_goTypes = []interface{}{
	(*PlayRequest)(nil),                // 0: game.PlayRequest
	(*PlayResponse)(nil),               // 1: game.PlayResponse
	(*Move)(nil),                       // 2: game.Move
	(*Board)(nil),                      // 3: game.Board
	(*PlayRequest_StartAction)(nil),    // 4: game.PlayRequest.StartAction
	(*PlayRequest_MoveAction)(nil),     // 5: game.PlayRequest.MoveAction
	(*PlayResponse_WaitingEvent)(nil),  // 6: game.PlayResponse.WaitingEvent
	(*PlayResponse_ReadyEvent)(nil),    // 7: game.PlayResponse.ReadyEvent
	(*PlayResponse_MoveEvent)(nil),     // 8: game.PlayResponse.MoveEvent
	(*PlayResponse_FinishedEvent)(nil), // 9: game.PlayResponse.FinishedEvent
	(*Board_Col)(nil),                  // 10: game.Board.Col
	(*Player)(nil),                     // 11: game.Player
	(Color)(0),                         // 12: game.Color
}
var file_proto_reversi_game_proto_depIdxs = []int32{
	11, // 0: game.PlayRequest.player:type_name -> game.Player
	4,  // 1: game.PlayRequest.start:type_name -> game.PlayRequest.StartAction
	6,  // 2: game.PlayResponse.waiting:type_name -> game.PlayResponse.WaitingEvent
	10, // 3: game.Board.cols:type_name -> game.Board.Col
	2,  // 4: game.PlayRequest.MoveAction.move:type_name -> game.Move
	11, // 5: game.PlayResponse.MoveEvent.player:type_name -> game.Player
	2,  // 6: game.PlayResponse.MoveEvent.move:type_name -> game.Move
	3,  // 7: game.PlayResponse.MoveEvent.board:type_name -> game.Board
	12, // 8: game.PlayResponse.FinishedEvent.winner:type_name -> game.Color
	3,  // 9: game.PlayResponse.FinishedEvent.board:type_name -> game.Board
	12, // 10: game.Board.Col.cells:type_name -> game.Color
	0,  // 11: game.GameService.Play:input_type -> game.PlayRequest
	1,  // 12: game.GameService.Play:output_type -> game.PlayResponse
	12, // [12:13] is the sub-list for method output_type
	11, // [11:12] is the sub-list for method input_type
	11, // [11:11] is the sub-list for extension type_name
	11, // [11:11] is the sub-list for extension extendee
	0,  // [0:11] is the sub-list for field type_name
}

func init() { file_proto_reversi_game_proto_init() }
func file_proto_reversi_game_proto_init() {
	if File_proto_reversi_game_proto != nil {
		return
	}
	file_proto_reversi_player_proto_init()
	file_proto_reversi_color_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_proto_reversi_game_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PlayRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_proto_reversi_game_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PlayResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_proto_reversi_game_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Move); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_proto_reversi_game_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Board); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_proto_reversi_game_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PlayRequest_StartAction); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_proto_reversi_game_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PlayRequest_MoveAction); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_proto_reversi_game_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PlayResponse_WaitingEvent); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_proto_reversi_game_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PlayResponse_ReadyEvent); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_proto_reversi_game_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PlayResponse_MoveEvent); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_proto_reversi_game_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PlayResponse_FinishedEvent); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_proto_reversi_game_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Board_Col); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	file_proto_reversi_game_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*PlayRequest_Start)(nil),
	}
	file_proto_reversi_game_proto_msgTypes[1].OneofWrappers = []interface{}{
		(*PlayResponse_Waiting)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_proto_reversi_game_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   11,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_reversi_game_proto_goTypes,
		DependencyIndexes: file_proto_reversi_game_proto_depIdxs,
		MessageInfos:      file_proto_reversi_game_proto_msgTypes,
	}.Build()
	File_proto_reversi_game_proto = out.File
	file_proto_reversi_game_proto_rawDesc = nil
	file_proto_reversi_game_proto_goTypes = nil
	file_proto_reversi_game_proto_depIdxs = nil
}
