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
	Tasks []Task `json:"tasks"`
	Smtp  Smtp   `json:"smtp"`
}

//Task is a struct that contains the information for a task
type Task struct {
	Name          string   `json:"name"`
	Frequency     int      `json:"frequency"`
	FrequencyUnit string   `json:"frequency_unit"`
	Url           string   `json:"url"`
	Css           string   `json:"css"`
	Targets       []string `json:"targets"`
	Emails        []string `json:"emails"`
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
