package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main(){
	r:=gin.Default()

	r.GET ("/",func (context *gin.Context){

		context.JSON (http.StatusOK, "results")
	})
	r.Run()
	fmt.Println("Hello")
}