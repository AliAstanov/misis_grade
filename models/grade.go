package models

import (
	"time"

	"github.com/google/uuid"
)

// Grade represents a grade in the system
type Grade struct {
	GradeID    uuid.UUID `json:"grade_id"`
	GradeName  string    `json:"grade_name"`
	GradeValue int       `json:"grade_value"`
	GradeDate  time.Time `json:"grade_date"`
	SubjectID  uuid.UUID `json:"subject_id"`
	GroupID    uuid.UUID `json:"group_id"`
	StudentID  uuid.UUID `json:"student_id"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}

type GetGradesList struct {
	Grades []*Grade `json:"grades"`
	Count  int      `json:"count"`
}

type GradeCreateReq struct {
	GradeName  string    `json:"grade_name"`
	GradeValue int       `json:"grade_value"`
	GradeDate  time.Time `json:"grade_date"`
	SubjectID  uuid.UUID `json:"subject_id"`
	GroupID    uuid.UUID `json:"group_id"`
	StudentID  uuid.UUID `json:"student_id"`
}

type GradeUpdateReq struct {
	GradeName  string    `json:"grade_name"`
	GradeValue int       `json:"grade_value"`
	GradeDate  time.Time `json:"grade_date"`
	SubjectID  uuid.UUID `json:"subject_id"`
	GroupID    uuid.UUID `json:"group_id"`
	StudentID  uuid.UUID `json:"student_id"`
}
