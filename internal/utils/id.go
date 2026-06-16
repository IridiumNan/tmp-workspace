package utils

import gonanoid "github.com/matoous/go-nanoid/v2"

func GetWorkSpaceID() (ID string) {
	ID, _ = gonanoid.Generate("abcdef1234567890", 12)
	// handle the conflict with other id
	return
}
