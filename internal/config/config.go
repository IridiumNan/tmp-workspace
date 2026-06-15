package config

import (
	_ "embed"
	"fmt"
	"os"
	"path"

	"github.com/BurntSushi/toml"
	"github.com/IridiumNan/tmp-workspace/internal/utils"
)

var defaultConfigPath = path.Join(utils.HomePath, ".config/tmp-workspace/config.toml")

//go:embed config.toml
var defaultConfig string

var Cfg Config

type CleanupConfig struct {
	AuthClean            bool     `toml:"auto_clean"`
	TriggerCommand       []string `toml:"trigger_command"`
	CheckIntervalMinutes uint16   `toml:"check_interval_minutes"`
}

type Config struct {
	WorkSpaceDir         string `toml:"workspaces"`
	MetadataPath         string `toml:"metadata_path"`
	LogPath              string `toml:"log_path"`
	ArchiveDir           string `toml:"archive_dir"`
	RetentionHours       uint16 `toml:"retention_hours"`
	AutoCreateLinkedDirs bool   `toml:"auto_create_linked_dirs"`

	LinkedDirs map[string]string `toml:"linked_dirs"`

	CleanupCfg CleanupConfig `toml:"cleanup"`
}

func handleNoConfig() (err error) {
	fmt.Println("default config file not found")

	err = utils.EnsureExist(path.Dir(defaultConfigPath))
	if err != nil {
		return
	}

	fmt.Println("create a new one -> ", defaultConfigPath)

	var cfgFile *os.File

	cfgFile, err = os.OpenFile(defaultConfigPath, os.O_CREATE|os.O_RDWR, 0o644)
	if err != nil {
		return fmt.Errorf("error when create a default config file -> %w", err)
	}

	_, err = cfgFile.Write([]byte(defaultConfig))
	if err != nil {
		return fmt.Errorf("error when write default config to file -> %s, err -> %w", defaultConfigPath, err)
	}

	return
}

func LoadConfig() (err error) {
	var cfgFile *os.File
	cfgFile, err = os.OpenFile(defaultConfigPath, os.O_RDONLY, 0o644)
	if err != nil {
		err = handleNoConfig()
		return
	}

	var rawCfgFile []byte
	_, err = cfgFile.Read(rawCfgFile)
	if err != nil {
		return fmt.Errorf("error when read config from %s, err -> %w", defaultConfigPath, err)
	}

	_, err = toml.Decode(string(rawCfgFile), &Cfg)
	if err != nil {
		return fmt.Errorf("error when decode toml config, err -> %w", err)
	}

	return
}
