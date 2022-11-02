package delivery

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Handler  represent the httpHandler
type Handler struct {
}

type Config struct {
	R       *gin.Engine
	BaseURL string
}

type Response struct {
	ID     string `json:"id"`
	Joke   string `json:"joke"`
	Status int    `json:"status"`
}

func NewHandler(c *Config) {
	handler := &Handler{}
	g := c.R.Group(c.BaseURL)
	g.GET("/badJoke", handler.getBadJoke)

}

func (h *Handler) getBadJoke(c *gin.Context) {
	res := Response{
		ID:     "1",
		Joke:   "Joke",
		Status: 1,
	}
	c.JSON(http.StatusOK, res)
}
