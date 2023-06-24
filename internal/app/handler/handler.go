package handler

import (
	"K1ngSochialMediaServer/internal/app/pkg"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

const (
	usersURL = "/users"
)

const (
	errorNotFoundUser = "user with id = %s not found in database"
)

type Handler struct {
	handApp *app.Application
}

func NewHandler(a *app.Application) *Handler {
	return &Handler{
		handApp: a,
	}
}

func (h *Handler) Register(router *gin.Engine) {
	router.GET(usersURL, h.List)
}

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
		
		ctx.JSON(http.StatusNotFound, gin.H{
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
