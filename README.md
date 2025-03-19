# Golang Football Web API

This is a golang web api contain Football Data including players, coaches, and clubs. Built using Go, gin-gonic, and deployed using Vercel.

Vercel Url: https://football-webapi.vercel.app/

### Available Api End Point

* GET /api/players - List of football players
* GET /api/clubs - List of football clubs
* GET /api/coaches - List of football coaches

### Getting Started on Local Devices

* Install Required Libraries
```bash
go get github.com/gin-gonic/gin // install Gin Gonic
npm install -g vercel // install Vercel
go run main.go // Run Project
```

# Testing on Vercel Url

### Get Players
* Method: GET
* Url: https://football-webapi.vercel.app/api/players

### Get Football Clubs
* Method: GET
* Url: https://football-webapi.vercel.app/api/clubs

### Get Coaches
* Method: GET
* Url: https://football-webapi.vercel.app/api/coaches
