package models

import (
	"time"

	"github.com/google/uuid"
)

// Course represents a course in the system
type Course struct {
	CourseID   uuid.UUID `json:"course_id"`
	CourseName string    `json:"course_name"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}

type GetCoursesList struct {
	Courses []*Course `json:"courses"`
	Count   int       `json:"count"`
}
type CourseCreateReq struct {
	Name string `json:"course_name"`
}

type CourseUpdateReq struct {
	Name      string    `json:"name"`
}
