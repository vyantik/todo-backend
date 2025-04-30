package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/vyantik/todo-backend"
)

func (h *Handler) signUp(c *gin.Context) {
	var input todo.User

	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	_, err := h.services.Authorization.CreateUser(input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]any{
		"message": "success",
	})
}

func (h *Handler) signIn(c *gin.Context) {

}
