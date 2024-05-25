package main

import (
	"github.com/gin-gonic/gin"
	"voteAPI/server"
)

func main() {
	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()
	srv := server.ServerHandler{}
	srv.Init()

	router.POST("/CreateVote", srv.CreateVote)
	router.POST("/DoVote", srv.DoVote)
	router.GET("/GetCountVotes/:voteid", srv.GetCountVotes)

	router.Run(":2020")
}
