package database

import (
	"context"
	"database/sql"
	"fmt"
	"servic_remender/internal/dto"

	_ "github.com/lib/pq"
)

type Repository struct {
	Db *sql.DB
}

type NotificationTypeRepository struct {
	Email string
	Tg    string
}

func Init(user string, password string, dbName string, port int) *Repository {
	conn := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable port=%d",
		user, password, dbName, port)
	db, err := sql.Open("postgres", conn)
	if err != nil {
		panic(err)
	}

	return &Repository{
		Db: db,
	}
}

func (r *Repository) Create(ctx context.Context, req dto.RequestCreateServicInDB) (dto.ResponseCreateServicInDB, error) {

	tx, errTx := r.Db.BeginTx(ctx, nil)
	if errTx != nil {
		return dto.ResponseCreateServicInDB{}, errTx
	}
	defer func() {
		if errTx != nil {
			tx.Rollback()
		}
	}()

	var ntId int
	err := tx.QueryRowContext(ctx, "INSERT INTO notificationType (email, tg) VALUES($1, $2) RETURNING id",
		req.NotificationType["email"], req.NotificationType["tg"]).Scan(&ntId)

	if err != nil {
		return dto.ResponseCreateServicInDB{}, ValidateErrors(err)
	}

	var remId int
	err = tx.QueryRowContext(ctx, `INSERT INTO reminder 
	(user_id, name_reminder, description_reminder, status_reminder, time_reminder, notificationID) 
	VALUES ($1, $2, $3, $4, $5, $6) RETURNING id`).Scan(&remId)
	if err != nil {
		return dto.ResponseCreateServicInDB{}, ValidateErrors(err)
	}

	if err := tx.Commit(); err != nil {
		return dto.ResponseCreateServicInDB{}, err
	}

	return dto.ResponseCreateServicInDB{ReminderID: remId}, nil
}

func (r *Repository) Get(ctx context.Context, req dto.RequestGetServicInDB) (dto.ResponseGetServicInDB, error) {

	var resp dto.ResponseGetServicInDB
	var ntId int

	tx, errTx := r.Db.BeginTx(ctx, nil)
	if errTx != nil {
		return dto.ResponseGetServicInDB{}, errTx
	}

	err := tx.QueryRowContext(ctx, `SELECT name_reminder, description_reminder, 
	status_reminder, time_reminder, notificationID
	FROM reminder`).Scan(&resp.Name, &resp.Description, &resp.Status, &resp.TimeStamp, &ntId)
	if err != nil {
		tx.Rollback()
		return dto.ResponseGetServicInDB{}, ValidateErrors(err)
	}

	var nt NotificationTypeRepository

	err = tx.QueryRowContext(ctx, "SELECT email, tg FROM notificationType").Scan(&nt.Email, &nt.Tg)
	if err != nil {
		tx.Rollback()
		return dto.ResponseGetServicInDB{}, ValidateErrors(err)
	}

	if nt.Email != "" {
		resp.NotificationType["email"] = true
	}

	if nt.Tg != "" {
		resp.NotificationType["tg"] = true
	}

	if err := tx.Commit(); err != nil {
		tx.Rollback()
		return dto.ResponseGetServicInDB{}, err
	}

	return resp, nil
}

func (r *Repository) Update(ctx context.Context, req dto.RequestUpdateServicInDB) (dto.ResponseUpdateServicInDB, error) {

	tx, errTx := r.Db.BeginTx(ctx, nil)
	if errTx != nil {
		return dto.ResponseUpdateServicInDB{}, errTx
	}
	var ntID int
	err := tx.QueryRow(`UPDATE reminder SET name_reminder = $1, description_reminder = $2, 
	status_reminder = $3, time_reminder = $4 RETURNING notificationID`).Scan(&ntID)
	if err != nil {
		tx.Rollback()
		return dto.ResponseUpdateServicInDB{}, ValidateErrors(err)
	}
	_, err = tx.Exec(`UPDATE notificationType SET email = $1, tg = $2 WHERE id = $3`,
		req.NotificationType["email"], req.NotificationType["tg"], ntID)
	if err != nil {
		tx.Rollback()
		return dto.ResponseUpdateServicInDB{}, ValidateErrors(err)
	}

	if err := tx.Commit(); err != nil {
		tx.Rollback()
		return dto.ResponseUpdateServicInDB{}, err
	}

	result := dto.ResponseUpdateServicInDB{
		Name:             req.Name,
		Description:      req.Description,
		ReminderID:       req.ReminderID,
		NotificationType: req.NotificationType,
		TimeStamp:        req.TimeStamp,
		Status:           req.Status,
	}

	return result, nil
}

func (r *Repository) Delete(ctx context.Context, req dto.RequestDeleteServicInDB) error {
	const op = "database.Delete"

	_, err := r.Db.Exec("DELETE FROM reminder WHERE id = $1", req.ReqeuestID)
	if err != nil {
		return ValidateErrors(err)
	}

	return nil
}
