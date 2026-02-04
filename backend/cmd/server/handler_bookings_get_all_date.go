package main

import (
	"encoding/json"
	"net/http"
	"time"
)

func (cfg *apiConfig) handlerBookingsGetAllDate(w http.ResponseWriter, r *http.Request) {
	type parameters struct {
		Date time.Time `json:"date"`
	}

	decoder := json.NewDecoder(r.Body)
	params := parameters{}
	if err := decoder.Decode(&params); err != nil {
		respondWithError(w, http.StatusInternalServerError, "could not decode params", err)
		return
	}

	bookings, err := cfg.db.GetAllBookingsOnDate(r.Context(), params.Date)
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
			TotalPax:        booking.GroupCount,
			AttendingGroups: booking.AttendingGroups,
		}
	}

	respondWithJSON(w, http.StatusOK, response)
}
