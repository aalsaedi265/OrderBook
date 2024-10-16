package main

import (
	"log"
	"orderbook_tradingEngine/internal/handlers"
	"orderbook_tradingEngine/internal/models"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)


func main() {

	// Initialize the database
    if err := models.InitDB(); err != nil {
        log.Fatal("Failed to connect to database:", err)
    }
	
	//web framework routing, middleware, and template rendering
	 router := gin.Default()

	// session middleware
	store := cookie.NewStore([]byte("secret"))
    router.Use(sessions.Sessions("mysession", store))


    // Serve static files
    router.Static("/static", "./static")

	// Load HTML templates
    router.LoadHTMLGlob("templates/*")

	//Define routes
	router.GET("/", handlers.ShowIndexPage)
    router.GET("/register", handlers.ShowRegistrationPage)
    router.POST("/register", handlers.Register)
    router.GET("/login", handlers.ShowLoginPage)
    router.POST("/login", handlers.Login)
    router.GET("/logout", handlers.Logout)

	//Protected routes
	authorized := router.Group("/")
	authorized.Use(handlers.AuthRequired)
	{
		authorized.GET("/dashboard", handlers.Dashboard)
        authorized.POST("/place_order", handlers.PlaceOrder)
        authorized.POST("/cancel_order", handlers.CancelOrder)

	}
	if err := router.Run(":8080"); err != nil{
		log.Fatal("server run failed: ", err )
	}

}