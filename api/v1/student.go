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

func (h handler) CreateStudent(ctx *gin.Context) {
	var reqBody models.StudentCreateReq
	var student = &models.Student{}

	if err := ctx.BindJSON(&reqBody); err != nil {
		log.Println("Invalid request body on create student:", err)
		ctx.JSON(400, gin.H{"error": "Invalid request body"})
		return
	}

	if err := halper.DataParser1(reqBody, student); err != nil {
		log.Println("Failed to parse request body:", err)
		ctx.JSON(400, gin.H{"error": "Failed to parse request body"})
		return
	}

	student.StudentID = uuid.New()
	student.CreatedAt = time.Now()
	student.UpdatedAt = time.Now()

	if err := h.storage.StudentRepo().CreateStudent(ctx, student); err != nil {
		log.Println("Failed to create student:", err)
		ctx.JSON(500, gin.H{"error": "Failed to create student"})
		return
	}

	ctx.JSON(201, gin.H{"message": "Student created successfully"})
}

func (h handler) GetStudentList(ctx *gin.Context) {

	var req models.GetListReq
	var err error

	limit := ctx.Query("limit")
	page := ctx.Query("page")

	if req.Limit, err = strconv.Atoi(limit); err != nil {
		log.Println("Invalid limit parameter:", err)
		ctx.JSON(400, gin.H{"error": "invalid limit parameter"})
		return
	}

	if req.Page, err = strconv.Atoi(page); err != nil {
		log.Println("Invalid page parameter:", err)
		ctx.JSON(400, gin.H{"error": "Invalid page parameter"})
		return
	}

	studentsList, err := h.storage.StudentRepo().GetStudentList(ctx, &req)
	if err != nil {
		log.Println("Failed to get student list:", err)
		ctx.JSON(500, gin.H{"error": "Failed to get student list"})
		return
	}

	ctx.JSON(200, studentsList)
}

func (h handler) GetStudentById(ctx *gin.Context) {
	id := ctx.Param("id")

	student, err := h.storage.StudentRepo().GetStudentByID(ctx, id)
	if err != nil {
		log.Println("Failed to get student by ID:", err)
		ctx.JSON(500, gin.H{"error": "Failed to get student by ID"})
		return
	}

	ctx.JSON(200, student)
}

func (h handler) UpdateStudent(ctx *gin.Context) {
	var reqBody models.StudentUpdateReq

	if err := ctx.BindJSON(&reqBody); err != nil {
		log.Println("Invalid request body")
		ctx.JSON(400, gin.H{"error": "Invalid request body"})
		return
	}

	id := ctx.Param("id")

	student, err := h.storage.StudentRepo().GetStudentByID(ctx, id)
	if err != nil {
		log.Println("student not found:", err)
		ctx.JSON(400, gin.H{"error": "Student not found"})
		return
	}

	if err := halper.DataParser1(reqBody, student); err != nil {
		log.Println("Failed to parse data:", err)
		ctx.JSON(400, gin.H{"error": "Failed to parse data"})
		return
	}

	student.UpdatedAt = time.Now()

	if err := h.storage.StudentRepo().UpdateStudent(ctx, student); err != nil {
		log.Println("Failed to UpdateStudent:", err)
		ctx.JSON(500, gin.H{"error": "Failed update student"})
		return
	}

	ctx.JSON(200, student)
}

func (h handler) DeleteStudent(ctx *gin.Context) {
	id := ctx.Param("id")

	if err := h.storage.StudentRepo().DeleteStudent(ctx, id); err != nil {
		log.Println("Failed to delete student:", err)
		ctx.JSON(500, gin.H{"error": "Failed to delete student"})
		return
	}

	ctx.JSON(200, gin.H{"message": "Student deleted successfully"})
}

