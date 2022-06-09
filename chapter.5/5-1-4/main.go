package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

type loadConfigError struct {
	msg string
	err error
}

func (e *loadConfigError) Error() string {
	return fmt.Sprintf("cannot load config: %s (%s)", e.msg, e.err.Error())
}

func (e *loadConfigError) Unwrap() error {
	return e.err
}

type Config struct {
	Name string `json:"name"`
	Age  int64  `json:"age"`
}

func LoadConfig(configFilePath string) (*Config, error) {
	var cfg Config
	data, err := os.ReadFile(configFilePath)
	if err != nil {
		return nil, &loadConfigError{
			msg: fmt.Sprintf("read file `%s`", configFilePath),
			err: err,
		}
	}
	if err = json.Unmarshal(data, &cfg); err != nil {
		return nil, &loadConfigError{
			msg: fmt.Sprintf("parse config file `%s`", configFilePath),
			err: err,
		}
	}
	return &cfg, nil
}

func main() {
	cfg, err := LoadConfig("config.json")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(cfg)
}
