package model

import "time"

type (
	status = uint8
)

const (
	StatusSaved   status = iota
	StatusDeleted status = iota
	StatusPending status = iota
)

type WorkSpace struct {
	ID         string    `json:"id"`
	Name       string    `json:"name"`
	IDPath     string    `json:"id_path"`
	NamePath   string    `json:"name_path"`
	Status     status    `json:"status"`
	UpdateTime time.Time `json:"update_time"`
	SavedPath  string    `json:"saved_path"`
}
