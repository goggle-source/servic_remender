package domain

import "time"

/*структура reminder описывает анпоминание
Name - название напоминание, обязательно должно быть и иметь размер максимальный размер - 300
Description - описание напоминания, оно уже не обязательно, максимальный размер - 300
userID - id пользователя, которму и принадлежит это напоминание
Timestamp - время, когда должно сработать напоминание, обязательно должно быть и не должна
быть меньше текущей даты (например: 1.02.15, текущая дата: 1.03.15, так нельзя!)
Notification_type - список методов доставки напоминания до пользователя(email, tg),
хотя бы 1 из способов должен быть доступен
*/
type Reminder struct {
	Name        string
	UserID      int
	Description string
	Timestamp   time.Time
	Nt          Notification_type
}

func NewReminder(name string, UserID int, Description string, timeStamp time.Time, Nt map[string]bool) (Reminder, error) {

	var noti_type Notification_type
	err := noti_type.MapInStruct(Nt)
	if err != nil {
		return Reminder{}, err
	}

	result := Reminder{
		Name:        name,
		Description: Description,
		Timestamp:   timeStamp,
		Nt:          noti_type,
	}

	err = result.Validate()
	if err != nil {
		return Reminder{}, err
	}

	return result, nil
}

func (r *Reminder) Validate() error {

	if !r.Timestamp.After(time.Now()) {
		return InvalidTimeStamp
	}

	//Потом нужно добавить проверку userID, как будет создан сервис user

	if len(r.Name) == 0 {
		return EqualName
	}

	if len(r.Description) > 300 {
		return ErrMaxParameter
	}

	if len(r.Name) > 300 {
		return ErrMaxParameter
	}

	return nil

}

type Notification_type struct {
	Email
	Tg
}

func (n *Notification_type) MapInStruct(nt map[string]bool) error {
	count := 0
	for key, value := range nt {
		if value {
			if key == "email" {
				count += 1
				n.Email.email = key
			}

			if key == "tg" {
				count += 1
				n.Tg.tg = key
			}
		}
	}
	if count == 0 {
		return EqualNotificationType
	}

	return nil
}

type Email struct {
	email string
}

type Tg struct {
	tg string
}
