package main

import (
	"database/sql"
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"time"

	"github.com/tfriezzz/tourtap/internal/database"
)

func (cfg *apiConfig) handlerGroupsCreate(w http.ResponseWriter, r *http.Request) {
	type parameters struct {
		Email           string    `json:"email"`
		Name            string    `json:"name"`
		Pax             int32     `json:"pax"`
		RequestedTourID int32     `json:"requested_tour_id"`
		RequestedDate   time.Time `json:"requested_date"`
	}

	decoder := json.NewDecoder(r.Body)
	params := parameters{}
	if err := decoder.Decode(&params); err != nil {
		respondWithError(w, http.StatusInternalServerError, "could not decode parameters", err)
		return
	}

	getBookingParams := database.GetBookingByTourDateParams{
		TourID: params.RequestedTourID,
		Date:   params.RequestedDate,
	}

	var bookingID int32

	booking, err := cfg.db.GetBookingByTourDate(r.Context(), getBookingParams)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			newBookingParams := database.CreateBookingParams{
				TourID: params.RequestedTourID,
				Date:   params.RequestedDate,
			}
			newBooking, err := cfg.db.CreateBooking(r.Context(), newBookingParams)
			if err != nil {
				respondWithError(w, http.StatusInternalServerError, "could not create booking", err)
				return
			}
			bookingID = newBooking.ID

		} else {
			respondWithError(w, http.StatusInternalServerError, "could not get booking", err)
			return

		}
	} else {
		bookingID = booking.ID
	}

	newGroupParams := database.CreateGroupParams{
		Email:           params.Email,
		Name:            params.Name,
		Pax:             params.Pax,
		RequestedTourID: params.RequestedTourID,
		RequestedDate:   params.RequestedDate,
		BookingID:       bookingID,
	}

	group, err := cfg.db.CreateGroup(r.Context(), newGroupParams)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "could not create group", err)
		return
	}

	respondWithJSON(w, http.StatusCreated, Group{
		ID:              group.ID,
		CreatedAt:       group.CreatedAt,
		UpdatedAt:       group.UpdatedAt,
		Email:           group.Email,
		Name:            group.Name,
		Pax:             group.Pax,
		Status:          group.Status,
		RequestedTourID: group.RequestedTourID,
		RequestedDate:   group.RequestedDate,
		BookingID:       group.BookingID,
	})

	log.Printf("group %v created\n", group.Email)
}
