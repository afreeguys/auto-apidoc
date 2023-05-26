package cmd

import (
	"auto-apidoc/api"
	"auto-apidoc/commom"
	"auto-apidoc/config"
	slog "auto-apidoc/log"
	"auto-apidoc/resource"
	"auto-apidoc/utils"
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"io/fs"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"
)

var (
	getAllDoc      bool
	fullJarPath    = ""
	fullConfigPath = ""
)

type cmdOption struct {
	SyncRemote    bool //是否同步远程
	GetRequireDoc bool //是否以自定义需求输出
	CompareOld    bool //是否对比历史
	OldVersion    int  //指定对比的历史版本号
}

func Run(agrvs []string) {
	option := cmdOption{}
	if len(agrvs) > 0 {
		for _, c := range agrvs {
			switch {
			//case strings.Contains(c, "-s"):
			//	option.SyncRemote = true
			case strings.Contains(c, "-r"):
				option.GetRequireDoc = true
			case strings.Contains(c, "-c"):
				option.CompareOld = true
				if len(c) > 2 {
					s := c[strings.Index(c, "-c")+2:]
					v, err := strconv.Atoi(s)
					if err != nil {
						fmt.Println("参数错误：-c 只能指定数字，例如：-c12：对比版本为12的接口数据")
						return
					}
					option.OldVersion = v
				}
			case strings.Contains(c, "-h"):
				useage()
				return
			}
		}
	}
	run(option)
}

func useage() {
	fmt.Println("使用方式：autojapi [option]")
	fmt.Println("option 不是必输，具体内容如下：")
	//fmt.Println("-s 输出文档同步远程客户端")
	fmt.Println("-c[version] 对比历史版本，只输出差异的接口；version为空对比上一版本，version不为空对比version版本")
	fmt.Println("-r 按照自定义需求输出文档")
	fmt.Println("-h 帮助文档")
}

func run(option cmdOption) {
	checkJarFile()
	project := InitProject()
	err := RunJava(project)
	if err != nil {
		log.Fatalf("run jar file err:%v", err)
	}
	defer func() {
		fmt.Println("------------------------------------------------------------------------")
		fmt.Println("clean lib directory...")
		project.CleanLibDir()
		fmt.Println("clean lib complete")
	}()
	fmt.Println("------------------------------------------------------------------------")
	file := project.GetJsonFile()
	defer file.Close()
	jsonModels, err := JsonToModel(file)
	if err != nil {
		log.Fatal(err)
	}
	if option.CompareOld {
		jsonModels = removeNoDiffMethod(project, jsonModels, option)
		if len(jsonModels) == 0 {
			fmt.Println("接口无变化，请去除-c使用")
			return
		}
	}
	if option.GetRequireDoc {
		getPRDMdDoc(project, jsonModels, option)
	} else {
		getControllerMdDoc(project, jsonModels, option)
	}

	fmt.Println("------------------------------------------------------------------------")
	fmt.Println("转换成功")
	fmt.Println("------------------------------------------------------------------------")
	fmt.Printf("目标文件地址:%v\n", project.DocPath())
	_ = utils.ExecCmdWithStdOut("explorer", project.DocPath())
}

