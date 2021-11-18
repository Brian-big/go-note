package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type note struct {
	ID      string `json:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`
}

var notes = []note{
	{"1", "start", "starting something is always important than never trying"},
	{"2", "plan", "failing to plan is planning to fail"},
	{"3", "get going", "no matter what happens, pursue your dreams to achieve your goal"},
	{"4", "consistency", "always be consistent in what you do"},
}

func getAllNotes(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, notes)
}

func insertNote(c *gin.Context){
	var newNote note
	if err := c.BindJSON(&newNote); err != nil{
		return
	}
	notes = append(notes, newNote)
	c.IndentedJSON(http.StatusCreated, newNote)
}
func getNoteById(c *gin.Context){
	id := c.Param("id")
	for _, n := range notes{
		if n.ID == id {
			c.IndentedJSON(http.StatusOK, n)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "note note found"})
}
func setupRouter() *gin.Engine {

	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})

	r.GET("/notes", getAllNotes)

	r.POST("/notes", insertNote)

	r.GET("/notes/:id", getNoteById)

	return r
}

func main() {
	r := setupRouter()
	// Listen and Server in 0.0.0.0:8080
	r.Run(":8010")
}
