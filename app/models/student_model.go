package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Student struct {

    Id           primitive.ObjectID  `json:"id,omitempty"`
    Name          string             `json:"name,omitempty" validate:"required"`
	Student_id    int                `json:"student_id,omitempty" validate:"required"`
    Register_no   int                `json:"register_no,omitempty" validate:"required"`
    Department    string             `json:"department,omitempty" validate:"required"`
}