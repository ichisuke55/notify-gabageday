package conf

import (
        "log"
	//"fmt"
        "github.com/BurntSushi/toml"
)

type Config struct {
        METALPAPER      GabageType
        BURNABLE        GabageType
        CANBOTTLE       GabageType
        PETFIBER        GabageType
        PLASTIC         GabageType
}

type GabageType struct {
        Weekdays        []string `toml:"weekdays"`
        Weeks           []int `toml:"weeks"`
        Message         string `"toml:"message""`
}

func ReadConfig() (*Config, error) {
	configFile := "conf/config.toml"

	conf := new(Config)

	_, err := toml.DecodeFile(configFile, &conf)
        if err != nil {
                log.Fatalf("error: %v", err)
        }

	return conf, nil
}


