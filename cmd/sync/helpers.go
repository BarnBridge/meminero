package sync

import (
	"encoding/json"
	"io/ioutil"

	"github.com/spf13/viper"
)

func readAndUnmarshalInto(v interface{}) error {
	data, err := ioutil.ReadFile(viper.GetString("file"))
	if err != nil {
		return err
	}

	err = json.Unmarshal(data, v)
	if err != nil {
		return err
	}

	return nil
}
