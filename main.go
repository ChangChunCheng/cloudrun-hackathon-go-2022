package main

import (
	"math/rand"
	"net/http"

	"github.com/gin-gonic/gin"
)

var (
	F = "F"
	R = "R"
	L = "L"
	T = "T"
)

var letterRunes = []rune("FRLT")

func RandStringRunes(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}

func getPOST(ctx *gin.Context) {
	j := make(map[string]interface{})
	err := ctx.BindJSON(&j)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.String(http.StatusOK, RandStringRunes(1))
}

func main() {
	r := gin.Default()
	r.POST("/", getPOST)
	r.Run(":8080")
}
