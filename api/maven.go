package api

import (
	"auto-apidoc/commom"
	"github.com/beevik/etree"
	"github.com/kataras/golog"
	"log"
	"os"
	"os/exec"
	"strings"
)

type pomResult struct {
	finalName  string //jar包文件名
	parentPath string //parent路径，如果不是聚合工程，则为空
	isModule   bool   //是否为聚合工程
}

//返回 target/**.jar
//mvnCmd 必须可以执行，如果存在环境变量path中，则只传mvn即可，否则需要mvn的绝对路径
//settingPath 执行的setting.xml绝对路径，为空则使用maven默认配置
func MavenBuild(mvnCmd, settingPath string) string {
	result, err := getBuildFinalNameFromPom("pom.xml")
	if err != nil {
		switch err {
		case commom.ErrIsModulesProject:
			log.Fatal("project is modules,please move to you want module")
		case commom.ErrPomNoBuild:
			log.Fatal(commom.ErrPomNoBuild)
		default:
			log.Fatal(err)
		}
	}
	artifactPath := "target" + commom.PathSeparator + result.finalName
	if result.isModule {
		log.Println("parent mvn clean...")
		err = exec.Command(mvnCmd, "-f", result.parentPath+commom.PathSeparator+"pom.xml", "clean").Run()
		if err != nil {
			log.Fatalf("parent mvn clean err:%v", err)
		}
		log.Println("parent mvn clean completed")
		log.Println("parent mvn package...")
		var cmdResult []byte
		if settingPath != "" {
			cmdResult, err = exec.Command(mvnCmd, "-f", result.parentPath+commom.PathSeparator+"pom.xml", "-DskipTests=true", "package", "--settings", settingPath).Output()
		} else {
			cmdResult, err = exec.Command(mvnCmd, "-f", result.parentPath+commom.PathSeparator+"pom.xml", "-DskipTests=true", "package").Output()
		}
		//err = exec.Command("bash","-c", fmt.Sprintf("mvn -f %v/pom.xml  -DskipTests=true package",result.parentPath)).Run()
		if err != nil {
			log.Println(string(cmdResult))
			log.Fatalf("parent mvn package err:%v", err)
		}
		log.Println("parent mvn package completed")
		_, err = os.Stat(artifactPath)
		if os.IsNotExist(err) {
			log.Fatalf("target not exist,please check the project. err:%v", err)
		}
		log.Printf("parent mvn result is [%v]\n", artifactPath)
		return artifactPath
	}
	log.Println("mvn clean...")
	err = exec.Command(mvnCmd, "clean").Run()
	if err != nil {
		log.Fatalf("mvn clean err:%v", err)
	}
	log.Println("mvn clean completed")
	log.Println("mvn package...")
	if settingPath != "" {
		err = exec.Command(mvnCmd, "-DskipTests=true", "package", "--settings", settingPath).Run()
	} else {
		err = exec.Command(mvnCmd, "-DskipTests=true", "package").Run()
	}
	if err != nil {
		log.Fatalf("mvn package err:%v", err)
	}
	log.Println("mvn package completed")
	_, err = os.Stat(artifactPath)
	if os.IsNotExist(err) {
		log.Fatalf("target not exist,please check the project. err:%v", err)
	}
	log.Printf("mvn result is [%v]\n", artifactPath)
	return artifactPath
}

//返回 **.jar
func getBuildFinalNameFromPom(pomPath string) (pomResult, error) {
	result := pomResult{}
	document := etree.NewDocument()
	err := document.ReadFromFile(pomPath)
	if err != nil {
		if os.IsNotExist(err) {
			log.Fatalf("pom.xml not exist ,please check the project")
		}
		return result, err
	}
	project := document.FindElement("project")
	modules := project.FindElement("modules")
	if modules != nil {
		return result, commom.ErrIsModulesProject
	}
	dir, err := os.Getwd()
	if err != nil {
		return result, err
	}
	//dir = "D:\\document\\workspace\\ads\\test"
	split := strings.Split(dir, commom.PathSeparator)
	// 获取parent
	parent := project.FindElement("parent")
	artifactId := parent.FindElement("artifactId")
	if split[len(split)-2] == artifactId.Text() {
		parentPath := ""
		tmpParent := dir
		for true {
			tmpParent = getParentPath(tmpParent)
			_, err := os.Stat(tmpParent + commom.PathSeparator + "pom.xml")
			if err == nil {
				parentPath = tmpParent
				continue
			}
			if os.IsNotExist(err) {
				break
			}
		}
		result.parentPath = parentPath
		result.isModule = true
	}
	packagingExt := "jar"
	element := project.FindElement("packaging")
	if element != nil {
		packagingExt = element.Text()
	}
	elements := project.FindElements("build")
	finalName := ""
	if len(elements) == 0 {
		log.Fatal("pom.xml have not <build>")
		return result, commom.ErrPomNoBuild
	}
	for _, element := range elements {
		path := etree.MustCompilePath("finalName")
		findElement := element.FindElementPath(path)
		if findElement != nil {
			finalName = findElement.Text()
			break
		}
	}
	result.finalName = finalName + "." + packagingExt
	return result, nil
}

//获取文件路径上一级
func getParentPath(path string) string {
	index := strings.LastIndex(strings.TrimRight(path, commom.PathSeparator), commom.PathSeparator)
	return path[:index]
}

// 验证maven 是否存在环境变量中
func CheckMavenExist(mvn string) bool {
	_, err := exec.Command(mvn, "-v").Output()
	if err != nil {
		return false
	}
	return true
}

//获取本地仓库地址
func GetLocalRepository(settingsPath string) string {
	if settingsPath == "" {
		envs := os.Getenv("path")
		split := strings.Split(envs, ";")
		for _, s := range split {
			if strings.Contains(strings.ToLower(s), "maven") {
				if strings.HasSuffix(s, string(os.PathSeparator)) {
					s = s[0 : len(s)-1]
				}
				sepatator := string(os.PathSeparator)
				index := strings.LastIndex(s, sepatator)
				settingsPath = s[0:index] + sepatator + "conf" + sepatator + "settings.xml"
				break
			}
		}
	}
	if settingsPath == "" {
		golog.Errorf("settings is blank")
		return ""
	}
	document := etree.NewDocument()
	err := document.ReadFromFile(settingsPath)
	if err != nil {
		golog.Errorf("settings[%v] parse err:%v", settingsPath, err)
		return ""
	}
	settingsElement := document.FindElement("settings")
	if settingsElement == nil {
		golog.Errorf("settings[%v] <settings> not found", settingsPath)
		return ""
	}
	element := settingsElement.FindElement("localRepository")
	if element == nil {
		golog.Errorf("settings[%v] <localRepository> not found", settingsPath)
		return ""
	}
	return element.Text()
}
