package grpc

import (
	"context"
	"servic_remender/internal/dto"

	reminder "github.com/goggle-source/grpc-proto-reminder/gen"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type serverAPI struct {
	reminder.UnimplementedReminderServer
	rem Remind
}

type Remind interface {
	Create(ctx context.Context, req dto.RequestCreategRPCInServic) (resp dto.ResponseCreategRPCInServic, err error)
	Get(ctx context.Context, req dto.RequestGetgRPCInServic) (resp dto.ResponseGetgRPCInServic, err error)
	Update(ctx context.Context, req dto.RequestUpdateGRPCInServic) (resp dto.ResponseUpdateGRPCInServic, err error)
	Delete(ctx context.Context, req dto.RequestDeletegRRPCInServic) (err error)
}

func CreateServerAPi(gRPCServer *grpc.Server, rem Remind) {
	reminder.RegisterReminderServer(gRPCServer, &serverAPI{rem: rem})
}

func (s *serverAPI) Create(ctx context.Context, in *reminder.CreateRequest) (*reminder.CreateResponse, error) {
	if in.GetName() == "" || in.GetDescription() == "" {
		return nil, status.Error(codes.InvalidArgument, "not have name")
	}

	if in.GetDescription() == "" {
		return nil, status.Error(codes.InvalidArgument, "not have description")
	}

	if !in.GetTimestamp().IsValid() {
		return nil, status.Error(codes.InvalidArgument, "imvalid time type")
	}

	nt := map[string]bool{
		"tg":    false,
		"email": false,
	}

	for _, value := range in.GetNotificationType() {
		if value == "email" {
			nt[value] = true
		}

		if value == "tg" {
			nt[value] = true
		}
	}

	t := in.GetTimestamp().AsTime()
	var reqCreate dto.RequestCreategRPCInServic
	reqCreate.Name = in.GetName()
	reqCreate.Description = in.GetDescription()
	reqCreate.Timestamp = t
	reqCreate.UserID = int(in.GetUserId())
	reqCreate.NotificationType = nt

	resp, err := s.rem.Create(ctx, reqCreate)
	if err != nil {
		// добавить validate для servicLayer errors
		return nil, status.Error(codes.Internal, "failed to create reminder")
	}

	return &reminder.CreateResponse{Id: int64(resp.ReminderID)}, nil
}

func (s *serverAPI) Get(ctx context.Context, in *reminder.GetRequest) (*reminder.GetResponse, error) {

	reqGet := dto.RequestGetgRPCInServic{
		ReminderID: int(in.GetId()),
	}

	result, err := s.rem.Get(ctx, reqGet)
	if err != nil {
		// добавить validate для servicLayer errors
		return nil, status.Error(codes.Internal, "failed to get")
	}

	nt := dto.GRPCInServicMapInSliceString(result.NotificationType)

	r := reminder.GetResponse{
		Name:             result.Name,
		Description:      result.Description,
		NotificationType: nt,
		Timestamp:        timestamppb.New(result.Timestamp),
	}

	return &r, nil
}

func (s *serverAPI) Update(ctx context.Context, in *reminder.UpdateRequest) (*reminder.UpdateResponse, error) {
	if in.GetName() == "" {
		return nil, status.Error(codes.InvalidArgument, "not have name")
	}

	if in.GetDescription() == "" {
		return nil, status.Error(codes.InvalidArgument, "not have description")
	}

	if !in.GetTimestamp().IsValid() {
		return nil, status.Error(codes.InvalidArgument, "imvalid time type")
	}

	nt := dto.GRPCInServicSliceStringInMap(in.GetNotificationType())

	t := in.GetTimestamp().AsTime()
	reqUpdate := dto.RequestUpdateGRPCInServic{}
	reqUpdate.ReminderID = int(in.GetId())
	reqUpdate.Name = in.GetName()
	reqUpdate.Description = in.GetDescription()
	reqUpdate.TimeStamp = t
	reqUpdate.NotificationType = nt

	resp, err := s.rem.Update(ctx, reqUpdate)
	if err != nil {
		// добавить validate для servicLayer errors
		return nil, status.Error(codes.InvalidArgument, "failed to update")
	}

	nt2 := dto.GRPCInServicMapInSliceString(resp.NotificationType)

	return &reminder.UpdateResponse{
		Id:               int64(resp.Reminder_id),
		Name:             resp.Name,
		Description:      resp.Description,
		NotificationType: nt2,
	}, nil
}

func (s *serverAPI) Delete(ctx context.Context, in *reminder.DeleteRequest) (*reminder.DeleteResponse, error) {

	reqDelete := dto.RequestDeletegRRPCInServic{
		ReminderID: int(in.GetId()),
	}

	err := s.rem.Delete(ctx, reqDelete)
	if err != nil {
		// добавить validate для servicLayer errors
		return nil, status.Error(codes.Internal, "failed to delete")
	}
	return &reminder.DeleteResponse{}, nil
}
