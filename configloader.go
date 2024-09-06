package configloader

import (
	"encoding/json"
	"io"
	"os"

	"gopkg.in/yaml.v2"
)

func LoadJSON(p string, config interface{}) error {
    file, err := os.Open(p)
    if err != nil {
        return err
    }
    defer file.Close()

    data, err := io.ReadAll(file)
    if err != nil {
        return err
    }

    return json.Unmarshal(data, config)
}

func LoadYAML(p string, config interface{}) error {
    file, err := os.Open(p)
    if err != nil {
        return err
    }
    defer file.Close()

    data, err := io.ReadAll(file)
    if err != nil {
        return err
    }

    return yaml.Unmarshal(data, config)
}

func LoadENV(config map[string]string) {
    for key := range config {
        if value, exists := os.LookupEnv(key); exists {
            config[key] = value
        }
    }
}
