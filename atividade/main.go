package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Contato struct {
	Name string `json:"name"`
	Email string `json:"email"`
	Message string `json:"message"`
}

func main() {
	router := gin.Default()

	router.Static("/static", "./static")
	router.LoadHTMLGlob("static/*")

	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})

	var Contatos []Contato

	router.POST("/message", func(c *gin.Context) {
		var contato Contato
		contato.Name = c.PostForm("name")
		contato.Email = c.PostForm("email")
		contato.Message = c.PostForm("message")

		Contatos = append(Contatos, contato)
		c.Redirect(http.StatusFound, "/message?name=" + contato.Name)
	})

	router.GET("/message", func(c *gin.Context) {
		name := c.Query("name")
		c.HTML(http.StatusOK, "message.html", gin.H{"Name": name})
	})

	router.GET("/messages", func(c *gin.Context) {
		c.HTML(http.StatusOK, "messages.html", gin.H{"Contatos": Contatos})
	})

	if err := router.Run(":8080"); err != nil {
		panic(err)
	}
}