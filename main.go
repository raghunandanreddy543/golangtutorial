package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func main() {
	r := gin.Default()
	r.GET("", func(context *gin.Context) {
		firstNumber := context.Query("fn")
		secondNumber := context.Query("sn")
		fn, _ := strconv.Atoi(firstNumber)
		sn, _ := strconv.Atoi(secondNumber)
		add := fn + sn
		msg := fmt.Sprintf("Addition is  %d", add)
		fmt.Println(msg)
		context.String(http.StatusOK, string(msg))
	})
	r.Run()
}
