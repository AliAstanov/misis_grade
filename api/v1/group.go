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

func (h handler) CreateGroup(ctx *gin.Context) {

	var reqBody models.GroupCreateReq
	var group = &models.Group{}

	if err := ctx.BindJSON(&reqBody); err != nil {
		log.Println("invalid request body on create group: ", err)
		ctx.JSON(400, gin.H{"error": "invalid request body"})
		return
	}

	if err := halper.DataParser1(reqBody, group); err != nil {
		log.Println("Failed to parse request body")
		ctx.JSON(400, gin.H{"error": "Failed to parse request body"})
		return
	}

	group.GroupID = uuid.New()
	group.CreatedAt = time.Now()
	group.UpdatedAt = time.Now()

	if err := h.storage.GroupRepo().CreateGroup(ctx, group); err != nil {
		log.Println("Failed CreateGroup: ", err)
		ctx.JSON(500, gin.H{"error": "Failed to create group"})
		return
	}

	ctx.JSON(201, gin.H{"message: ": "Group created successfully"})
}

func (h handler) GetGroupList(ctx *gin.Context) {
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

	groupList, err := h.storage.GroupRepo().GetGroupList(ctx, &req)
	if err != nil {
		log.Println("Failed to get group list:", err)
		ctx.JSON(500, gin.H{"error": "Failed to get group list"})
		return
	}

	ctx.JSON(200, groupList)
}

func (h handler) GetGroupByID(ctx *gin.Context) {
	id := ctx.Param("id")

	group, err := h.storage.GroupRepo().GetGroupByID(ctx, id)
	if err != nil {
		log.Println("Failed to get group by ID:", err)
		ctx.JSON(500, gin.H{"error": "Failed to get group by ID"})
		return
	}

	ctx.JSON(200, group)
}

func (h handler) UpdateGroup(ctx *gin.Context) {
	var reqBody models.GroupUpdateReq
	id := ctx.Param("id")

	if err := ctx.BindJSON(&reqBody); err != nil {
		log.Println("Invalid request body:", err)
		ctx.JSON(400, gin.H{"error": "Invalid request body"})
		return
	}

	group, err := h.storage.GroupRepo().GetGroupByID(ctx, id)
	if err != nil {
		log.Println("Group not found:", err)
		ctx.JSON(404, gin.H{"error": "Group not found"})
		return
	}

	if err := halper.DataParser1(reqBody, group); err != nil {
		log.Println("Failed to parse data:", err)
		ctx.JSON(500, gin.H{"error": "Failed to parse data"})
		return
	}

	group.UpdatedAt = time.Now()

	if err := h.storage.GroupRepo().UpdateGroup(ctx, group); err != nil {
		log.Println("Failed to update group:", err)
		ctx.JSON(500, gin.H{"error": "Failed to update group"})
		return
	}

	ctx.JSON(200, group)
}

func (h handler) DeleteGroup(ctx *gin.Context) {
	id := ctx.Param("id")

	if err := h.storage.GroupRepo().DeleteGroup(ctx, id); err != nil {
		log.Println("Failed to delete group:", err)
		ctx.JSON(500, gin.H{"error": "Failed to delete group"})
		return
	}

	ctx.JSON(200, gin.H{"message": "Group deleted successfully"})
}

