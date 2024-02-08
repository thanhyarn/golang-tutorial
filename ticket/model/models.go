package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Ticket struct {
    ID          primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
    Title       string             `json:"title,omitempty"`
    Description string             `json:"description,omitempty"`
    Status      string             `json:"status,omitempty"`
    AssignedTo  primitive.ObjectID `json:"assignedTo,omitempty" bson:"assignedTo,omitempty"`
    Department  string             `json:"department,omitempty"`
    CreatedAt   time.Time          `json:"createdAt,omitempty" bson:"createdAt,omitempty"`
    UpdatedAt   time.Time          `json:"updatedAt,omitempty" bson:"updatedAt,omitempty"`
	ParentTicketID primitive.ObjectID `json:"parent_ticket_id,omitempty" bson:"parent_ticket_id,omitempty"`
}
