package v1

import (
	"log"
	"misis_baholar/halper"
	"misis_baholar/models"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func (h handler) CreateCourse(ctx *gin.Context) {
	var reqBody models.CourseCreateReq
	var course = &models.Course{}

	if err := ctx.BindJSON(&reqBody); err != nil {
		log.Println("Invalid request body on create course:", err)
		ctx.JSON(400, gin.H{"error": "Invalid request body"})
		return
	}

	if err := halper.DataParser1(reqBody, course); err != nil {
		log.Println("Failed to parse request body:", err)
		ctx.JSON(400, gin.H{"error": "Failed to parse request body"})
		return
	}
	course.CourseID = uuid.New()
	course.CreatedAt = time.Now()
	course.UpdatedAt = time.Now()

	if err := h.storage.CourseRepo().CreateCourse(ctx, course); err != nil {
		log.Println("Failed to create course:", err)
		ctx.JSON(500, gin.H{"error": "Failed to create course"})
		return
	}

	ctx.JSON(201, gin.H{"message": "Course created successfully"})
}

func (h handler) GetCourseList(ctx *gin.Context) {

	var req models.GetListReq
	var err error

	limit := ctx.Query("limit")
	page := ctx.Query("page")

	if req.Limit, err = strconv.Atoi(limit); err != nil {
		log.Println("Invalid limit parameter:", err)
		ctx.JSON(400, gin.H{"error": "Invalid limit parameter"})
	}

	if req.Page, err = strconv.Atoi(page); err != nil {
		log.Println("invalid page parameter:", err)
		ctx.JSON(400, gin.H{"error": "Invalid page parameter"})
		return
	}

	CourseList, err := h.storage.CourseRepo().GetCourseList(ctx, &req)
	if err != nil {
		log.Println("Failed CourseList:", err)
		ctx.JSON(500, gin.H{"error": "Failed get course list on server"})
		return
	}

	ctx.JSON(200, CourseList)
}

func (h handler) GetCourseById(ctx *gin.Context) {
	id := ctx.Param("id")

	course, err := h.storage.CourseRepo().GetCourseByID(ctx, id)
	if err != nil {
		log.Println("Failed to get course by ID:", err)
		ctx.JSON(500, gin.H{"error": "Failed to get course by ID"})
		return
	}

	ctx.JSON(200, course)
}

func (h handler) UpdateCourse(ctx *gin.Context) {
	var reqBody models.CourseUpdateReq
	id := ctx.Param("id")

	if err := ctx.BindJSON(&reqBody); err != nil {
		log.Println("Invalid request body:", err)
		ctx.JSON(400, gin.H{"error": "Invalid request body"})
		return
	}

	course, err := h.storage.CourseRepo().GetCourseByID(ctx, id)
	if err != nil {
		log.Println("Course not found:", err)
		ctx.JSON(404, gin.H{"error": "Course not found"})
		return
	}

	if err := halper.DataParser1(reqBody, course); err != nil {
		log.Println("Failed to parse data:", err)
		ctx.JSON(500, gin.H{"error": "Failed to parse data"})
		return
	}

	course.UpdatedAt = time.Now()

	if err := h.storage.CourseRepo().UpdateCourse(ctx, course); err != nil {
		log.Println("Failed on UpdateCourse:", err)
		ctx.JSON(500, gin.H{"error": "Failed on Update course"})
		return
	}

	ctx.JSON(200, course)
}

func (h handler) DeleteCourse(ctx *gin.Context) {
	id := ctx.Param("id")

	if err := h.storage.CourseRepo().DeleteCourse(ctx, id); err != nil {
		log.Println("Failed to DeleteCourse:", err)
		ctx.JSON(500, gin.H{"error": "Failed to delete course"})
		return
	}

	ctx.JSON(200, gin.H{"message": "Course deleted successfully"})
}
