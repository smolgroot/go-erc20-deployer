package utils

import (
	"encoding/json"
	"fmt"
	"os"
)

type Config struct {
	ChainId         int64  `json:"chainId"`
	RPCAddress      string `json:"RPCAddress"`
	KeyStore        string `json:"keystore"`
	ContracrAddress string `json:"contractAddress"`
}

func LoadConfiguration(file string) Config {
	var config Config
	configFile, err := os.Open(file)
	defer configFile.Close()
	if err != nil {
		fmt.Println(err.Error())
	}
	jsonParser := json.NewDecoder(configFile)
	jsonParser.Decode(&config)
	return config
}
