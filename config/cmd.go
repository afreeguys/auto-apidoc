package config

import (
	"auto-apidoc/commom"
	"auto-apidoc/utils"
	_ "embed"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
	"os/user"
)

var (
	HomeDir homeDirS
)

type homeDirS string

func (h homeDirS) ConfigPath() string {
	return string(h) + commom.PathSeparator + "config" + commom.PathSeparator
}

var globalConfig *GlobalConfig

func (h homeDirS) GlobalConfig() GlobalConfig {
	if globalConfig != nil {
		return *globalConfig
	}
	gfile := h.ConfigPath() + "global.json"
	if !utils.CheckPath(gfile) {
		return GlobalConfig{}
	}
	file, err := os.Open(gfile)
	if err != nil {
		log.Fatal(fmt.Sprintf("open file[%v] err:%v", gfile, err))
	}
	defer file.Close()
	config := &GlobalConfig{}
	all, err := io.ReadAll(file)
	if err != nil {
		log.Fatal(fmt.Sprintf("read all file[%v] err:%v", gfile, err))
	}
	err = json.Unmarshal(all, config)
	if err != nil {
		log.Fatal(fmt.Sprintf("file[%v] to json err:%v", gfile, err))
	}
	globalConfig = config
	return *config
}

type GlobalConfig struct {
	Maven          *MavenConfig `json:"maven"` //maven 配置
	DocumentOutDir string       //文档导出路径
}

type MavenConfig struct {
	MvnPath         string `json:"mvnPath"`
	SettingPath     string `json:"settingPath"`
	LocalRepository string `json:"localRepository"`
}

func (g GlobalConfig) Save() error {
	filePath := HomeDir.ConfigPath() + "global.json"
	_ = os.Remove(filePath)
	file, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer file.Close()
	marshal, err := json.Marshal(g)
	if err != nil {
		return err
	}
	err = os.WriteFile(filePath, marshal, 0777)
	if err != nil {
		return err
	}
	return nil
}

func init() {
	current, err := user.Current()
	if err != nil {
		log.Fatalf("get current user err:%v", err)
	}
	HomeDir = homeDirS(current.HomeDir + commom.PathSeparator + commom.ConfigPath)
}

func (pc ProjectConfig) GetJsonFile() *os.File {
	file, err := os.Open(pc.OutDir + commom.PathSeparator + pc.Name + commom.PathSeparator + fmt.Sprintf("V%v", pc.Version) + commom.PathSeparator + commom.JsonFileName)
	if err != nil {
		log.Fatal(err)
	}
	return file
}

func (pc ProjectConfig) GetJsonFileForLast() *os.File {
	file, err := os.Open(pc.OutDir + commom.PathSeparator + pc.Name + commom.PathSeparator + fmt.Sprintf("V%v", pc.Version-1) + commom.PathSeparator + commom.JsonFileName)
	if err != nil {
		fmt.Printf("GetJsonFileForLast last version is [%v] err:%v\n", pc.Version-1, err)
	}
	return file
}

// 获取指定版本数据
func (pc ProjectConfig) GetJsonFileForVersion(version int) *os.File {
	file, err := os.Open(pc.OutDir + commom.PathSeparator + pc.Name + commom.PathSeparator + fmt.Sprintf("V%v", version) + commom.PathSeparator + commom.JsonFileName)
	if err != nil {
		fmt.Printf("GetJsonFileForLast last version is [%v] err:%v\n", version, err)
	}
	return file
}

func (pc *ProjectConfig) CleanLibDir() {
	path := pc.OutDir + commom.PathSeparator + pc.Name + commom.PathSeparator + "lib"
	stat, _ := os.Stat(path)
	if stat != nil {
		err := os.RemoveAll(path)
		if err != nil {
			log.Fatalf("clean lib dir[%v] err:%v", path, err)
		}
	}
}

func (pc *ProjectConfig) SaveFile(fileName string, content []byte) error {
	path := pc.OutDir + commom.PathSeparator + pc.Name + commom.PathSeparator + fmt.Sprintf("V%v", pc.Version) + commom.PathSeparator + "markdown"
	_, err := os.Stat(path)
	if err != nil {
		if os.IsNotExist(err) {
			err := os.MkdirAll(path, 0777)
			if err != nil {
				return err
			}
		}
	}
	fullName := path + commom.PathSeparator + fileName
	return os.WriteFile(fullName, content, 0777)
}