func removeNoDiffMethod(pc config.ProjectConfig, list []*api.ControllerDto, option cmdOption) []*api.ControllerDto {
	var oldJsonFile *os.File
	if option.OldVersion == 0 {
		oldJsonFile = pc.GetJsonFileForLast()
	} else {
		oldJsonFile = pc.GetJsonFileForVersion(option.OldVersion)
	}
	if oldJsonFile == nil {
		slog.Error("open last json File err")
		return list
	}
	oldList, err := JsonToModel(oldJsonFile)
	if err != nil {
		slog.Error("read file[%v] to json err:%v", oldJsonFile.Name(), err)
		return list
	}
	if len(oldList) == 0 {
		return list
	}
	oldMethodMap := make(map[string]*api.RequestDto)
	for _, dto := range oldList {
		if len(dto.RequestNodes) == 0 {
			continue
		}
		for _, r := range dto.RequestNodes {
			oldMethodMap[dto.ClassName+":"+r.MethodName+":"+r.Url] = r
		}
	}
	result := make([]*api.ControllerDto, 0, 10)
	for _, dto := range list {
		if len(dto.RequestNodes) == 0 {
			continue
		}
		rList := make([]*api.RequestDto, 0, 10)
		for _, r := range dto.RequestNodes {
			requestDto, ok := oldMethodMap[dto.ClassName+":"+r.MethodName+":"+r.Url]
			if !ok {
				rList = append(rList, r)
				continue
			}
			if r.MethodMd5 == requestDto.MethodMd5 {
				continue
			}
			rList = append(rList, r)
		}
		if len(rList) == 0 {
			continue
		}
		result = append(result, &api.ControllerDto{
			Author:       dto.Author,
			Description:  dto.Description,
			BaseUrl:      dto.BaseUrl,
			ClassName:    dto.ClassName,
			PackageName:  dto.PackageName,
			RequestNodes: rList,
			SrcFileName:  dto.SrcFileName,
		})
	}
	if len(result) == 0 {
		slog.Info("json is equal with last vesion")
		return result
	}
	return result
}

//以controller维度输出接口文档
func getControllerMdDoc(project config.ProjectConfig, list []*api.ControllerDto, option cmdOption) {
	var pathId int
	if option.SyncRemote {
		pathId = api.AddPath(project.Name, project.Name)
		if pathId == 0 {
			log.Fatal("添加文档目录失败")
		}
	}
	for _, c := range list {
		text, err := c.ToMarkDownText()
		if err != nil {
			log.Fatal(fmt.Sprintf("%v to markdown err:%v", c.ClassName, err))
		}
		if option.SyncRemote {
			docId := api.AddDoc(c.ClassName, project.Name+"_"+c.ClassName, text, pathId)
			if docId == 0 {
				log.Fatal("添加文档内容失败")
			}
		}
		err = project.SaveFile(c.ClassName+".md", []byte(text))
		if err != nil {
			log.Fatal(fmt.Sprintf("%v saveFile err:%v", c.ClassName, err))
		}
	}
}

//以需求维度输出接口文档
func getPRDMdDoc(project config.ProjectConfig, list []*api.ControllerDto, option cmdOption) {
	var pathId int
	if option.SyncRemote {
		pathId = api.AddPath(project.Name, project.Name)
		if pathId == 0 {
			log.Fatal("添加文档目录失败")
		}
	}
	fmt.Println("请输入需求名称：")
	var docName string
	_, err := fmt.Scanln(&docName)
	if err != nil {
		log.Fatal("输入内容错误")
	}
	docName = fmt.Sprintf("[%v]-%v-接口文档", time.Now().Format("2006.01.02"), docName)
	methodMap := make(map[string]*api.RequestDto)
	controllerList := make([]string, 0, 10)
	controllerMethodNameMap := make(map[string][]string)
	controllerIndex := 0
	for _, c := range list {
		if len(c.RequestNodes) == 0 {
			continue
		}
		controllerIndex++
		controllerList = append(controllerList, fmt.Sprintf("%v %v", controllerIndex, c.ClassName))
		methodList := make([]string, 0, 10)
		methodIndex := 0
		for _, dto := range c.RequestNodes {
			methodIndex++
			methodList = append(methodList, fmt.Sprintf("%v.%v %v(%v)", controllerIndex, methodIndex, dto.Description, dto.MethodName))
			methodMap[strconv.Itoa(controllerIndex)+"."+strconv.Itoa(methodIndex)] = dto
		}
		controllerMethodNameMap[strconv.Itoa(controllerIndex)] = methodList
	}
	fmt.Println("===========================")
	fmt.Println("controller 如下：")
	for _, s := range controllerList {
		fmt.Println(s)
	}
	fmt.Println("===========================")
	fmt.Print("请输入你要选择的controller(使用/分割)：")
	var cIndexInput string
	for cIndexInput == "" {
		_, err = fmt.Scanln(&cIndexInput)
		if err != nil {
			log.Fatal("输入内容错误")
		}
	}
	fmt.Println(cIndexInput + " 中包含的接口如下：")
	conIndexList := strings.Split(cIndexInput, "/")
	for _, ci := range conIndexList {
		methodList, ok := controllerMethodNameMap[ci]
		if !ok {
			fmt.Printf("序号：%v 找不到对应的controller\n", ci)
		}
		for _, m := range methodList {
			fmt.Println(m)
		}
	}
	fmt.Println("===========================")
	fmt.Print("请输入你要选择的接口(使用/分割)，全选输入a：")
	var methodInput string
	for methodInput == "" {
		_, err = fmt.Scanln(&methodInput)
		if err != nil {
			log.Fatal("输入内容错误")
		}
	}

	dtos := make([]*api.RequestDto, 0, 10)
	if methodInput == "a" {
		for _, ci := range conIndexList {
			methodList, ok := controllerMethodNameMap[ci]
			if !ok {
				fmt.Printf("序号：%v 找不到对应的controller\n", ci)
			}
			for _, m := range methodList {
				d, ok := methodMap[m[0:strings.Index(m, " ")]]
				if !ok {
					continue
				}
				dtos = append(dtos, d)
			}
		}
	} else {
		split := strings.Split(methodInput, "/")
		for _, mi := range split {
			d, ok := methodMap[mi]
			if !ok {
				continue
			}
			dtos = append(dtos, d)
		}
	}
	c := api.ControllerDto{
		Description:  docName,
		ClassName:    docName,
		RequestNodes: dtos,
	}
	text, err := c.ToMarkDownText()
	if err != nil {
		log.Fatal(fmt.Sprintf("%v to markdown err:%v", c.ClassName, err))
	}
	if option.SyncRemote {
		docId := api.AddDoc(c.ClassName, project.Name+"_"+c.ClassName, text, pathId)
		if docId == 0 {
			log.Fatal("添加文档内容失败")
		}
	}
	err = project.SaveFile(c.ClassName+".md", []byte(text))
	if err != nil {
		log.Fatal(fmt.Sprintf("%v saveFile err:%v", c.ClassName, err))
	}

}

