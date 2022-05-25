package service

import (
	"fmt"

	"example.tld/ps/model"
	"github.com/google/uuid"
)

func get(pet model.Pet) {
	fmt.Printf("Retrieving Pet with ID [%s]", pet.Id)
	// todo call repository
}

func post(pet model.Pet) {
	fmt.Printf("Creating Pet with ID [%s]", pet.Id)
	pet.Id = uuid.NewString()
	// todo call repository
}
