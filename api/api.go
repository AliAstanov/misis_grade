package api

import (
	v1 "misis_baholar/api/v1"
	"misis_baholar/storage"

	"github.com/gin-gonic/gin"
)

func Api(storage storage.StorageI) {
	router := gin.Default()

	h := v1.NewHandler(storage)

	router.GET("/ping", h.Ping)

	router.POST("/")

	// Teacher endpoints
	router.POST("/create-teacher", h.CreateTeacher)
	router.GET("/get-teacher_list", h.GetTeachers)
	router.GET("/get-teacher-by-id/:id", h.GetTeacherById)
	router.PUT("/update-teachers/:id", h.UpdateTeacher)
	router.DELETE("/delete-teacher/:id", h.DeleteTeacher)

	// Course endpoints
	router.POST("/create-course", h.CreateCourse)
	router.GET("/get-course-list", h.GetCourseList)
	router.GET("/get-course-by-id/:id", h.GetCourseById)
	router.PUT("/update-course/:id", h.UpdateCourse)
	router.DELETE("/delete-course/:id", h.DeleteCourse)

	// Group endpoints
	router.POST("/create-group", h.CreateGroup)
	router.GET("/get-group-list", h.GetGroupList)
	router.GET("/get-group-by-id/:id", h.GetGroupByID)
	router.PUT("/update-group/:id", h.UpdateGroup)
	router.DELETE("/delete-group/:id", h.DeleteGroup)

	// Subject endpoints
	router.POST("/create-subject", h.CreateSubject)
	router.GET("/get-subject-list", h.GetSubjectList)
	router.GET("/get-subject-by-id/:id", h.GetSubjectById)
	router.PUT("/update-subject/:id", h.UpdateSubject)
	router.DELETE("/delete-subject/:id", h.DeleteSubject)

	// Student endpoints
	router.POST("/create-student", h.CreateStudent)
	router.GET("/get-student-list", h.GetStudentList)
	router.GET("/get-student-by-id/:id", h.GetStudentById)
	router.PUT("/update-student/:id", h.UpdateStudent)
	router.DELETE("/delete-student/:id", h.DeleteStudent)

	// Grade endpoints
	router.POST("/create-grade", h.CreateGrade)
	router.GET("/get-grade-list", h.GetGradeList)
	router.GET("/get-grade-by-id/:id", h.GetGradeById)
	router.PUT("/update-grade/:id", h.UpdateGrade)
	router.DELETE("/delete-grade/:id", h.DeleteGrade)

	router.Run(":8080")
}
