package servicelayer

import (
	"context"
	"log/slog"
	"servic_remender/internal/domain"
	"servic_remender/internal/dto"
)

const (
	Waiting        = "reminder is wait"
	sending        = "reminder is the senting"
	ErrSending     = "reminder senting error"
	SuccessRequest = "reminder sent successfully"
)

type Repository interface {
	Create(ctx context.Context, req dto.RequestCreateServicInDB) (dto.ResponseCreateServicInDB, error)
	Get(ctx context.Context, req dto.RequestGetServicInDB) (dto.ResponseGetServicInDB, error)
	Update(ctx context.Context, req dto.RequestUpdateServicInDB) (dto.ResponseUpdateServicInDB, error)
	Delete(ctx context.Context, req dto.RequestDeleteServicInDB) error
}

type Servic struct {
	log *slog.Logger
	DB  Repository
}

func CreateNewServic(log *slog.Logger, Db Repository) *Servic {
	return &Servic{
		log: log,
		DB:  Db,
	}
}

func (s *Servic) Create(ctx context.Context, req dto.RequestCreategRPCInServic) (resp dto.ResponseCreategRPCInServic, err error) {
	const op = "servicLayer.Create"

	rem, err := domain.NewReminder(req.Name, req.UserID, req.Description, req.Timestamp, req.NotificationType)
	if err != nil {
		return dto.ResponseCreategRPCInServic{}, ValidationError(op, err)
	}

	reqDB := dto.RequestCreateServicInDB{
		Name:             rem.Name,
		Description:      rem.Description,
		TimeStamp:        rem.Timestamp,
		UserID:           rem.UserID,
		Status:           Waiting,
		NotificationType: req.NotificationType,
	}

	result, err := s.DB.Create(ctx, reqDB)
	if err != nil {
		return dto.ResponseCreategRPCInServic{}, ValidationError(op, err)
	}

	resp.ReminderID = result.ReminderID

	return resp, nil
}

func (s *Servic) Get(ctx context.Context, req dto.RequestGetgRPCInServic) (resp dto.ResponseGetgRPCInServic, err error) {
	const op = "servicLayer.Get"

	reqDB := dto.RequestGetServicInDB{
		ReminderID: req.ReminderID,
	}

	result, err := s.DB.Get(ctx, reqDB)
	if err != nil {
		return dto.ResponseGetgRPCInServic{}, ValidationError(op, err)
	}

	resp.Name = result.Name
	resp.Description = result.Description
	resp.NotificationType = result.NotificationType
	resp.Timestamp = result.TimeStamp

	return resp, nil
}

func (s *Servic) Update(ctx context.Context, req dto.RequestUpdateGRPCInServic) (resp dto.ResponseUpdateGRPCInServic, err error) {
	const op = "servic.Update"

	rem, err := domain.NewReminder(req.Name, req.UserID, req.Description, req.TimeStamp, req.NotificationType)
	if err != nil {
		return dto.ResponseUpdateGRPCInServic{}, ValidationError(op, err)
	}

	res := dto.RequestUpdateServicInDB{
		Name:             rem.Name,
		Description:      rem.Description,
		ReminderID:       req.ReminderID,
		Status:           Waiting,
		TimeStamp:        rem.Timestamp,
		NotificationType: req.NotificationType,
	}

	resultDB, err := s.DB.Update(ctx, res)
	if err != nil {
		return dto.ResponseUpdateGRPCInServic{}, ValidationError(op, err)
	}

	resp.Name = resultDB.Name
	resp.Description = resultDB.Description
	resp.ReminderID = req.ReminderID
	resp.TimeStamp = resultDB.TimeStamp
	resp.NotificationType = resultDB.NotificationType

	return resp, nil

}

func (s *Servic) Delete(ctx context.Context, req dto.RequestDeletegRRPCInServic) (err error) {
	const op = "servic.Delete"

	reqDB := dto.RequestDeleteServicInDB{
		ReqeuestID: req.ReminderID,
	}

	err = s.DB.Delete(ctx, reqDB)
	if err != nil {
		return ValidationError(op, err)
	}

	return nil
}

