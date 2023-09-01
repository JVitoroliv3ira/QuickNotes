package services

import (
	"quick_notes/src/repositories"
	"quick_notes/src/types"
)

func Create(note types.Note) (types.Note, error) {
	return repositories.Save(note)
}
