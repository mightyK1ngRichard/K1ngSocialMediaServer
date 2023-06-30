package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
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
	if err != nil {
		textError := fmt.Sprintf(errorNotFoundUser, id)
		ctx.JSON(http.StatusNotFound, gin.H{
			"error": textError,
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"user": user,
	})
}
