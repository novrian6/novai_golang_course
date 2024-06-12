package main

import (
	"errors"
	"fmt"
	"github.com/Sirupsen/logrus"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	logrus.Error(errors.New("whoops"))
	fmt.Println("hello")

	r:=gin.Default()

	r.get("/", func (context *gin.Context){
		context.JSON (http.StatusOK, gin.H{"status":"Ok"})
	})
	r.run(":8082")
}