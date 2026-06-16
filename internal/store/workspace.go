package store

import (
	"fmt"
	"os"
	"path"
	"time"

	"github.com/IridiumNan/tmp-workspace/internal/config"
	"github.com/IridiumNan/tmp-workspace/internal/model"
	"github.com/IridiumNan/tmp-workspace/internal/utils"
)

const (
	subNameDir = "byName"
	subIDDIR   = "byID"
)

// GetWorkSpaceByName only join the path and return the expected name-path
func GetWorkSpaceByName(name string) (namePath string) {
	return path.Join(config.Cfg.WorkSpaceDir, subNameDir, namePath)
}

// GetWorkspaceByID only join the path and return the expect name-path
func GetWorkspaceByID(id string) (idPath string) {
	return path.Join(config.Cfg.WorkSpaceDir, subIDDIR, id)
}

// NewTmpSpace create new tmp-workspace in disk and return the space for metadata update
func NewTmpSpace(spaceName string) (space *model.WorkSpace, err error) {
	id := utils.GetWorkSpaceID()
	if spaceName == "" {
		spaceName = id
	}

	idPath := GetWorkspaceByID(id)
	namePath := GetWorkSpaceByName(spaceName)

	utils.BakIfExist(idPath, true)
	utils.BakIfExist(namePath, true)

	err = utils.EnsureExist(idPath, true)
	if err != nil {
		err = fmt.Errorf("error when create the idPath -> %s, err -> %w", idPath, err)
		return
	}

	// create soft link which points to the id-pth
	err = os.Symlink(idPath, namePath)
	if err != nil {
		err = fmt.Errorf("error when create the namePath link: %s -> %s, err -> %w", idPath, namePath, err)
		return
	}

	space = &model.WorkSpace{
		ID:         id,
		Name:       spaceName,
		IDPath:     idPath,
		NamePath:   namePath,
		Status:     model.StatusPending,
		UpdateTime: time.Now(),
		SavedPath:  model.NotSavedPath,
	}

	return
}
