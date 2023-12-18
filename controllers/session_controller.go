package controllers

import (
	"fmt"
	"net/http"
	"note-app/helpers"
	"note-app/models"

	"github.com/gin-gonic/gin"
)

func LoginPage(c *gin.Context) {
	c.HTML(
		http.StatusOK,
		"home/login.html",
		gin.H{},
	)
}

func SignupPage(c *gin.Context) {
	c.HTML(
		http.StatusOK,
		"home/signup.html",
		gin.H{},
	)
}

func SignUp(c *gin.Context) {
	email := c.PostForm("email")
	password := c.PostForm("password")
	confirm_password := c.PostForm("confirm_password")

	available := models.UserCheckAvailability(email)
	fmt.Println(available)
	if !available {
		c.HTML(
			http.StatusIMUsed,
			"home/signup.html",
			gin.H{
				"alert": "Email is already registered",
			},
		)
		return
	}

	if password != confirm_password {
		c.HTML(
			http.StatusNotAcceptable,
			"home/signup.html",
			gin.H{
				"alert": "Password missmatch",
			},
		)
		return
	}
	user := models.UserCreate(email, password)
	if user.ID == 0 {
		c.HTML(
			http.StatusNotAcceptable,
			"home/signup.html",
			gin.H{
				"alert": "Unable to create user",
			},
		)
	} else {
		helpers.SessionsSet(c, user.ID)
		c.Redirect(http.StatusMovedPermanently, "/")
	}
}

func Login(c *gin.Context) {
	email := c.PostForm("email")
	password := c.PostForm("password")
	user := models.UserCheck(email, password)
	if user == nil {
		helpers.SessionsSet(c, user.ID)
		c.Redirect(http.StatusMovedPermanently, "/")
	} else {
		c.HTML(
			http.StatusOK,
			"home/login.html",
			gin.H{
				"alert": "Email and/or password mismatch",
			},
		)
	}
}

func Logout(c *gin.Context) {
	helpers.SessionsClear(c)
	c.HTML(
		http.StatusOK,
		"home/login.html",
		gin.H{
			"alert": "Logout successfully",
		},
	)
}
