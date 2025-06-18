package handlers

import (
	"fmt"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/Longreader/dynamic_user_segmentation_service.git/internal/models"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
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

	var newPath = fmt.Sprintf("./storage/%s", time.Now().Format("1-06"))
	err := os.MkdirAll(newPath, os.ModePerm)
	if err != nil {
		logrus.Fatal(err)
	}

	var logPath = fmt.Sprintf("./storage/%s/%s.csv", time.Now().Format("1-06"), time.Now().Format("1-2-06"))

	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	err = h.services.ComparisonInterface.SetUserSegments(input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	file, err := os.OpenFile(logPath, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0755)

	if err != nil {
		logrus.Error(err)
	}
	defer file.Close()

	for _, i := range input.SegmentsSet {
		file.WriteString(fmt.Sprintf("%d;%s;%s;%s;\n", input.UserID, i, "SET", time.Now().Format("1.2.06 3:4:5 -07 MST")))
	}
	for _, i := range input.SegmentsDelete {
		file.WriteString(fmt.Sprintf("%d;%s;%s;%s;\n", input.UserID, i, "DELETE", time.Now().Format("1.2.06 3:4:5 -07 MST")))
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

	input.UserID = idInt

	smgts, err = h.services.ComparisonInterface.GetActiveSegments(input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	output.Segments = smgts
	output.UserID = input.UserID

	c.JSON(http.StatusOK, output)

}
