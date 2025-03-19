package controllers

import (
	"golang-webapi/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetCoaches(c *gin.Context) {
	coaches := []models.Coach{
		{ID: 1, Name: "Pep Guardiola", Club: "Manchester City"},
		{ID: 2, Name: "Carlo Ancelotti", Club: "Real Madrid"},
		{ID: 3, Name: "Thomas Tuchel", Club: "Bayern Munich"},
		{ID: 4, Name: "Xavi Hernandez", Club: "Barcelona"},
		{ID: 5, Name: "Jurgen Klopp", Club: "Liverpool"},
		{ID: 6, Name: "Erik ten Hag", Club: "Manchester United"},
		{ID: 7, Name: "Luis Enrique", Club: "Paris Saint-Germain"},
		{ID: 8, Name: "Mauricio Pochettino", Club: "Chelsea"},
		{ID: 9, Name: "Mikel Arteta", Club: "Arsenal"},
		{ID: 10, Name: "Massimiliano Allegri", Club: "Juventus"},
		{ID: 11, Name: "Simone Inzaghi", Club: "Inter Milan"},
		{ID: 12, Name: "Stefano Pioli", Club: "AC Milan"},
		{ID: 13, Name: "Diego Simeone", Club: "Atletico Madrid"},
		{ID: 14, Name: "Edin Terzic", Club: "Borussia Dortmund"},
		{ID: 15, Name: "Rudi Garcia", Club: "Napoli"},
		{ID: 16, Name: "Ange Postecoglou", Club: "Tottenham Hotspur"},
		{ID: 17, Name: "Marco Rose", Club: "RB Leipzig"},
		{ID: 18, Name: "Jose Luis Mendilibar", Club: "Sevilla"},
		{ID: 19, Name: "Luis Castro", Club: "Al-Nassr"},
		{ID: 20, Name: "Jorge Jesus", Club: "Al-Hilal"},
	}

	c.JSON(http.StatusOK, coaches)
}
