package handlers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/Longreader/dynamic_user_segmentation_service.git/internal/models"
	"github.com/gin-gonic/gin"
)

// @Summary      GetUser
// @Description  Search for user in database
// @Tags         users
// @Produce      json
// @Param        id   path      int  true  "User outter ID"
// @Success      200  {int}		integer	"User ID"
// @Failure      400  {object}  errorResponse
// @Failure      404  {object}	errorResponse
// @Failure      500  {object}  errorResponse
// @Router       /users/{id} [get]
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

	input.UserID = idInt

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

// @Summary      CreateUser
// @Description  Create a new user
// @Tags         users
// @Accept       json
// @Produce      json
// @Param        input	body	models.User  true  "User outter ID"
// @Success      200  {int}  integer "User ID"
// @Failure      400  {object}  errorResponse
// @Failure      404  {object}  errorResponse
// @Failure      500  {object}  errorResponse
// @Router       /users/ [post]
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

// @Summary      DeleteUser
// @Description  Full delete user from database
// @Tags         users
// @Produce      json
// @Param        id   path     int  true  "User outter ID"
// @Success      200  {string}		string	"OK"
// @Failure      400  {object}  errorResponse
// @Failure      404  {object}	errorResponse
// @Failure      500  {object}  errorResponse
// @Router       /users/{id} [delete]
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

	input.UserID = idInt

	err = h.services.UserInterface.DeleteUser(input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"status": "OK",
	})

}