//配置项目
func InitProject() config.ProjectConfig {
	wd, err := os.Getwd()
	if err != nil {
		log.Fatalf("open dir err:%v", err)
	}
	split := strings.Split(wd, commom.PathSeparator)
	projectName := split[len(split)-1]

	g := config.HomeDir.GlobalConfig()
	projectName = config.CheckProjectName(projectName, wd)
	proConfig := config.InitProjectConfig(projectName)
	if proConfig == nil {
		proConfig = &config.ProjectConfig{Name: projectName, ProjectDir: wd}
		if g.DocumentOutDir == "" {
			fmt.Println("请输入导出路径:")
			var outPutDir string
			_, err = fmt.Scanln(&outPutDir)
			if err != nil {
				log.Fatalf("input err:%v", err)
			}
			if outPutDir == "" {
				outPutDir = string(config.HomeDir)
			}
			g.DocumentOutDir = outPutDir
			proConfig.OutDir = outPutDir
			_ = g.Save()
		} else {
			proConfig.OutDir = g.DocumentOutDir
		}
	}
	mvnCmd := "mvn"
	settingsPath := ""
	if !api.CheckMavenExist(mvnCmd) {
		if g.Maven == nil {
			stdinR := bufio.NewReader(os.Stdin)
			fmt.Println("请输入maven home路径:")
			mvnBinB, _, err := stdinR.ReadLine()
			if err != nil {
				log.Fatalf("input maven home err:%v", err)
			}
			mvnBin := string(mvnBinB)
			mvnBin = strings.ReplaceAll(mvnBin, "/", commom.PathSeparator)
			fmt.Println("请输入maven setting.xml路径:")
			settingB, _, err := stdinR.ReadLine()
			if err != nil {
				log.Fatalf("input maven setting.xml err:%v", err)
			}
			settingPath := string(settingB)
			_ = filepath.WalkDir(mvnBin, func(path string, d fs.DirEntry, err error) error {
				if !strings.Contains(path, "bin") {
					return nil
				}
				if d.IsDir() {
					return nil
				}
				if strings.Contains(path, "mvn.cmd") || strings.Contains(path, "mvn.bat") {
					mvnBin = path
					return filepath.SkipDir
				}
				return nil
			})
			//if strings.HasSuffix(mvnBin, commom.PathSeparator) {
			//	mvnBin = mvnBin + "bin" + commom.PathSeparator + "mvn.cmd"
			//} else {
			//	mvnBin = mvnBin + commom.PathSeparator + "bin" + commom.PathSeparator + "mvn.cmd"
			//}
			g.Maven = &config.MavenConfig{
				MvnPath:     mvnBin,
				SettingPath: settingPath,
			}
			err = g.Save()
			if err != nil {
				log.Fatal(fmt.Sprintf("save global.json err:%v\n", err))
			}
			mvnCmd = g.Maven.MvnPath
			settingsPath = g.Maven.SettingPath
		} else {
			mvnCmd = g.Maven.MvnPath
			settingsPath = g.Maven.SettingPath
		}
	}
	repository := api.GetLocalRepository(settingsPath)
	if repository == "" {
		log.Fatal(fmt.Sprintf("GetLocalRepository result is blank\n"))
	}
	proConfig.Version++
	proConfig.MavenRepository = repository
	proConfig.Save()
	return *proConfig
}

