package tools

import (
	"encoding/json"
	"strings"

	"github.com/google/uuid"
)

const RequestKey = "RequestID"

func Data2Str(data interface{}) (ret string) {
	tmp, _ := json.Marshal(data)
	ret = string(tmp)
	return
}

/**
生成UUID
*/
func GenUUID() (string, error) {
	uuidTmp, err := uuid.NewRandom()
	if err != nil {
		return "", err
	}
	return strings.Replace(uuidTmp.String(), "-", "", -1), nil
}
