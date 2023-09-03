package services

import (
	"quick_notes/src/repositories"
	"quick_notes/src/types"
)

func Create(note types.Note) (types.Note, error) {
	return repositories.Save(note)
}

func ListAll() (types.Notes, error) {
	return repositories.FindAll()
}

func FindById(id string) (types.Note, error) {
	return repositories.FindById(id)
}

func DeleteById(id string) error {
	return repositories.DeleteById(id)
}
