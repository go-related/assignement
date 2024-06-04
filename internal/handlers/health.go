package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *Handler) GetHealth(c *gin.Context) {
	status, err := h.Service.CheckStatus()
	if err != nil {
		AbortWithMessage(c, http.StatusInternalServerError, err, "service can't connect to the hotel-api")
	}
	returnOk(c, http.StatusOK, status)
}
