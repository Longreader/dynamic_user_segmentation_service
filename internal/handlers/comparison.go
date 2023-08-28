package handlers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/Longreader/dynamic_user_segmentation_service.git/internal/models"
	"github.com/gin-gonic/gin"
)

// @Summary      UserSegments
// @Description  Create/Delete users segments
// @Tags         comparison
// @Accept       json
// @Produce      json
// @Param        input	body	models.UserSetSegment  true  "Segments and User data"
// @Success      200  {string}  string "OK"
// @Failure      400  {object}  errorResponse
// @Failure      404  {object}  errorResponse
// @Failure      500  {object}  errorResponse
// @Router       /users/add [post]
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
		"status": "OK",
	})

}

// @Summary      UserSegments
// @Description  Search for active segments of user
// @Tags         comparison
// @Produce      json
// @Param        id   path      integer  true  "User ID"
// @Success      200  {object}  models.UserSegments
// @Failure      400  {object}  errorResponse
// @Failure      404  {object}	errorResponse
// @Failure      500  {object}  errorResponse
// @Router       /users/acitve/{id} [get]
func (h *Handler) getActive(c *gin.Context) {

	var input models.User
	var output models.UserSegments
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

	output.Segments = smgts
	output.UserId = input.UserId

	c.JSON(http.StatusOK, output)

}
