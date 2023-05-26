package config

import (
	"auto-apidoc/commom"
	"auto-apidoc/utils"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

type ProjectConfig struct {
	Name            string `json:"name"`
	OutDir          string `json:"outDir"`
	Version         int    `json:"version"`
	ProjectDir      string `json:"projectDir"`
	MavenRepository string `json:"mavenRepository"`
}

func (pc *ProjectConfig) Save() {
	jsonFile := HomeDir.ConfigPath() + pc.Name + ".json"
	_ = os.Remove(jsonFile)
	bytes, _ := json.Marshal(pc)
	_ = os.WriteFile(jsonFile, bytes, 0777)
}

// 接口文档地址，包含版本
func (pc *ProjectConfig) DocPath() string {
	return pc.OutDir + commom.PathSeparator + pc.Name + commom.PathSeparator + fmt.Sprintf("V%v", pc.Version)
}

func InitProjectConfig(projectName string) *ProjectConfig {
	if projectName == "" {
		log.Fatal("projectName cannot be blank")
	}
	c := &ProjectConfig{}
	jsonFile := HomeDir.ConfigPath() + projectName + ".json"
	file, err := os.Open(jsonFile)
	if err != nil {
		if os.IsNotExist(err) {
			_, err := os.Create(jsonFile)
			if err != nil {
				log.Fatalf("create file[%v] err:%v", jsonFile, err)
			}
			return nil
		}
		log.Fatalf("open config file error:%v", err)
	}
	defer file.Close()
	all, err := ioutil.ReadAll(file)
	if err != nil {
		log.Fatalf("read config file error:%v", err)
	}
	if len(all) == 0 {
		return nil
	}
	err = json.Unmarshal(all, c)
	if err != nil {
		log.Fatalf("parse config json error:%v", err)
	}
	return c
}

func CheckProjectName(projectName, projectDir string) string {
	if projectName == "" {
		log.Fatal("projectName cannot be blank")
	}
	pc := &ProjectConfig{}
	jsonFile := HomeDir.ConfigPath() + projectName + ".json"
	err := utils.JsonFileToStruct(jsonFile, pc, nil)
	if err != nil {
		log.Fatal(err)
	}
	if pc.ProjectDir == projectDir {
		return pc.Name
	}
	projectName = projectName + "new"
	for true {
		pc = &ProjectConfig{}
		jsonFile := HomeDir.ConfigPath() + projectName + ".json"
		err := utils.JsonFileToStruct(jsonFile, pc, nil)
		if err != nil {
			log.Fatal(err)
		}
		if pc.Name == "" {
			return projectName
		}
		if pc.ProjectDir == projectDir {
			return projectName
		}
		projectName = projectName + "new"
	}
	return ""
}
