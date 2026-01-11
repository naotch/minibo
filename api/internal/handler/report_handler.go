package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type IReposrtService interface {
	ReportTotal(userID uint) (int, error)
}

type ReportHandler struct {
	service IReposrtService
}

func NewReportHandler(service IReposrtService) *ReportHandler {
	return &ReportHandler{service: service}
}

func (h *ReportHandler) ReportTotal(c *gin.Context) {
	val := c.MustGet("userId")
	var userID uint
	if fuid, ok := val.(float64); ok {
		userID = uint(fuid)
	} else {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "System error"})
		return
	}

	total, err := h.service.ReportTotal(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"total": total})
}
