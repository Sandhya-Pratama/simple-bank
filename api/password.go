package api

import (
	"net/http"

	"github.com/Sandhya-Pratama/simple-bank/util"
	"github.com/gin-gonic/gin"
)

type hashPasswordRequest struct {
	Password string `json:"password" binding:"required"`
}

type hashPasswordResponse struct {
	HashedPassword string `json:"hashed_password"`
}

func (server *Server) hashPassword(ctx *gin.Context) {
	var req hashPasswordRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	hashedPassword, err := util.HashPassword(req.Password)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, hashPasswordResponse{
		HashedPassword: hashedPassword,
	})
}

type validatePasswordRequest struct {
	Password       string `json:"password" binding:"required"`
	HashedPassword string `json:"hashed_password" binding:"required"`
}

func (server *Server) validatePassword(ctx *gin.Context) {
	var req validatePasswordRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	err := util.CheckPassword(req.Password, req.HashedPassword)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Password is valid"})
}
