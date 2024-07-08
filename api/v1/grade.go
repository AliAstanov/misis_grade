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

func (h handler) CreateGrade(ctx *gin.Context) {

	var reqBody models.GradeCreateReq
	var grade = &models.Grade{}

	if err := ctx.BindJSON(&reqBody); err != nil {
		log.Println("invalid request body on CreateGrade:", err)
		ctx.JSON(400, gin.H{"error": "invalid request body"})
		return
	}

	if err := halper.DataParser1(reqBody, grade); err != nil {
		log.Println("Failed to parse request body", err)
		ctx.JSON(400, gin.H{"error": "Failed to parse request body"})
		return
	}

	grade.GradeID = uuid.New()
	grade.CreatedAt = time.Now()
	grade.UpdatedAt = time.Now()

	if err := h.storage.GradeRepo().CreateGrade(ctx, grade); err != nil {
		log.Println("Failed CreateGrade:", err)
		ctx.JSON(500, gin.H{"error": "Failed to create grade"})
		return
	}

	ctx.JSON(201, gin.H{"message": "Grade created successfully"})
}

func (h handler) GetGradeList(ctx *gin.Context) {

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
		log.Println("Invalid page parameter", err)
		ctx.JSON(400, gin.H{"error": "Invalid page parameter"})
		return
	}

	GradeList, err := h.storage.GradeRepo().GetGradeList(ctx, &req)
	if err != nil {
		log.Println("Failed to get grade list:", err)
		ctx.JSON(500, gin.H{"error": "Failed getgrades in server"})
		return
	}

	ctx.JSON(200, GradeList)
}

func (h handler) GetGradeById(ctx *gin.Context) {
	id := ctx.Param("id")

	grade, err := h.storage.GradeRepo().GetGradeByID(ctx, id)
	if err != nil {
		log.Println("Failed to get grade by ID:", err)
		ctx.JSON(500, gin.H{"error": "Failed to get grade by ID"})
		return
	}

	ctx.JSON(200, grade)
}

func (h handler) UpdateGrade(ctx *gin.Context) {
	var reqBody models.GradeUpdateReq
	id := ctx.Param("id")

	if err := ctx.BindJSON(&reqBody); err != nil {
		log.Println("Invalid request body:", err)
		ctx.JSON(400, gin.H{"error": "Invalid request body"})
		return
	}

	grade, err := h.storage.GradeRepo().GetGradeByID(ctx, id)
	if err != nil {
		log.Println("Grade not found", err)
		ctx.JSON(404, gin.H{"error": "Grade not found"})
		return
	}

	if err := halper.DataParser1(reqBody, grade); err != nil {
		log.Println("Failed to parse data:", err)
		ctx.JSON(500, gin.H{"error": "Failed to parse data"})
		return
	}

	grade.UpdatedAt = time.Now()

	if err := h.storage.GradeRepo().UpdateGrade(ctx, grade); err != nil {
		log.Println("Failed to update grade:", err)
		ctx.JSON(500, gin.H{"error": "Failed to update grade"})
		return
	}

	ctx.JSON(200, grade)
}

func (h handler) DeleteGrade(ctx *gin.Context) {
	id := ctx.Param("id")

	if err := h.storage.GradeRepo().DeleteGrade(ctx, id); err != nil {
		log.Println("Failed to delete grade:", err)
		ctx.JSON(500, gin.H{"error": "Failed to delete grade"})
		return
	}

	ctx.JSON(200, gin.H{"error": "Grade deleted successfully"})
}
