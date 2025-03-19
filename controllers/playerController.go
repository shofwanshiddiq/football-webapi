package controllers

import (
	"golang-webapi/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetPlayers(c *gin.Context) {
	players := []models.Player{
		{ID: 1, Name: "Lionel Messi", Position: "Forward", Club: "Inter Miami"},
		{ID: 2, Name: "Cristiano Ronaldo", Position: "Forward", Club: "Al-Nassr"},
		{ID: 3, Name: "Kevin De Bruyne", Position: "Midfielder", Club: "Manchester City"},
		{ID: 4, Name: "Kylian Mbappe", Position: "Forward", Club: "Paris Saint-Germain"},
		{ID: 5, Name: "Neymar Jr.", Position: "Forward", Club: "Al-Hilal"},
		{ID: 6, Name: "Erling Haaland", Position: "Forward", Club: "Manchester City"},
		{ID: 7, Name: "Robert Lewandowski", Position: "Forward", Club: "Barcelona"},
		{ID: 8, Name: "Luka Modric", Position: "Midfielder", Club: "Real Madrid"},
		{ID: 9, Name: "Mohamed Salah", Position: "Forward", Club: "Liverpool"},
		{ID: 10, Name: "Karim Benzema", Position: "Forward", Club: "Al-Ittihad"},
		{ID: 11, Name: "Bruno Fernandes", Position: "Midfielder", Club: "Manchester United"},
		{ID: 12, Name: "Harry Kane", Position: "Forward", Club: "Bayern Munich"},
		{ID: 13, Name: "Jude Bellingham", Position: "Midfielder", Club: "Real Madrid"},
		{ID: 14, Name: "Vinicius Jr.", Position: "Forward", Club: "Real Madrid"},
		{ID: 15, Name: "Rodri", Position: "Midfielder", Club: "Manchester City"},
		{ID: 16, Name: "Joshua Kimmich", Position: "Midfielder", Club: "Bayern Munich"},
		{ID: 17, Name: "Antoine Griezmann", Position: "Forward", Club: "Atletico Madrid"},
		{ID: 18, Name: "Trent Alexander-Arnold", Position: "Defender", Club: "Liverpool"},
		{ID: 19, Name: "Pedri", Position: "Midfielder", Club: "Barcelona"},
		{ID: 20, Name: "Casemiro", Position: "Midfielder", Club: "Manchester United"},
	}

	c.JSON(http.StatusOK, players)
}
