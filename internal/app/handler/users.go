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
			ctx.JSON(http.StatusNotFound, gin.H{
				"error": nil,
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
		textError := fmt.Sprintf(errorNotFoundUser, id)
		ctx.JSON(http.StatusNotFound, gin.H{
			"error": textError,
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
