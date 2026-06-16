package store

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"time"

	"github.com/IridiumNan/tmp-workspace/internal/config"
	"github.com/IridiumNan/tmp-workspace/internal/model"
	"github.com/IridiumNan/tmp-workspace/internal/utils"
)

var Metadata *model.MetaData

func firstCreateMetadata() (err error) {
	// logic for first create metadata -> if idPath exist then use the id as name

	Metadata = &model.MetaData{
		LastCheckTime: time.Now(),
		TmpSpaces:     []model.WorkSpace{},
		SavedSpaces:   []model.WorkSpace{},
		DeletedSpaces: []model.WorkSpace{},
	}

	metadataPath := config.Cfg.MetadataPath
	fmt.Printf("metadata file not found, try to create a new one -> %s\n", metadataPath)

	err = utils.EnsureExist(metadataPath, false)

	return
}

func LoadMetadata() (err error) {
	metadataPath := config.Cfg.MetadataPath

	if _, err = os.Stat(metadataPath); os.IsNotExist(err) {
		return firstCreateMetadata()
	}

	metadataFile, openErr := os.OpenFile(metadataPath, os.O_RDWR, 0o644)
	if openErr != nil {
		return fmt.Errorf("error when open metadata file : %w", openErr)
	}

	rawMetadata, _ := io.ReadAll(metadataFile)

	err = json.Unmarshal(rawMetadata, Metadata)
	if err != nil {
		return fmt.Errorf("error when unmarshal metadata from file : %w", err)
	}

	return
}

func NewTmpSpaceMetadata(space *model.WorkSpace) (err error) {
	return
}
