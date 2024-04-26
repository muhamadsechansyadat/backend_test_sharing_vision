package controller

import (
	"backend_test_sharing_vision/api/service"
	"backend_test_sharing_vision/models"
	"backend_test_sharing_vision/util"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// PostController -> PostController
type PostController struct {
	service service.PostService
}

// NewPostController : NewPostController
func NewPostController(s service.PostService) PostController {
	return PostController{
		service: s,
	}
}

// GetPosts : GetPosts controller
func (p PostController) GetPosts(ctx *gin.Context) {
	var posts models.Post

	// keyword := ctx.Query("keyword")
	limitStr := ctx.Param("limit")
	offsetStr := ctx.Param("offset")

	// Konversi string limitStr dan offsetStr ke int64
	limit, err := strconv.Atoi(limitStr)
	if err != nil {
		// Penanganan kesalahan jika konversi gagal
		fmt.Println("Failed to convert limit to int:", err)
		return
	}

	offset, err := strconv.Atoi(offsetStr)
	if err != nil {
		// Penanganan kesalahan jika konversi gagal
		fmt.Println("Failed to convert offset to int:", err)
		return
	}

	status := ctx.Query("status")

	data, total, err := p.service.FindAll(posts, limit, offset, status)

	if err != nil {
		util.ErrorJSON(ctx, http.StatusBadRequest, "Failed to find questions")
		return
	}
	respArr := make([]map[string]interface{}, 0, 0)

	for _, n := range *data {
		resp := n.ResponseMap()
		respArr = append(respArr, resp)
	}

	ctx.JSON(http.StatusOK, &util.Response{
		Success: true,
		Message: "Successfully retrieved articles",
		Data: map[string]interface{}{
			"rows":       respArr,
			"total_rows": total,
		}})
}

// AddPost : AddPost controller
func (p *PostController) AddPost(ctx *gin.Context) {
	var post models.Post
	ctx.ShouldBindJSON(&post)

	if post.Title == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "Title is required"})
		return
	}
	if len(post.Title) < 20 {
		ctx.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "Title length must be at least 20 characters"})
		return
	}
	if post.Content == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "Content is required"})
		return
	}
	if len(post.Content) < 200 {
		ctx.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "Content length must be at least 200 characters"})
		return
	}
	if post.Category == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "Category is required"})
		return
	}
	if len(post.Category) < 3 {
		ctx.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "Category length must be at least 3 characters"})
		return
	}
	err := p.service.Save(post)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "Failed to create article"})
		return
	}
	ctx.JSON(http.StatusCreated, &util.Response{
		Success: true,
		Message: "Successfully created post"})
}

// GetPost : get post by id
func (p *PostController) GetPost(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.ParseInt(idParam, 10, 64) //type conversion string to int64
	if err != nil {
		util.ErrorJSON(c, http.StatusBadRequest, "id invalid")
		return
	}
	var post models.Post
	post.ID = id
	foundPost, err := p.service.Find(post)
	if err != nil {
		util.ErrorJSON(c, http.StatusBadRequest, "Error Finding Post")
		return
	}
	response := foundPost.ResponseMap()

	c.JSON(http.StatusOK, &util.Response{
		Success: true,
		Message: "Successfully retrieved post",
		Data:    &response})

}

func (p *PostController) DeletePost(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.ParseInt(idParam, 10, 64) //type conversion string to int64
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "Id Invalid"})
		return
	}
	if err := p.service.Delete(id); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "Error deleted post"})
		return
	}
	c.JSON(http.StatusOK, &util.Response{
		Success: true,
		Message: "Post deleted successfully",
	})
}

// UpdatePost : get update by id
func (p PostController) UpdatePost(ctx *gin.Context) {
	idParam := ctx.Param("id")

	id, err := strconv.ParseInt(idParam, 10, 64)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "Id Invalid"})
		return
	}
	var post models.Post
	post.ID = id

	postRecord, err := p.service.Find(post)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "TPost with given id not found"})
		return
	}
	ctx.ShouldBindJSON(&postRecord)

	if postRecord.Title == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "Title is required"})
		return
	}
	if len(postRecord.Title) < 20 {
		ctx.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "Title length must be at least 20 characters"})
		return
	}
	if postRecord.Content == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "Content is required"})
		return
	}
	if len(postRecord.Content) < 200 {
		ctx.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "Content length must be at least 200 characters"})
		return
	}
	if postRecord.Category == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "Category is required"})
		return
	}
	if len(postRecord.Category) < 3 {
		ctx.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "Category length must be at least 3 characters"})
		return
	}

	if err := p.service.Update(postRecord); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "Failed to update post"})
		return
	}
	response := postRecord.ResponseMap()

	ctx.JSON(http.StatusOK, &util.Response{
		Success: true,
		Message: "Successfully Updated Post",
		Data:    response,
	})
}
