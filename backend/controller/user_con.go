package controller

import (
	"net/http"
	"strconv"

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
		helper.Error(c, http.StatusBadRequest, err.Error())
		return
	}

	if err := h.service.Register(c.Request.Context(), user); err != nil {
		helper.Error(c, http.StatusInternalServerError, err.Error())
		return
	}

	helper.Success(c, user, "user registered")
}

func (h *UserController) Login(c *gin.Context) {
	var cred model.LoginCredential
	if err := c.ShouldBindJSON(&cred); err != nil {
		helper.Error(c, http.StatusBadRequest, err.Error())
		return
	}

	user, token, err := h.service.Login(c.Request.Context(), cred)
	if err != nil {
		helper.Error(c, http.StatusUnauthorized, err.Error())
		return
	}

	if err := helper.Write(c, token); err != nil {
		helper.Error(c, http.StatusInternalServerError, "failed to set cookie")
		return
	}

	helper.Success(c, user, "login successful", token)
}

func (h *UserController) GetById(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		helper.Error(c, http.StatusBadRequest, "invalid user id")
		return
	}

	user, err := h.service.GetById(c.Request.Context(), id)
	if err != nil {
		helper.Error(c, http.StatusNotFound, err.Error())
		return
	}

	helper.Success(c, user, "user found")
}

func (h *UserController) GetByEmail(c *gin.Context) {
	email := c.Param("email")
	// if err != nil {
	// 	helper.Error(c, http.StatusBadRequest, "invalid user email")
	// 	return
	// }

	user, err := h.service.GetByEmail(c.Request.Context(), email)
	if err != nil {
		helper.Error(c, http.StatusNotFound, err.Error())
		return
	}

	helper.Success(c, user, "user found")
}

func (h *UserController) Update(c *gin.Context) {
	var user model.User

	if err := c.ShouldBindJSON(&user); err != nil {
		helper.Error(c, http.StatusBadRequest, "invalid request body")
		return
	}

	updatedUser, err := h.service.Update(c.Request.Context(), user)
	if err != nil {
		helper.Error(c, http.StatusInternalServerError, err.Error())
		return
	}

	helper.Success(c, updatedUser, "user updated successfully")
}

func (h *UserController) Delete(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		helper.Error(c, http.StatusBadRequest, "invalid user id")
		return
	}

	err = h.service.Delete(c.Request.Context(), id)
	if err != nil {
		helper.Error(c, http.StatusNotFound, err.Error())
		return
	}
}

func (h *UserController) GetAll(c *gin.Context) {
	users, err := h.service.GetAll(c)
	if err != nil {
		helper.Error(c, http.StatusNotFound, err.Error())
		return
	}

	helper.Success(c, users, "users found")
}

func (h *UserController) GetByNim(c *gin.Context) {
	nim := c.Param("nim")
	user, err := h.service.GetByNim(c, nim)
	if err != nil {
		helper.Error(c, http.StatusNotFound, err.Error())
		return
	}

	helper.Success(c, user, "user found")
}

func (h *UserController) GetByName(c *gin.Context) {
	name := c.Param("name")
	users, err := h.service.GetByNim(c, name)
	if err != nil {
		helper.Error(c, http.StatusNotFound, err.Error())
		return
	}

	helper.Success(c, users, "users found")
}

func (h *UserController) GetByRole(c *gin.Context) {
	role := c.Param("role")
	users, err := h.service.GetByNim(c, role)
	if err != nil {
		helper.Error(c, http.StatusNotFound, err.Error())
		return
	}

	helper.Success(c, users, "users found")
}

func (h *UserController) ChangePassword(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		helper.Error(c, http.StatusBadRequest, "invalid user id")
		return
	}

	var req model.ChangePasswordCredential
	if err := c.ShouldBindJSON(&req); err != nil {
		helper.Error(c, http.StatusBadRequest, err.Error())
		return
	}

	if err := h.service.ChangePassword(c.Request.Context(), id, req.OldPassword, req.NewPassword); err != nil {
		helper.Error(c, http.StatusBadRequest, err.Error())
		return
	}

	helper.Success(c, nil, "password changed successfully")
}
