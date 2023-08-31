package repositories

import (
	"encoding/json"
	"os"
	"quick_notes/src/types"
	"quick_notes/src/utils"
	"strconv"
)

const fileName = "db.json"

func Save(note types.Note) (types.Note, error) {
	notes, err := FindAll()
	if err != nil {
		return note, err
	}
	notes.CurrentId = notes.CurrentId + 1
	note.Id = strconv.Itoa(notes.CurrentId)
	notes.Data[note.Id] = note
	err = persist(notes)
	return note, err
}

func FindAll() (types.Notes, error) {
	var notes types.Notes
	file, err := utils.Open(fileName)
	if err != nil {
		return notes, nil
	}
	defer file.Close()
	err = json.NewDecoder(file).Decode(&notes)
	if err != nil {
		return notes, err
	}
	return notes, nil
}

func persist(notes types.Notes) error {
	file, err := os.Create(fileName)
	if err != nil {
		return err
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ")

	return encoder.Encode(notes)
}
