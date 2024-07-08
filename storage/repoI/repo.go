package repoi

import (
	"context"
	"misis_baholar/models"
)

type TeacherRepoI interface {
	CreateTeacher(ctx context.Context, req *models.Teacher) error
	GetTeacherList(ctx context.Context, req *models.GetListReq) (*models.GetTeachersList, error)
	GetTeacherByID(ctx context.Context, id string) (*models.Teacher, error)
	UpdateTeacher(ctx context.Context, req *models.Teacher) error
	DeleteTeacher(ctx context.Context, id string) error
}

// CourseRepoI interfeysi
type CourseRepoI interface {
	CreateCourse(ctx context.Context, req *models.Course) error
	GetCourseList(ctx context.Context, req *models.GetListReq) (*models.GetCoursesList, error)
	GetCourseByID(ctx context.Context, id string) (*models.Course, error)
	UpdateCourse(ctx context.Context, req *models.Course) error
	DeleteCourse(ctx context.Context, id string) error
}

// GroupRepoI interfeysi
type GroupRepoI interface {
	CreateGroup(ctx context.Context, req *models.Group) error
	GetGroupList(ctx context.Context, req *models.GetListReq) (*models.GetGroupsList, error)
	GetGroupByID(ctx context.Context, id string) (*models.Group, error)
	UpdateGroup(ctx context.Context, req *models.Group) error
	DeleteGroup(ctx context.Context, id string) error
}

// SubjectRepoI interfeysi
type SubjectRepoI interface {
	CreateSubject(ctx context.Context, req *models.Subject) error
	GetSubjectList(ctx context.Context, req *models.GetListReq) (*models.GetSubjectsList, error)
	GetSubjectByID(ctx context.Context, id string) (*models.Subject, error)
	UpdateSubject(ctx context.Context, req *models.Subject) error
	DeleteSubject(ctx context.Context, id string) error
}

// StudentRepoI interfeysi
type StudentRepoI interface {
	CreateStudent(ctx context.Context, req *models.Student) error
	GetStudentList(ctx context.Context, req *models.GetListReq) (*models.GetStudentsList, error)
	GetStudentByID(ctx context.Context, id string) (*models.Student, error)
	UpdateStudent(ctx context.Context, req *models.Student) error
	DeleteStudent(ctx context.Context, id string) error
}

// GradeRepoI interfeysi
type GradeRepoI interface {
	CreateGrade(ctx context.Context, req *models.Grade) error
	GetGradeList(ctx context.Context, req *models.GetListReq) (*models.GetGradesList, error)
	GetGradeByID(ctx context.Context, id string) (*models.Grade, error)
	UpdateGrade(ctx context.Context, req *models.Grade) error
	DeleteGrade(ctx context.Context, id string) error
}
