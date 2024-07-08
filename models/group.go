package models

import (
	"time"

	"github.com/google/uuid"
)

// Group represents a group in the system
type Group struct {
	GroupID   uuid.UUID `json:"group_id"`
	GroupName string    `json:"group_name"`
	CourseID  uuid.UUID `json:"course_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type GetGroupsList struct {
	Groups []*Group `json:"groups"`
	Count  int      `json:"count"`
}

type GroupCreateReq struct {
	GroupName string    `json:"group_name"`
	CourseID  uuid.UUID `json:"course_id"`	
}

type GroupUpdateReq struct {
	GroupName string    `json:"group_name"`
	CourseID  uuid.UUID `json:"course_id"`
}
