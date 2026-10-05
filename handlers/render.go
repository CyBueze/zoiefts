package handlers

import (
	"log"

	"github.com/a-h/templ"
	"github.com/gin-gonic/gin"
)

// render writes a templ component to the response as HTML.
func render(c *gin.Context, status int, component templ.Component) {
	c.Status(status)
	c.Header("Content-Type", "text/html; charset=utf-8")
	if err := component.Render(c.Request.Context(), c.Writer); err != nil {
		log.Printf("render error: %v", err)
	}
}