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

func (h handler) CreateTeacher(ctx *gin.Context) {
	var reqBody models.TeacherCreateReq
	var teacher = &models.Teacher{}

	if err := ctx.BindJSON(&reqBody); err != nil {
		ctx.JSON(400, gin.H{"error": "Invalid request body"})
		log.Println("Invalid request body:", err)
		return
	}
	if err := halper.DataParser1(reqBody, teacher); err != nil {
		ctx.JSON(400, gin.H{"error": "Failed to parse request body"})
		log.Println("Failed to parse request body:", err)
		return
	}

	teacher.TeacherID = uuid.New()
	teacher.CreatedAt = time.Now()
	teacher.UpdatedAt = time.Now()

	if err := h.storage.TeacherRepo().CreateTeacher(ctx, teacher); err != nil {
		log.Println("error on CreateTeacher:", err)
		ctx.JSON(500, gin.H{"error": "Failed to create teacher"})
		return
	}

	ctx.JSON(201, gin.H{"message": "Teacher created successfully"})
}

func (h handler) GetTeachers(ctx *gin.Context) {
	var req models.GetListReq
	// query key "limit" va  "page"
	limit := ctx.Query("limit")
	page := ctx.Query("page")

	var err error
	if req.Limit, err = strconv.Atoi(limit); err != nil {
		ctx.JSON(400, gin.H{"error": "Invalid limit parameter"})
		log.Println("invalid limit parametr")
		return
	}
	if req.Page, err = strconv.Atoi(page); err != nil {
		ctx.JSON(400, gin.H{"error": "Invalid page parameter"})
		log.Println("Invalid page parametr:", err)
		return
	}
	teacherList, err := h.storage.TeacherRepo().GetTeacherList(ctx, &req)
	if err != nil {
		ctx.JSON(500, gin.H{"error": "Failed to get teacher list"})
		log.Println("Failed to get teacher list", err)
		return
	}
	ctx.JSON(200, teacherList)
}

func (h handler) GetTeacherById(ctx *gin.Context) {
	id := ctx.Param("id")

	teacher, err := h.storage.TeacherRepo().GetTeacherByID(ctx, id)
	if err != nil {
		ctx.JSON(500, gin.H{"error": "Failed to get teacher by ID"})
		log.Println("Failed to get teacher by ID:", err)
		return
	}
	ctx.JSON(200, teacher)
}

func (h handler) UpdateTeacher(ctx *gin.Context) {
	var reqBody models.TeacherUpdateReq

	// JSON ma'lumotlarini o'qish
	if err := ctx.BindJSON(&reqBody); err != nil {
		ctx.JSON(400, gin.H{"error": "Invalid request body"})
		log.Println("Invalid request body:", err)
		return
	}
	// ID-ni olish
	id := ctx.Param("id")
	// Mavjud o'qituvchini olish
	teacher, err := h.storage.TeacherRepo().GetTeacherByID(ctx, id)
	if err != nil {
		ctx.JSON(404, gin.H{"error": "Teacher not found"})
		log.Println("Teacher not found:", err)
		return
	}

	// DataParser1 ni qo'llash va xatoni tekshirish
	if err := halper.DataParser1(reqBody, teacher); err != nil {
		ctx.JSON(500, gin.H{"error": "Failed to parse data"})
		log.Println("Failed to parse data:", err)
		return
	}

	// Yangilangan vaqtni o'rnatish
	teacher.UpdatedAt = time.Now()

	// O'qituvchini yangilash
	if err := h.storage.TeacherRepo().UpdateTeacher(ctx, teacher); err != nil {
		ctx.JSON(500, gin.H{"error": "Failed to update teacher"})
		log.Println("Failed to update teacher:", err)
		return
	}
	// Javobni qaytarish
	ctx.JSON(200, teacher)
}

func (h handler) DeleteTeacher(ctx *gin.Context) {
	id := ctx.Param("id")

	err := h.storage.TeacherRepo().DeleteTeacher(ctx, id)
	if err != nil {
		ctx.JSON(500, gin.H{"error": "Failed to delete teacher"})
		log.Println("Failed to delete teacher:", err)
		return
	}
	ctx.JSON(200, gin.H{"message": "Teacher deleted succesfully"})
}