func RunJava(config config.ProjectConfig) error {
	err := utils.ExecCmdWithStdOut("java", "-jar", fullJarPath, config.ProjectDir, config.Name,
		fmt.Sprintf("V%v", config.Version), config.OutDir+commom.PathSeparator+config.Name, config.MavenRepository)
	if err != nil {
		return err
	}
	return nil
}

//验证初始文件夹 返回jar的全路径
func checkJarFile() string {
	fileName := string(config.HomeDir)
	_, err := os.Stat(fileName)
	if err != nil {
		if os.IsNotExist(err) {
			err := os.MkdirAll(fileName, 0777)
			if err != nil {
				log.Fatalf("mkdir err:%v", err)
			}
			err = os.MkdirAll(config.HomeDir.ConfigPath(), 0777)
			if err != nil {
				_ = os.RemoveAll(fileName)
				log.Fatalf("mkdir config dir err:%v", err)
			}
			fullJarPath = fileName + commom.PathSeparator + commom.JarFileName
			err = os.WriteFile(fullJarPath, resource.JarExecuteFile, 0777)
			if err != nil {
				_ = os.RemoveAll(fileName)
				log.Fatalf("copy jar file err:%v", err)
			}
			fullConfigPath = fileName + commom.PathSeparator + commom.ConfigFileName
			configFile, err := os.Create(fullConfigPath)
			if err != nil {
				_ = os.RemoveAll(fileName)
				log.Fatalf("copy jar file err:%v", err)
			}
			defer configFile.Close()

			return fileName + commom.PathSeparator + commom.JarFileName
		}
		log.Fatalf("check jar file exist err:%v", err)
	}
	fullJarPath = fileName + commom.PathSeparator + commom.JarFileName
	info, err := os.Stat(fullJarPath)
	if err != nil {
		if os.IsNotExist(err) {
			err = os.WriteFile(fullJarPath, resource.JarExecuteFile, 0777)
			if err != nil {
				log.Fatalf("copy new jar file err:%v", err)
			}
		}
	}
	if info != nil && info.Size() != int64(len(resource.JarExecuteFile)) {
		err := os.Remove(fullJarPath)
		if err != nil {
			log.Fatalf("remove old jar[%v] error:%v", fullJarPath, err)
		}
		err = os.WriteFile(fullJarPath, resource.JarExecuteFile, 0777)
		if err != nil {
			log.Fatalf("copy new jar file err:%v", err)
		}
	}
	return fileName + commom.PathSeparator + commom.JarFileName
}

func JsonToModel(input io.Reader) ([]*api.ControllerDto, error) {
	//utf8Reader := transform.NewReader(input, simplifiedchinese.GBK.NewDecoder())
	utf8Reader := input
	all, err := io.ReadAll(utf8Reader)
	if err != nil {
		return nil, err
	}
	dtos := make([]*api.ControllerDto, 0, 10)
	err = json.Unmarshal(all, &dtos)
	if err != nil {
		return nil, err
	}
	return dtos, nil
}
