package controller

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"latih.in-be/helper"
	"latih.in-be/model"
	"latih.in-be/service"
)

type QuestionController struct {
	service service.QuestionService
}

func NewQuestionController(s service.QuestionService) *QuestionController {
	return &QuestionController{service: s}
}

func (h *QuestionController) Create(c *gin.Context) {
	var data model.Question
	if err := c.ShouldBindJSON(&data); err != nil {
		helper.Error(c, http.StatusBadRequest, err.Error())
		return
	}

	if err := h.service.Create(c.Request.Context(), data); err != nil {
		helper.Error(c, http.StatusInternalServerError, err.Error())
		return
	}
	helper.Success(c, data, "exam created")
}

func (h *QuestionController) GetById(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		helper.Error(c, http.StatusBadRequest, "invalid data id")
		return
	}

	data, err := h.service.GetById(c.Request.Context(), id)
	if err != nil {
		helper.Error(c, http.StatusNotFound, err.Error())
		return
	}

	helper.Success(c, data, "data found")
}

func (h *QuestionController) GetAll(c *gin.Context) {
	data, err := h.service.GetAll(c.Request.Context())
	if err != nil {
		helper.Error(c, http.StatusNotFound, err.Error())
		return
	}

	helper.Success(c, data, "data found")
}

func (h *QuestionController) Update(c *gin.Context) {
	var data model.Question

	if err := c.ShouldBindJSON(&data); err != nil {
		helper.Error(c, http.StatusBadRequest, "invalid request body")
		return
	}

	updatedData, err := h.service.Update(c, data)
	if err != nil {
		helper.Error(c, http.StatusInternalServerError, err.Error())
		return
	}
	helper.Success(c, updatedData, "data updated")
}

func (h *QuestionController) Delete(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		helper.Error(c, http.StatusBadRequest, "invalid id")
		return
	}
	err = h.service.Delete(c, id)
	if err != nil {
		helper.Error(c, http.StatusInternalServerError, err.Error())
		return
	}
	helper.Success(c, nil, "data deleted")
}
