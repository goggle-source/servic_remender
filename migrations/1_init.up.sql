CREATE TABLE IF NOT EXISTS reminder
(
    id  INTEGER PRIMARY KEY,
    user_id INTEGER NOT NULL,
    name_reminder TEXT NOT NULL,
    message_reminder TEXT NOT NULL,
    status_reminder TEXT NOT NULL,
    time_reminder TIME NOT NULL,
    notificationID INTEGER REFERENCES notificationType (id)
)

CREATE TABLE IF NOT EXISTS notificationType
(
    id INTEGER PRIMARY KEY,
    email BOOLEAN NOT NULL,
    tg BOOLEAN NOT NULL
)