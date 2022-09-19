package model

import (
	vo "github.com/alexeykirinyuk/go-product-api/internal/service/value_objects"
	"time"
)

type EventType string

const (
	EventTypeProductCreated = EventType("created")
	EventTypeProductUpdated = EventType("updated")
	EventTypeProductDeleted = EventType("removed")
)

type EventStatus int16

const (
	_ EventStatus = iota
	EventStatusCreated
)

type ProductCreatedPayload struct {
	ID          uint64      `json:"id"`
	Name        string      `json:"name"`
	Category    string      `json:"category"`
	Description string      `json:"description"`
	Brand       string      `json:"brand"`
	Cost        float32     `json:"cost"`
	Currency    vo.Currency `json:"currency"`
}

type ProductUpdatedPayload struct {
	ID          uint64      `json:"id"`
	Name        string      `json:"name"`
	Category    string      `json:"category"`
	Description string      `json:"description"`
	Brand       string      `json:"brand"`
	Cost        float32     `json:"cost"`
	Currency    vo.Currency `json:"currency"`
}

type ProductRemovedPayload struct {
	ID uint64 `json:"id"`
}

type Payload struct {
	Created *ProductCreatedPayload `json:"created"`
	Updated *ProductUpdatedPayload `json:"updated"`
	Removed *ProductRemovedPayload `json:"removed"`
}

type ProductEvent struct {
	ID        uint64
	ProductID uint64
	Type      EventType
	Status    EventStatus
	Updated   time.Time
	Payload   Payload
}
