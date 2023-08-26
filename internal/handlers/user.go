package handlers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/Longreader/dynamic_user_segmentation_service.git/internal/models"
	"github.com/gin-gonic/gin"
)

func (h *Handler) getUser(c *gin.Context) {
	// GET обработчик запроса получения user_id

	var input models.User

	idStr := c.Param("id")

	if idStr == "" {
		newErrorResponse(c, http.StatusBadRequest, fmt.Sprintln("Request must contatin 'user_id'"))
		return
	}

	idInt, err := strconv.Atoi(idStr)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, fmt.Sprintln("Request must contatin 'user_id'"))
		return
	}

	input.UserId = idInt

	usrOut, err := h.services.UserInterface.GetUser(input)
	if err != nil {
		if err.Error() == "sql: no rows in result set" {
			c.JSON(http.StatusOK, map[string]interface{}{
				"user_id": 0,
			})
			return
		} else {
			newErrorResponse(c, http.StatusInternalServerError, err.Error())
			return
		}
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"user_id": usrOut,
	})
}

func (h *Handler) postUser(c *gin.Context) {
	// POST обработчик запроса создания объекта user

	var input models.User

	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.services.UserInterface.CreateUser(input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})

}

func (h *Handler) deleteUser(c *gin.Context) {
	// DELETE обработчик запроса удаления объекта user

	var input models.User

	idStr := c.Param("id")

	if idStr == "" {
		newErrorResponse(c, http.StatusBadRequest, fmt.Sprintln("Request must contatin 'user_id'"))
		return
	}

	idInt, err := strconv.Atoi(idStr)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, fmt.Sprintln("Request must contatin 'user_id'"))
		return
	}

	input.UserId = idInt

	err = h.services.UserInterface.DeleteUser(input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"deleted": "true",
	})

}
