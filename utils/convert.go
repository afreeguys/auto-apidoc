package utils

import (
	"auto-apidoc/log"
	"encoding/json"
)

func JsonToObject(bytes []byte, dist interface{}) error {
	if len(bytes) == 0 {
		return nil
	}
	return json.Unmarshal(bytes, dist)
}

func ObjectToJson(dist interface{}) string {
	if dist == nil {
		return ""
	}
	marshal, err := json.Marshal(dist)
	if err != nil {
		log.Error("to json string err:%v", err)
		return ""
	}
	return string(marshal)
}
