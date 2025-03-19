package controllers

import (
	"golang-webapi/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetClubs(c *gin.Context) {
	clubs := []models.Club{
		{ID: 1, Name: "Manchester City", City: "Manchester"},
		{ID: 2, Name: "Real Madrid", City: "Madrid"},
		{ID: 3, Name: "Bayern Munich", City: "Munich"},
		{ID: 4, Name: "Barcelona", City: "Barcelona"},
		{ID: 5, Name: "Liverpool", City: "Liverpool"},
		{ID: 6, Name: "Manchester United", City: "Manchester"},
		{ID: 7, Name: "Paris Saint-Germain", City: "Paris"},
		{ID: 8, Name: "Chelsea", City: "London"},
		{ID: 9, Name: "Arsenal", City: "London"},
		{ID: 10, Name: "Juventus", City: "Turin"},
		{ID: 11, Name: "Inter Milan", City: "Milan"},
		{ID: 12, Name: "AC Milan", City: "Milan"},
		{ID: 13, Name: "Atletico Madrid", City: "Madrid"},
		{ID: 14, Name: "Borussia Dortmund", City: "Dortmund"},
		{ID: 15, Name: "Napoli", City: "Naples"},
		{ID: 16, Name: "Tottenham Hotspur", City: "London"},
		{ID: 17, Name: "RB Leipzig", City: "Leipzig"},
		{ID: 18, Name: "Sevilla", City: "Seville"},
		{ID: 19, Name: "Al-Nassr", City: "Riyadh"},
		{ID: 20, Name: "Al-Hilal", City: "Riyadh"},
	}

	c.JSON(http.StatusOK, clubs)
}
