package core

import (
	"os/user"
	"path/filepath"

	"github.com/ipfsync/common"

	"github.com/spf13/viper"
)

func NewConfig() (*viper.Viper, error) {
	cfg := viper.New()

	usr, err := user.Current()
	if err != nil {
		return nil, err
	}

	dataDir := filepath.Join(usr.HomeDir, ".ipfsync")

	// Make sure dataDir exists and writable
	if err := common.CheckWritable(dataDir); err != nil {
		return nil, err
	}

	cfg.SetConfigName("config")
	cfg.AddConfigPath(dataDir)
	cfg.AddConfigPath(".")

	// Defaults

	// IPFS repository directory
	cfg.SetDefault("repoDir", filepath.Join(dataDir, "ipfs"))

	return cfg, nil
}
