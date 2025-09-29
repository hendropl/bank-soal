package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"latih.in-be/helper"
	"latih.in-be/model"
	"latih.in-be/service"
)

type UserController struct {
	service service.UserService
}

func NewUserController(s service.UserService) *UserController {
	return &UserController{service: s}
}

func (h *UserController) Register(c *gin.Context) {
	var user model.RegisterCredential
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.service.Register(c.Request.Context(), user); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "user registered"})
}

func (h *UserController) Login(c *gin.Context) {
	var cred model.LoginCredential
	if err := c.ShouldBindJSON(&cred); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := h.service.Login(c.Request.Context(), cred)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	token, err := helper.GenerateToken(user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if err := helper.Write(c, token); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "login success",
		"user":    user,
		"token":   token,
	})
}
