package handlers

import (
	"github.com/gin-gonic/gin"
	"githumb/go-related/nuitteassignment/internal/models"
	"net/http"
	"time"
)

func (h *Handler) GetHotels(c *gin.Context) {

	var (
		checkinTime      *time.Time
		checkoutTime     *time.Time
		currency         string
		guestNationality string
		hotelIds         []uint
		occupancies      []models.Occupancies
	)
	checkinTimeParameterValue, err := getQueryParamDate(c, "checkin")
	if err == nil {
		checkinTime = &checkinTimeParameterValue
	}

	checkoutTimeParameterValue, err := getQueryParamDate(c, "checkout")
	if err == nil {
		checkoutTime = &checkoutTimeParameterValue
	}
	currency = c.Query("currency")
	guestNationality = c.Query("guestNationality")
	hotelIdParameterValue, err := getQueryParamIdList(c, "hotelIds")
	if err == nil {
		hotelIds = hotelIdParameterValue
	}
	occupanciesParameterValue, err := getQueryParamDataList[models.Occupancies](c, "occupancies")
	if err == nil {
		occupancies = occupanciesParameterValue
	}

	result, err := h.Service.CheckHotelRate(checkinTime, checkoutTime, currency, guestNationality, hotelIds, occupancies)
	if err != nil {
		AbortWithMessage(c, http.StatusInternalServerError, err, "service can't connect to the hotel-api")
	}
	returnOk(c, http.StatusOK, result)
}
