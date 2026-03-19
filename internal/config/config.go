package config

import (
	"encoding/json"
	"fmt"
	"os"
)

type Config struct {
	DbUrl           string `json:"db_url"`
	CurrentUsername string `json:"current_user_name"`
}

const config_file_name = "/.gatorconfig.json"

func Read() (Config, error) {
	config_file, err := getConfigFile()
	if err != nil {
		return Config{}, fmt.Errorf("Error encountered when looking for the config file: %w", err)
	}
	//Get the file and read the json into bytes
	config_file_bytes, err := os.ReadFile(config_file)
	if err != nil {
		return Config{}, fmt.Errorf("Error encountered when reading the config file: %w", err)
	}
	//Read the data into a Config struct
	var res Config
	err = json.Unmarshal(config_file_bytes, &res)
	if err != nil {
		return Config{}, fmt.Errorf("Error encountered when unmarshaling the data: %w", err)
	}

	return res, nil
}

func (c *Config) SetUser(u string) {
	c.CurrentUsername = u
	write(c)
}

func getConfigFile() (string, error) {
	//Get the home dir
	home_dir, err := os.UserHomeDir()
	if err != nil {
		return "", fmt.Errorf("Error encountered when finding the HOME directory: %w", err)
	}
	return home_dir + config_file_name, nil
}

func write(c *Config) error {
	config_file, _ := getConfigFile()
	struct_bytes, err := json.Marshal(c)
	if err != nil {
		return fmt.Errorf("Error marshaling the config struct: %w", err)
	}
	err = os.WriteFile(config_file, struct_bytes, 0644)
	if err != nil {
		return fmt.Errorf("Error writing the config struct: %w", err)
	}
	return nil
}
