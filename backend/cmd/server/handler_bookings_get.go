package main

import (
	"net/http"
)

func (cfg *apiConfig) handlerBookingsGet(w http.ResponseWriter, r *http.Request) {
	// fmt.Println("hi from handler")
	bookings, err := cfg.db.GetBookings(r.Context())
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "could not get bookings", err)
		return
	}

	response := make([]BookingsRow, len(bookings))

	for i, booking := range bookings {
		response[i] = BookingsRow{
			BookingID:       booking.BookingID,
			TourName:        booking.TourName,
			Date:            booking.Date,
			GroupCount:      booking.GroupCount,
			TotalPax:        booking.TotalPax,
			AttendingGroups: booking.AttendingGroups,
		}
	}

	respondWithJSON(w, http.StatusOK, response)
}
