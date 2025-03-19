package main

import (
	"golang-webapi/controllers"

	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	// Home route - API Documentation Page
	router.GET("/", func(c *gin.Context) {
		c.Data(http.StatusOK, "text/html; charset=utf-8", []byte(`
			<!DOCTYPE html>
			<html lang="en">
			<head>
				<meta charset="UTF-8">
				<meta name="viewport" content="width=device-width, initial-scale=1.0">
				<title>Football API</title>
				<style>
					body { font-family: Arial, sans-serif; text-align: center; padding: 50px; }
					h1 { color: #2c3e50; }
					p { font-size: 18px; }
					ul { text-align: left; display: inline-block; }
					a { text-decoration: none; color: #2980b9; }
					a:hover { color: #e74c3c; }
				</style>
			</head>
			<body>
				<h1>Welcome to the Football API</h1>
				<p>This is a simple Golang web API providing data about football players, clubs, and coaches.</p>
				
				<h2>Available API Endpoints</h2>
				<ul>
					<li><a href="/api/players">GET /api/players</a> - List of football players</li>
					<li><a href="/api/clubs">GET /api/clubs</a> - List of football clubs</li>
					<li><a href="/api/coaches">GET /api/coaches</a> - List of football coaches</li>
				</ul>

				<p>Use Postman, Curl, or a browser to access these endpoints.</p>
			</body>
			</html>
		`))
	})

	// API Routes
	router.GET("/api/players", controllers.GetPlayers)
	router.GET("/api/clubs", controllers.GetClubs)
	router.GET("/api/coaches", controllers.GetCoaches)

	// Start the server
	router.Run(":8080")
}
