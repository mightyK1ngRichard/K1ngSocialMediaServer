package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *Handler) GetUserPosts(ctx *gin.Context) {
	id := ctx.Query("user_id")
	if id == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": errorEmptyIdLine,
		})
		return
	}

	posts, err := h.handApp.Repository.GetPostsOfUser(id)
	if err != nil {
		h.handApp.Logger.Error(err)
		ctx.JSON(http.StatusNotFound, gin.H{
			"error": nil,
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"posts": posts,
	})
}
