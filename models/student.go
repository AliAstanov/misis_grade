package models

import (
	"time"

	"github.com/google/uuid"
)

// Student represents a student in the system
type Student struct {
	StudentID uuid.UUID `json:"student_id"`
	Name      string    `json:"name"`
	Surname   string    `json:"surname"`
	GroupID   uuid.UUID `json:"group_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type GetStudentsList struct {
	Students []*Student `json:"students"`
	Count    int        `json:"count"`
}

type StudentCreateReq struct {
	Name    string    `json:"name"`
	Surname string    `json:"surname"`
	GroupID uuid.UUID `json:"group_id"`
}

type StudentUpdateReq struct {
	Name    string    `json:"name"`
	Surname string    `json:"surname"`
	GroupID uuid.UUID `json:"group_id"`
}
