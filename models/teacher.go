package models

import (
	"time"

	"github.com/google/uuid"
)

// Teacher represents a teacher in the system
type Teacher struct {
	TeacherID uuid.UUID `json:"teacher_id"`
	Name      string    `json:"name"`
	Surname   string    `json:"surname"`
	Email     string    `json:"email"`
	Password  string    `json:"password"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type GetTeachersList struct {
	Teachers []*Teacher `json:"name"`
	Count    int        `json:"count"`
}

type TeacherCreateReq struct {
	Name     string `json:"name"`
	Surname  string `json:"surname"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type TeacherUpdateReq struct {
	Name     string `json:"name"`
	Surname  string `json:"surname"`
	Email    string `json:"email"`
	Password string `json:"password"`
}
