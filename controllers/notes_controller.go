package controllers

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"note-app/controllers/helpers"
	"note-app/models"

	"github.com/gin-gonic/gin"
)

func NotesIndex(c *gin.Context) {
	currentUser := helpers.GetUserFromRequest(c)
	if currentUser == nil || currentUser.ID == 0 {
		c.HTML(
			http.StatusUnauthorized,
			"notes/index.html",
			gin.H{
				"alert": "Unauthorized user",
			},
		)
		return
	}
	notes := models.NotesAll(currentUser)
	c.HTML(
		http.StatusOK,
		"notes/index.html",
		gin.H{
			"notes": notes,
		},
	)
}

func NotesNew(c *gin.Context) {
	c.HTML(
		http.StatusOK,
		"notes/new.html",
		gin.H{},
	)
}

type FormData struct {
	Name    string `json:"name"`
	Content string `json:"content`
}

func NotesCreate(c *gin.Context) {
	currentUser := helpers.GetUserFromRequest(c)
	if currentUser == nil || currentUser.ID == 0 {
		c.HTML(
			http.StatusUnauthorized,
			"notes/index.html",
			gin.H{
				"alert": "Unauthorized user",
			},
		)
		return
	}

	var data FormData
	c.Bind(&data)
	models.NotesCreate(currentUser, data.Name, data.Content)
	c.Redirect(http.StatusMovedPermanently, "notes")
}

func NotesShow(c *gin.Context) {
	currentUser := helpers.GetUserFromRequest(c)
	if currentUser == nil || currentUser.ID == 0 {
		c.HTML(
			http.StatusUnauthorized,
			"notes/index.html",
			gin.H{
				"alert": "Unauthorized user",
			},
		)
		return
	}

	idstr := c.Param("id")
	id, err := strconv.ParseUint(idstr, 10, 64)
	if err != nil {
		log.Fatal("Error : %v", err)
	}
	note := models.NotesFind(currentUser, id)
	c.HTML(
		http.StatusOK,
		"notes/show.html",
		gin.H{
			"note": note,
		},
	)
}

func NotesEdit(c *gin.Context) {
	currentUser := helpers.GetUserFromRequest(c)
	if currentUser == nil || currentUser.ID == 0 {
		c.HTML(
			http.StatusUnauthorized,
			"notes/index.html",
			gin.H{
				"alert": "Unauthorized user",
			},
		)
		return
	}

	idstr := c.Param("id")
	id, err := strconv.ParseUint(idstr, 10, 64)
	if err != nil {
		fmt.Printf("Error : %v", err)
	}
	note := models.NotesFind(currentUser, id)
	c.HTML(
		http.StatusOK,
		"notes/edit.html",
		gin.H{
			"note": note,
		},
	)
}

func NotesUpdate(c *gin.Context) {
	currentUser := helpers.GetUserFromRequest(c)
	if currentUser == nil || currentUser.ID == 0 {
		c.HTML(
			http.StatusUnauthorized,
			"notes/index.html",
			gin.H{
				"alert": "Unauthorized user",
			},
		)
		return
	}

	idstr := c.Param("id")
	id, err := strconv.ParseUint(idstr, 10, 64)
	if err != nil {
		fmt.Printf("Error : %v", err)
	}
	note := models.NotesFind(currentUser, id)
	name := c.PostForm("name")
	content := c.PostForm("content")
	note.NotesUpdate(name, content)
	c.Redirect(http.StatusMovedPermanently, "/notes/"+idstr)
}

func NotesDelete(c *gin.Context) {
	currentUser := helpers.GetUserFromRequest(c)
	if currentUser == nil || currentUser.ID == 0 {
		c.HTML(
			http.StatusUnauthorized,
			"notes/index.html",
			gin.H{
				"alert": "Unauthorized user",
			},
		)
		return
	}

	idstr := c.Param("id")
	id, err := strconv.ParseUint(idstr, 10, 64)
	if err != nil {
		fmt.Printf("Error : %v", err)
	}
	models.NotesMarkDelete(currentUser, id)
	c.Redirect(http.StatusSeeOther, "/notes")
}
