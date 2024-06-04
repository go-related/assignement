package services

import (
	"githumb/go-related/nuitteassignment/internal/models"
	"githumb/go-related/nuitteassignment/internal/models/response"
)

func convertOccupancies(list []models.Occupancies) []models.OccupancyDTO {
	out := make([]models.OccupancyDTO, len(list))
	for counter, occupancy := range list {
		out[counter] = models.OccupancyDTO{
			Rooms:    occupancy.RoomNr,
			Adults:   occupancy.AdultNr,
			Children: occupancy.ChildrenNr,
		}
	}
	return out
}

func convertHotelResponse(input *response.HotelResponseDTO) []models.HotelResponse {
	var output []models.HotelResponse
	if input != nil && input.Hotels != nil && input.Hotels.Hotels != nil {
		for _, hotel := range *input.Hotels.Hotels {
			output = append(output, models.HotelResponse{
				HotelId:  hotel.Code,
				Currency: hotel.Currency,
				Price:    hotel.MinRate,
			})
		}
	}
	return output
}
