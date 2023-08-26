package handlers

import (
	"fmt"
	"net/http"

	"github.com/Longreader/dynamic_user_segmentation_service.git/internal/models"
	"github.com/gin-gonic/gin"
)

func (h *Handler) getSegment(c *gin.Context) {
	// GET обработчик запроса получения segment

	var input models.Segment

	input.Segment = c.Param("segment")

	if input.Segment == "" {
		newErrorResponse(c, http.StatusBadRequest, fmt.Sprintln("Request must contatin 'segment'"))
		return
	}

	sgmtOut, err := h.services.SegmentInterface.GetSegment(input)
	if err != nil {
		if err.Error() == "sql: no rows in result set" {
			c.JSON(http.StatusOK, map[string]interface{}{
				"segment": "No such segment in base",
			})
			return
		} else {
			newErrorResponse(c, http.StatusInternalServerError, err.Error())
			return
		}
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"segment": sgmtOut,
	})
}

func (h *Handler) postSegment(c *gin.Context) {
	// POST обработчик запроса создания объекта segment

	var input models.Segment

	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.services.SegmentInterface.CreateSegment(input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})

}

func (h *Handler) deleteSegment(c *gin.Context) {
	// DELETE обработчик запроса удаления объекта segment

	var input models.Segment

	input.Segment = c.Param("segment")
	if input.Segment == "" {
		newErrorResponse(c, http.StatusBadRequest, fmt.Sprintln("Request must contatin 'segment'"))
		return
	}

	err := h.services.SegmentInterface.DeleteSegment(input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"deleted": "true",
	})

}
