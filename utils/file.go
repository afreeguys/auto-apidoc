package utils

import (
	"auto-apidoc/log"
	"encoding/json"
	"io/ioutil"
	"os"
)

func JsonFileToStruct(jsonFile string, c interface{}, notExist func() error) error {
	file, err := os.Open(jsonFile)
	if err != nil {
		if os.IsNotExist(err) {
			if notExist != nil {
				err := notExist()
				if err != nil {
					log.Error("notExist func execute err", err)
					return nil
				}
			}
			file, err := os.Create(jsonFile)
			if err != nil {
				log.Error("create file[%v] err:%v", jsonFile, err)
				return nil
			}
			defer file.Close()
			return nil
		}
		log.Error("open config file error:%v", err)
	}
	defer file.Close()
	all, err := ioutil.ReadAll(file)
	if err != nil {
		log.Error("read config file error:%v", err)
		return nil
	}
	if len(all) == 0 {
		return nil
	}
	err = json.Unmarshal(all, c)
	if err != nil {
		log.Error("parse config json error:%v", err)
		return nil
	}
	return nil
}
