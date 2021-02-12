package conf

import (
	"io/ioutil"
	"encoding/json"
)

type StructUrl struct {
        WEBHOOKURL string `toml:"webhookUrl"`
}

func ReadJson() (*StructUrl, error) {
	const configFile = "conf/url.json"

	conf := new(StructUrl)

	cValue, err := ioutil.ReadFile(configFile)
	if err != nil {
		return conf, err
	}

	err = json.Unmarshal([]byte(cValue), conf)
        if err != nil {
                return conf, err
        }

	return conf, nil
}


