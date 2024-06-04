package services

import "githumb/go-related/nuitteassignment/internal/models"

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
