package config

import (
    "encoding/json"
    "io/ioutil"
    "os"
    "path/filepath"
)

const configFileName = "gitctx_config.json"

type Config struct {
    Accounts map[string]string `json:"accounts"`
}

func LoadConfig() (Config, error) {
    var config Config
    configFilePath := filepath.Join(os.Getenv("HOME"), configFileName)

    if _, err := os.Stat(configFilePath); os.IsNotExist(err) {
        config.Accounts = make(map[string]string)
        return config, nil
    }

    data, err := ioutil.ReadFile(configFilePath)
    if err != nil {
        return config, err
    }

    err = json.Unmarshal(data, &config)
    return config, err
}

func SaveConfig(config Config) error {
    configFilePath := filepath.Join(os.Getenv("HOME"), configFileName)

    data, err := json.MarshalIndent(config, "", "  ")
    if err != nil {
        return err
    }

    return ioutil.WriteFile(configFilePath, data, 0644)
}
