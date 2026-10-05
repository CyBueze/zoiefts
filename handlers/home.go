package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"zoie/config"
	"zoie/views/pages"
)

type HomeHandler struct {
	Cfg *config.Config
}

func NewHomeHandler(cfg *config.Config) *HomeHandler {
	return &HomeHandler{Cfg: cfg}
}

func (h *HomeHandler) Show(c *gin.Context) {
	render(c, http.StatusOK, pages.Home(h.Cfg))
}