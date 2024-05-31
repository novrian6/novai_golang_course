package middleware

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Render(c *gin.Context, data gin.H, templateName string){

	switch c.Request.Header.Get ("Accept"){
	case "application/json":
		c.JSON (http.StatusOK, data["payload"])
	default:
		c.HTML (http.StatusOK, templateName, data)

	}

}