package handler

import "github.com/gin-gonic/gin"

type Handler struct {

}

func (h *Handler) InitRouter() *gin.Engine {
	router := gin.New()

	auth := router.Group()
}