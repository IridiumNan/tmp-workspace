package config

import (
	_ "embed"
	"fmt"
	"io"
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
	ConfigName string `toml:"config_name"`

	WorkSpaceDir string `toml:"workspace_dir"`
	MetadataPath string `toml:"metadata_path"`
	LogPath      string `toml:"log_path"`
	ArchiveDir   string `toml:"archive_dir"`

	AutoCreateLinkedDirs bool   `toml:"auto_create_linked_dirs"`
	RetentionHours       uint16 `toml:"retention_hours"`

	LinkedDirs map[string]string `toml:"linked_dirs"`

	CleanupCfg CleanupConfig `toml:"cleanup"`
}

func handleNoConfig() (err error) {
	fmt.Println("default config file not found")

	err = utils.EnsureExist(path.Dir(defaultConfigPath), true)
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

	_, err = toml.Decode(defaultConfig, &Cfg)
	if err != nil {
		return fmt.Errorf("error when decode default config -> %w", err)
	}

	return
}

// LoadConfig : load from the default config path and decode to global variant config.Cfg
func LoadConfig() (err error) {
	var cfgFile *os.File
	cfgFile, err = os.OpenFile(defaultConfigPath, os.O_RDONLY, 0o644)
	if err != nil {
		err = handleNoConfig()
		return
	}

	defer func() {
		closeErr := cfgFile.Close()
		if closeErr == nil {
			return
		}

		if err == nil {
			err = fmt.Errorf("close config file fail: %w", closeErr)

			return
		}

		fmt.Println("close config file fail: %w", closeErr)
	}()

	// use the io.ReadAll to read byte from config file directly
	rawCfgFile, err := io.ReadAll(cfgFile)
	if err != nil {
		return fmt.Errorf("error when read config from %s, err -> %w", defaultConfigPath, err)
	}

	_, err = toml.Decode(string(rawCfgFile), &Cfg)
	if err != nil {
		return fmt.Errorf("error when decode toml config, err -> %w", err)
	}

	return
}
