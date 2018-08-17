/*
Sniperkit-Bot
- Status: analyzed
*/

package config

import (
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

const configFileName = "webtaskrunner.yaml"

// Load reads the webtaskrunner yaml configuration and deserializes it into
// a Configuration struct which is returned on success
func Load() (*Config, error) {
	configuration := Config{}

	configContent, err := ioutil.ReadFile(configFileName)
	if err != nil {
		return nil, err
	}
	err = yaml.Unmarshal(configContent, &configuration)
	if err != nil {
		return nil, err
	}

	return &configuration, nil
}
