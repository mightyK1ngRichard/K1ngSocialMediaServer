package handler

import (
	"K1ngSochialMediaServer/internal/app/pkg"
	"github.com/gin-gonic/gin"
)

const (
	//StaticImgURL = "/static/img"
	usersURL = "/users"
	posts    = "/posts"
	comments = "/comments"
	image    = "user/:id/upload"
	test     = "test"
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
	router.GET(posts, h.GetUserPosts)
	router.GET(comments, h.GetAllCommentsOfPostsByUserID)
	router.POST(image, h.UploadImage)
	router.POST(test, h.Test)
}
