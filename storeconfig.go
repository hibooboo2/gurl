package main

import (
	"encoding/json"
	"io/ioutil"

	"github.com/spf13/viper"
)

func UpdateSettings(v *viper.Viper) error {
	settings := v.AllSettings()
	data, err := json.MarshalIndent(settings, "", "    ")
	if err != nil {
		return err
	}

	err = ioutil.WriteFile(v.ConfigFileUsed(), data, 0644)
	if err != nil {
		return err
	}
	return nil
}
