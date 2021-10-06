package helpers

import (
	"encoding/json"

	"github.com/sirupsen/logrus"
)

func Stringify(data interface{}) interface{} {
	res, err := json.Marshal(data)

	if err != nil {
		logrus.Error(err)
		return nil
	}

	return string(res)
}
