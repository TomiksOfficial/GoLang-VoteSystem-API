package server

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"voteAPI/postgresql"
	"voteAPI/types"
)

type ServerHandler struct {
	pgdb *postgresql.PGDatabase
}

func (handler *ServerHandler) Init() bool {
	login := ``
	password := ``
	ip := ``
	port := ``
	dbname := ``

	connString := fmt.Sprint(`host=`, ip, ` port=`, port, ` dbname=`, dbname, ` user=`, login, ` password=`, password)

	var err error
	handler.pgdb, err = postgresql.GetPool(context.Background(), connString)
	if err != nil {
		return false
	}

	return true
}

func (handler *ServerHandler) CreateVote(c *gin.Context) {
	var vote types.Vote
	err := c.ShouldBindBodyWithJSON(&vote)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": http.StatusBadRequest,
			"error":  err.Error(),
		})
		return
	}

	err = handler.pgdb.CreateVote(context.Background(), &vote)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": http.StatusBadRequest,
			"error":  err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": http.StatusOK,
	})
}

func (handler *ServerHandler) DoVote(c *gin.Context) {
	var vote types.UserVote
	err := c.ShouldBindBodyWithJSON(&vote)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": http.StatusBadRequest,
			"error":  err.Error(),
		})
		return
	}

	err = handler.pgdb.DoVote(context.Background(), &vote)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": http.StatusBadRequest,
			"error":  err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": http.StatusOK,
	})
}

func (handler *ServerHandler) GetCountVotes(c *gin.Context) {
	var voteinfo types.VoteInfo
	err := c.ShouldBindUri(&voteinfo)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": http.StatusBadRequest,
			"error":  err.Error(),
		})
		return
	}

	var data map[string]any
	data, err = handler.pgdb.GetCountVotes(context.Background(), &voteinfo)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": http.StatusBadRequest,
			"error":  err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, data)
}
