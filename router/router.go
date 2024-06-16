package router

import "github.com/gin-gonic/gin"

func Initialize() {
	// Inicializa o Router utilizan as configurações Defaults do gin
	router := gin.Default()

	// Definindo uma Rota
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Activated",
		})
	})

	// Setando a porta da API
	router.Run(":8080")
}
