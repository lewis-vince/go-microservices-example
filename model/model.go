package model

import "time"

const (
	DOG  = "dog"
	CAT  = "cat"
	BIRD = "bird"
	ROCK = "rock"
)

type Pet struct {
	Id          string    `json:"id,omitempty"`
	Name        string    `json:"name,omitempty"`
	Species     string    `json:"species,omitempty"`
	DateOfBirth time.Time `json:"dateOfBirth,omitempty"`
}
