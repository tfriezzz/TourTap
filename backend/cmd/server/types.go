package main

import (
	"time"

	"github.com/google/uuid"
	"github.com/tfriezzz/tourtap/internal/database"
)

type User struct {
	ID        uuid.UUID `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Email     string    `json:"email"`
	Password  string    `json:"password"`
}

type Tour struct {
	ID        int32     `json:"id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	BasePrice string    `json:"base_price"`
}

type Group struct {
	ID              uuid.UUID            `json:"id"`
	CreatedAt       time.Time            `json:"created_at"`
	UpdatedAt       time.Time            `json:"updated_at"`
	Email           string               `json:"email"`
	Name            string               `json:"name"`
	Pax             int32                `json:"pax"`
	Status          database.GroupStatus `json:"customer_status"`
	RequestedTourID int32                `json:"requested_tour_id"`
	RequestedDate   time.Time            `json:"requested_date"`
	BookingID       int32                `json:"booking_id"`
}

type BookingsRow struct {
	BookingID       int32     `json:"booking_id"`
	TourName        string    `json:"tour_name"`
	Date            time.Time `json:"date"`
	GroupCount      int64     `json:"group_count"`
	TotalPax        any       `json:"total_pax"`
	AttendingGroups any       `json:"attending_groups"`
}

type PaymentWebhookPayload struct {
	Event       int32       `json:"group_id"`
	PaymentData PaymentData `json:"status"`
}

type PaymentData struct{}
