package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main(){
	fmt.Println("Hello")
	r:=gin.Default()

	r.GET ("/",func (context *gin.Context){

		context.JSON (http.StatusOK, "results")
	})
	r.Run()
}
