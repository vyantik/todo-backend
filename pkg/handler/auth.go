package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/vyantik/todo-backend"
)

// @Summary Регистрация нового пользователя
// @Tags Аутентификация
// @Description Регистрация нового пользователя в системе
// @Accept json
// @Produce json
// @Param input body todo.User true "Данные пользователя"
// @Success 200 {object} map[string]interface{} "Успешная регистрация"
// @Failure 400 {object} ErrorResponse "Неверный формат данных"
// @Failure 500 {object} ErrorResponse "Внутренняя ошибка сервера"
// @Router /auth/sign-up [post]
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

type signInInput struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

// @Summary Вход в систему
// @Tags Аутентификация
// @Description Аутентификация пользователя и получение JWT токена
// @Accept json
// @Produce json
// @Param input body signInInput true "Учетные данные"
// @Success 200 {object} map[string]interface{} "Успешный вход"
// @Failure 400 {object} ErrorResponse "Неверный формат данных"
// @Failure 500 {object} ErrorResponse "Неверные учетные данные"
// @Router /auth/sign-in [post]
func (h *Handler) signIn(c *gin.Context) {
	var input signInInput
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	token, err := h.services.Authorization.GenerateToken(input.Username, input.Password)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, "invalid username or password")
		return
	}

	c.JSON(http.StatusOK, map[string]any{
		"message": "success",
		"token":   token,
	})
}
