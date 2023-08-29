package handlers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

// @Summary      DownloadAudic
// @Description  Download history file by sent date
// @Tags         utils
// @Produce      octet-stream
// @Param        date   path      string  true  "MONTH-YEAR"
// @Success      200  {object}	[]byte
// @Failure      400  {object}  errorResponse
// @Failure      404  {object}	errorResponse
// @Failure      500  {object}  errorResponse
// @Router       /utils/audit/{date} [get]
func (h *Handler) dowloadAudit(c *gin.Context) {

	dateString := c.Param("date")

	if dateString == "" {
		newErrorResponse(c, http.StatusBadRequest, fmt.Sprintln("Request must contatin 'date:MONTH-YEAR'"))
		return
	}
	fileName, err := h.services.AuditInterface.SendAuditInformation(dateString)

	if err != nil {
		logrus.Fatal(err)
	}

	// Установка заголовка ответа для передачи файла CSV
	c.Header("Content-Disposition", fmt.Sprintf("attachment; filename=./storage/csv_container/%s", fileName))
	c.Header("Content-Type", "text/csv")

	// Отправка данных CSV в ответ на GET запрос
	c.File(fmt.Sprintf("./storage/csv_container/%s.csv", fileName))
}
