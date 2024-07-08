package models

import (
	"time"

	"github.com/google/uuid"
)

// Subject represents a subject in the system
type Subject struct {
	SubjectID   uuid.UUID `json:"subject_id"`
	SubjectName string    `json:"subject_name"`
	GroupID     uuid.UUID `json:"group_id"`
	TeacherID   uuid.UUID `json:"teacher_id"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type GetSubjectsList struct {
	Subjects []*Subject `json:"subjects"`
	Count    int        `json:"count"`
}

type SubjectCreateReq struct {
	SubjectName string    `json:"subject_name"`
	GroupID     uuid.UUID `json:"group_id"`
	TeacherID   uuid.UUID `json:"teacher_id"`
}

type SubjectUpdateReq struct {
	SubjectName string    `json:"subject_name"`
	GroupID     uuid.UUID `json:"group_id"`
	TeacherID   uuid.UUID `json:"teacher_id"`
}
