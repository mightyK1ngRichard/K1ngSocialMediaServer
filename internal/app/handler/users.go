package handler

import (
	"K1ngSochialMediaServer/internal/utils"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"path/filepath"
)

func (h *Handler) List(ctx *gin.Context) {
	id := ctx.Query("id")
	if id == "" {
		users, err := h.handApp.Repository.GetAllUsers()
		if err != nil {
			h.handApp.Logger.Error(err)
			message := fmt.Sprintf("error from database: %s", err)
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"error": message,
			})
			return
		}

		ctx.JSON(http.StatusOK, gin.H{
			"users": users,
		})
		return
	}

	user, err := h.handApp.Repository.GetUserById(id)
	posts, errPost := h.handApp.Repository.GetPostsOfUser(id)

	if err != nil {
		h.handApp.Logger.Error(err)
		message := fmt.Sprintf("error from database: %s", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": message,
		})
		return
	}
	if errPost != nil {
		h.handApp.Logger.Error(errPost)
	}
	user.Posts = posts
	ctx.JSON(http.StatusOK, gin.H{
		"user": user,
	})
}

func (h *Handler) UploadImage(ctx *gin.Context) {
	file, err := ctx.FormFile("file")
	userID := ctx.Param("id")

	if err != nil {
		h.handApp.Logger.Errorf("error from handler/users FormFile returned error: %s", err)
		ctx.JSON(http.StatusBadRequest, "file not found")
		return
	}

	extension := filepath.Ext(file.Filename)
	fileName, err2 := h.handApp.Repository.AddUserImage(extension, userID)
	if err2 != nil {

		h.handApp.Logger.Errorf("error from handler/users db addUserImage: %s", err2)
		ctx.JSON(http.StatusInternalServerError, "Error saving file")
		return
	}

	// TODO: Сделать для видео и тд отдельные работы с бд.
	switch utils.GetFileType(file) {
	case "image":
		fileName = "./static/img/" + fileName
	case "video":
		fileName = "./static/video/" + fileName
	case "file":
		fileName = "./static/files/" + fileName
	case "gif":
		fileName = "./static/gif/" + fileName
	default:
		ctx.JSON(http.StatusBadRequest, "invalid file type")
		return
	}

	if err := ctx.SaveUploadedFile(file, fileName); err != nil {
		h.handApp.Logger.Error(err)
		ctx.JSON(http.StatusInternalServerError, fileName)
		return
	}
	ctx.JSON(http.StatusOK, fileName)
}

// Test TODO: удалить.
func (h *Handler) Test(ctx *gin.Context) {
	type TestData struct {
		Name string `json:"name"`
	}

	header := ctx.GetHeader("Content-Type")

	switch header {
	case "application/x-www-form-urlencoded":
		name := ctx.PostForm("name")
		ctx.JSON(http.StatusOK, gin.H{
			"name": name,
		})
		return

	case "application/json":
		var data TestData
		if err := ctx.ShouldBindJSON(&data); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		name := data.Name
		ctx.JSON(http.StatusOK, gin.H{
			"name": name,
		})
		return

	default:
		h.handApp.Logger.Error("uncorrected header")
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "uncorrected header: " + header,
		})
	}
}
