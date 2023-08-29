package handlers

import (
	"fmt"
	"net/http"

	"github.com/Longreader/dynamic_user_segmentation_service.git/internal/models"
	"github.com/gin-gonic/gin"
)

// @Summary      GetSegmentID
// @Description  Search for segment in database
// @Tags         segments
// @Produce      json
// @Param        segment   path      string  true  "Segment name"
// @Success      200  {int}		integer	"Segment ID"
// @Failure      400  {object}  errorResponse
// @Failure      404  {object}	errorResponse
// @Failure      500  {object}  errorResponse
// @Router       /segments/{segment} [get]
func (h *Handler) getSegment(c *gin.Context) {
	// GET обработчик запроса получения segment

	var input models.Segment

	input.Segment = c.Param("segment")

	if input.Segment == "" {
		newErrorResponse(c, http.StatusBadRequest, fmt.Sprintln("Request must contatin 'segment'"))
		return
	}

	sgmtIDOut, err := h.services.SegmentInterface.GetSegment(input)
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
		"segment": sgmtIDOut,
	})
}

// @Summary      CreateSegment
// @Description  Create a new segments
// @Tags         segments
// @Accept       json
// @Produce      json
// @Param        input	body	models.Segment  true  "Segment name"
// @Success      200  {int}  integer "Segment ID"
// @Failure      400  {object}  errorResponse
// @Failure      404  {object}  errorResponse
// @Failure      500  {object}  errorResponse
// @Router       /segments/ [post]
func (h *Handler) postSegment(c *gin.Context) {
	// POST обработчик запроса создания объекта segment

	var input models.Segment

	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.services.SegmentInterface.CreateSegment(input)
	if err != nil {
		if err.Error() != "ERROR: duplicate key value violates unique constraint \"segments_segment_key\" (SQLSTATE 23505)" {
			newErrorResponse(c, http.StatusInternalServerError, err.Error())
			return
		}
	}

	if (input.Persent != 0) && (input.Persent <= 100) {
		usrs, err := h.services.UserInterface.GetRandUsers(input)
		if err != nil {
			newErrorResponse(c, http.StatusInternalServerError, err.Error())
			return
		}
		for _, usr := range usrs {
			err := h.services.ComparisonInterface.SetUserSegments(models.UserSetSegment{
				SegmentsSet: []string{input.Segment},
				UserId:      usr,
			})
			if err != nil {
				newErrorResponse(c, http.StatusInternalServerError, err.Error())
				return
			}
		}
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})

}

// @Summary      DelSegment
// @Description  Full delete segment from database
// @Tags         segments
// @Produce      json
// @Param        segment   path      string  true  "Segment name"
// @Success      200  {string}		string	"OK"
// @Failure      400  {object}  errorResponse
// @Failure      404  {object}	errorResponse
// @Failure      500  {object}  errorResponse
// @Router       /segments/{segment} [delete]
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
		"status": "OK",
	})

}
