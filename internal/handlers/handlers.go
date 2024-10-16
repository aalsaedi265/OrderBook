package handlers

import (
	"net/http"
	"orderbook_tradingEngine/internal/models"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func ShowRegistrationPage(c *gin.Context){
	c.HTML(http.StatusOK, "register.html", gin.H{})
}

func Register(c *gin.Context){
	username := c.PostForm("username")
	password := c.PostForm("password")

	// Hash the password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		c.HTML(http.StatusInternalServerError, "register.html", gin.H{
            "Error": "Failed to process registration",
		})
		return
	}

	user := models.User{
		Username: username,
		Password: string(hashedPassword),
		Balance: 100000.00,	
	}
	result := models.DB.Create(&user)
	if result.Error != nil{
		c.HTML(http.StatusBadRequest, "register.html",gin.H{
			"Error" : "Username already exists",
		})
		return
	}

	// Set session and redirect to dashboard
	session := sessions.Default(c)
	session.Set("user_id", user.ID)
	session.Save()

	c.Redirect(http.StatusFound, "/dashboard")
}

func ShowLoginPage(c *gin.Context){
	c.HTML(http.StatusOK, "login.html", gin.H{})
}

func Login(c *gin.Context){
	username := c.PostForm("username")
	password := c.PostForm("password")

	var user models.User
	result := models.DB.Where("username = ?", username).First(&user)
	if result.Error != nil{
		c.HTML(http.StatusBadRequest, "login.html", gin.H{
			"Error": "Invalid username or password",
		})
		return
	}
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil{
		c.HTML(http.StatusBadRequest, "login.html", gin.H{
			"Error": "Invalid username or password",
		})
		return
	}
	session := sessions.Default(c)
	session.Set("user_id", user.ID)
	session.Save()

	c.Redirect(http.StatusFound, "/dashboard")
}

func Logout(c *gin.Context){
	session := sessions.Default(c)
	session.Clear()
	session.Save()
	c.Redirect(http.StatusFound, "/login")
}

func ShowIndexPage(c *gin.Context){
	c.HTML(http.StatusOK, "index.html", gin.H{})
}

func AuthRequired(c *gin.Context) {
	session := sessions.Default(c)
	userID := session.Get("user_id")
	if userID == nil {
		c.Redirect(http.StatusFound, "/login")
		c.Abort()
		return
	}
	c.Next()
}