package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *Handler) GetAllCommentsOfPostsByUserID(ctx *gin.Context) {
	id := ctx.Query("user_id")
	if id == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": errorEmptyIdLine,
		})
		return
	}

	comments, err := h.handApp.Repository.GetAllCommentsOfPostsByUserID(id)
	if err != nil {
		h.handApp.Logger.Error(err)
		ctx.JSON(http.StatusNotFound, gin.H{
			"error": nil,
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"comments": comments,
	})
}
