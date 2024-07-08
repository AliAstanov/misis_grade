package storage

import (
	"misis_baholar/storage/postgres"
	repoi "misis_baholar/storage/repoI"

	"github.com/jackc/pgx/v5"
)

type StorageI interface {
	TeacherRepo() repoi.TeacherRepoI
	StudentRepo() repoi.StudentRepoI
	CourseRepo() repoi.CourseRepoI
	GroupRepo() repoi.GroupRepoI
	SubjectRepo() repoi.SubjectRepoI
	GradeRepo() repoi.GradeRepoI
}

type storage struct {
	teacherRepo repoi.TeacherRepoI
	studentRepo repoi.StudentRepoI
	courseRepo  repoi.CourseRepoI
	groupRepo repoi.GroupRepoI
	subjectRepo repoi.SubjectRepoI
	gradeRepo repoi.GradeRepoI
}

func NewStorage(db *pgx.Conn) StorageI {
	return &storage{
		teacherRepo: postgres.NewTeacherRepo(db),
		studentRepo: postgres.NewStudentRepo(db),
		courseRepo: postgres.NewCourseRepo(db),
		groupRepo: postgres.NewGroup(db),
		subjectRepo: postgres.NewSubject(db),
		gradeRepo: postgres.NewGrade(db),
	}
}
func (s storage) TeacherRepo() repoi.TeacherRepoI {
	return s.teacherRepo
}
func(s storage) StudentRepo() repoi.StudentRepoI{
	return s.studentRepo
}
func (s storage)CourseRepo() repoi.CourseRepoI{
	return s.courseRepo
}
func (s storage)GroupRepo() repoi.GroupRepoI{
	return s.groupRepo
}
func(s storage)SubjectRepo()repoi.SubjectRepoI{
	return s.subjectRepo
}
func(s storage)GradeRepo() repoi.GradeRepoI{
	return s.gradeRepo
}
