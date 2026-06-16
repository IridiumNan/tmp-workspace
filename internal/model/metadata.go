package model

import "time"

type MetaData struct {
	LastCheckTime time.Time   `json:"last_check_time"`
	TmpSpaces     []WorkSpace `json:"tmp_spaces"`
	SavedSpaces   []WorkSpace `json:"saved_spaces"`
	DeletedSpaces []WorkSpace `json:"deleted_spaces"`
}
