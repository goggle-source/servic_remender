package grpc

import (
	"context"
	"servic_remender/internal/models"

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
	Create(ctx context.Context, req models.ServiceCreateRequest) (id int, err error)
	Get(ctx context.Context, id int) (resp models.ServiceGetResponse, err error)
	Update(ctx context.Context, req models.ServiceUpdateRequest) (resp models.ServiceUpdateResponse, err error)
	Delete(ctx context.Context, id int) (err error)
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

	t := in.GetTimestamp().AsTime()
	reqCreate := models.ServiceCreateRequest{}
	reqCreate.Name = in.GetName()
	reqCreate.Message = in.GetDescription()
	reqCreate.Timestamp = t
	reqCreate.Weekday = in.GetWeekday()
	reqCreate.Notification_type = in.GetNotificationType()

	id, err := s.rem.Create(ctx, reqCreate)
	if err != nil {
		return nil, status.Error(codes.Internal, "failed to create reminder")
	}

	return &reminder.CreateResponse{Id: int64(id)}, nil
}

func (s *serverAPI) Get(ctx context.Context, in *reminder.GetRequest) (*reminder.GetResponse, error) {

	result, err := s.rem.Get(ctx, int(in.GetId()))
	if err != nil {
		return nil, status.Error(codes.Internal, "failed to get")
	}

	r := reminder.GetResponse{
		Name:             result.Name,
		Description:      result.Message,
		Weekday:          result.Weekday,
		NotificationType: result.Notification_type,
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

	t := in.GetTimestamp().AsTime()
	reqUpdate := models.ServiceUpdateRequest{}
	reqUpdate.ID = int(in.GetId())
	reqUpdate.Name = in.GetName()
	reqUpdate.Message = in.GetDescription()
	reqUpdate.Timestamp = t
	reqUpdate.Weekday = in.GetWeekday()
	reqUpdate.Notification_type = in.GetNotificationType()

	resp, err := s.rem.Update(ctx, reqUpdate)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, "failed to update")
	}

	return &reminder.UpdateResponse{
		Id:               int64(resp.ID),
		Name:             resp.Name,
		Description:      resp.Message,
		Weekday:          resp.Weekday,
		NotificationType: resp.Notification_type,
	}, nil
}

func (s *serverAPI) Delete(ctx context.Context, in *reminder.DeleteRequest) (*reminder.DeleteResponse, error) {

	err := s.rem.Delete(ctx, int(in.GetId()))
	if err != nil {
		return nil, status.Error(codes.Internal, "failed to delete")
	}
	return &reminder.DeleteResponse{}, nil
}
