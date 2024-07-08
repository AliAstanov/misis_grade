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

func (h handler) CreateSubject(ctx *gin.Context) {
	var reqBody models.SubjectCreateReq
	var subject = &models.Subject{}

	if err := ctx.BindJSON(&reqBody); err != nil {
		log.Println("Invalid request body:", err)
		ctx.JSON(400, gin.H{"error": "Invalid request body"})
		return
	}

	if err := halper.DataParser1(reqBody, subject); err != nil {
		log.Println("Failed to parse request body:", err)
		ctx.JSON(400, gin.H{"error": "Failed to parse request body"})
		return
	}

	subject.SubjectID = uuid.New()
	subject.CreatedAt = time.Now()
	subject.UpdatedAt = time.Now()

	if err := h.storage.SubjectRepo().CreateSubject(ctx, subject); err != nil {
		log.Println("Failed to create subject:", err)
		ctx.JSON(500, gin.H{"error": "Failed to create subject"})
		return
	}

	ctx.JSON(201, gin.H{"message": "Subject created successfully"})
}

func (h handler) GetSubjectList(ctx *gin.Context) {
	var req models.GetListReq
	var err error

	limit := ctx.Query("limit")
	page := ctx.Query("page")

	if req.Limit, err = strconv.Atoi(limit); err != nil {
		log.Println("Invalid limit parameter:", err)
		ctx.JSON(400, gin.H{"error": "Invalid limit parameter"})
		return
	}

	if req.Page, err = strconv.Atoi(page); err != nil {
		log.Println("Invalid page parameter:", err)
		ctx.JSON(400, gin.H{"error": "Invalid page parameter"})
		return
	}

	subjectList, err := h.storage.SubjectRepo().GetSubjectList(ctx, &req)
	if err != nil {
		log.Println("Failed to get subject list:", err)
		ctx.JSON(500, gin.H{"error": "Failed to get subject list"})
		return
	}

	ctx.JSON(200, subjectList)
}

func (h handler) GetSubjectById(ctx *gin.Context) {
	id := ctx.Param("id")

	subject, err := h.storage.SubjectRepo().GetSubjectByID(ctx, id)
	if err != nil {
		log.Println("Failed to get subject by ID:", err)
		ctx.JSON(500, gin.H{"error": "Failed to get subject by ID"})
		return
	}

	ctx.JSON(200, subject)
}

func (h handler) UpdateSubject(ctx *gin.Context) {
	var reqBody models.SubjectUpdateReq

	if err := ctx.BindJSON(&reqBody); err != nil {
		log.Println("Invalid request body:", err)
		ctx.JSON(400, gin.H{"error": "Invalid request body"})
		return
	}

	id := ctx.Param("id")

	subject, err := h.storage.SubjectRepo().GetSubjectByID(ctx, id)
	if err != nil {
		log.Println("Subject not found:", err)
		ctx.JSON(404, gin.H{"error": "Subject not found"})
		return
	}

	if err := halper.DataParser1(reqBody, subject); err != nil {
		log.Println("Failed to parse data:", err)
		ctx.JSON(500, gin.H{"error": "Failed to parse data"})
		return
	}

	subject.UpdatedAt = time.Now()

	if err := h.storage.SubjectRepo().UpdateSubject(ctx, subject); err != nil {
		log.Println("Failed to update subject:", err)
		ctx.JSON(500, gin.H{"error": "Failed to update subject"})
		return
	}

	ctx.JSON(200, subject)
}

func (h handler) DeleteSubject(ctx *gin.Context) {
	id := ctx.Param("id")

	if err := h.storage.SubjectRepo().DeleteSubject(ctx, id); err != nil {
		log.Println("Failed to delete subject:", err)
		ctx.JSON(500, gin.H{"error": "Failed to delete subject"})
		return
	}

	ctx.JSON(200, gin.H{"message": "Subject deleted successfully"})
}
