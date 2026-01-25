package main

import (
	"fmt"
	"net/http"
	"time"
)

type BookingsRow struct {
	BookingID       int32       `json:"booking_id"`
	TourName        string      `json:"tour_name"`
	Date            time.Time   `json:"date"`
	GroupCount      int64       `json:"group_count"`
	TotalPax        interface{} `json:"total_pax"`
	AttendingGroups interface{} `json:"attending_groups"`
}

func (cfg *apiConfig) handlerBookingsGet(w http.ResponseWriter, r *http.Request) {
	fmt.Println("hi from handler")
	bookings, err := cfg.db.GetBookings(r.Context())
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "could not get bookings", err)
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
