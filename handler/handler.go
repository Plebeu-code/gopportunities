package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func ShowOpeningHandler(ctx *gin.Context) {
  ctx.JSON(http.StatusOK, gin.H{
    "message": "GET oppening",
  })
}

func CreateOpeningHandler(ctx *gin.Context) {
  ctx.JSON(http.StatusOK, gin.H{
    "message": "POST oppening",
  })
}

func DeleteOpeningHandler(ctx *gin.Context) {
  ctx.JSON(http.StatusOK, gin.H{
    "message": "DELETE oppening",
  })
}

func UpdateOpeningHandler(ctx *gin.Context) {
  ctx.JSON(http.StatusOK, gin.H{
    "message": "PUT oppening",
  })
}

func ListOpeningsHandler(ctx *gin.Context) {
  ctx.JSON(http.StatusOK, gin.H{
    "message": "GET oppening",
  })
}