package config

import (
	"encoding/json"
	"log"
	"os"
	"path"
)

const (
	CONFIG_FILE = "config.json"
)

//Config is the main configuration struct
type Config struct {
	Smtp   Smtp   `json:"smtp"`
	Server Server `json:"server"`
}

type Server struct {
	Chrome string `json:"chrome"`
}

type Smtp struct {
	Host     string `json:"host"`
	Port     int    `json:"port"`
	Secure   bool   `json:"secure"`
	User     string `json:"user"`
	Password string `json:"password"`
}

// New Loads the config file
func New() *Config {
	currentPath, err := os.Getwd()
	if err != nil {
		log.Fatalln(err)
	}

	bytes, err := os.ReadFile(path.Join(currentPath, CONFIG_FILE))
	if err != nil {
		log.Fatalln("Loading config file file:", err)
	}
	config := &Config{}
	if err := json.Unmarshal(bytes, config); err != nil {
		log.Fatalln("config file format error:", err)
	}
	return config
}
