// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package subscriptiontesting

import (
	"time"
)

type Message struct {
	ID        string    `json:"id"`
	Text      string    `json:"text"`
	CreatedBy string    `json:"createdBy"`
	CreatedAt time.Time `json:"createdAt"`
}
