CREATE TABLE IF NOT EXISTS reminder
(
    id  INTEGER PRIMARY KEY,
    user_id INTEGER NOT NULL,
    name_reminder TEXT NOT NULL,
    message_reminder TEXT NOT NULL,
    status_reminder TEXT NOT NULL,
    time_reminder TIME NOT NULL,
    notificationID INTEGER REFERENCES notificationType (id),
    weekday_reminderID INTEGER REFERENCES weekdayReminder (id)
)

CREATE TABLE IF NOT EXISTS notificationType
(
    id INTEGER PRIMARY KEY,
    email BOOLEAN NOT NULL,
    tg BOOLEAN NOT NULL
)

CREATE TABLE IF NOT EXISTS weekdayReminder
(
    id INTEGER PRIMARY KEY,
    monday BOOLEAN NOT NULL,
    tuesday BOOLEAN NOT NULL,
    wednesday BOOLEAN NOT NULL,
    thursday BOOLEAN NOT NULL,
    friday BOOLEAN NOT NULL,
    saturday BOOLEAN NOT NULL,
    sunday BOOLEAN NOT NULL
)