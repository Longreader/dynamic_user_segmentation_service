package handlers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/Longreader/dynamic_user_segmentation_service.git/internal/models"
	"github.com/gin-gonic/gin"
)

func (h *Handler) postComarison(c *gin.Context) {
	var input models.UserSetSegment

	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	err := h.services.ComparisonInterface.SetUserSegments(input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"success": true,
	})

}

func (h *Handler) getActive(c *gin.Context) {

	var input models.User
	var smgts []string

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

	smgts, err = h.services.ComparisonInterface.GetActiveSegmnents(input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"segments": smgts,
	})

}
