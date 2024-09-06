package configloader

import (
    "os"
    "testing"
)

func TestLoadJSON(t *testing.T) {
    type Config struct {
        AppName string `json:"app_name"`
        Port    int    `json:"port"`
    }

    var config Config
    err := LoadJSON("config.json", &config)
    if err != nil {
        t.Fatalf("Failed to load JSON: %v", err)
    }

    if config.AppName != "TestApp" || config.Port != 8080 {
        t.Fatalf("Unexpected config values: %+v", config)
    }
}

func TestLoadYAML(t *testing.T) {
    type Config struct {
        AppName string `yaml:"app_name"`
        Port    int    `yaml:"port"`
    }

    var config Config
    err := LoadYAML("config.yaml", &config)
    if err != nil {
        t.Fatalf("Failed to load YAML: %v", err)
    }

    if config.AppName != "TestApp" || config.Port != 8080 {
        t.Fatalf("Unexpected config values: %+v", config)
    }
}

func TestLoadENV(t *testing.T) {
    os.Setenv("APP_NAME", "EnvApp")
    os.Setenv("PORT", "9090")

    config := map[string]string{
        "APP_NAME": "",
        "PORT":     "",
    }

    LoadENV(config)

    if config["APP_NAME"] != "EnvApp" || config["PORT"] != "9090" {
        t.Fatalf("Unexpected config values: %+v", config)
    }
}
