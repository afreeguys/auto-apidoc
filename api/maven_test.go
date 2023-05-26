package api

import (
	"auto-apidoc/commom"
	"fmt"
	"log"
	"strings"
	"testing"
)

func TestMavenBuild1(t *testing.T) {
	//result, err := getBuildFinalNameFromPom("D:\\document\\workspace\\ads\\test\\pom.xml")
	result, err := getBuildFinalNameFromPom("D:\\document\\workspace\\test\\pom.xml")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(result)
}

func TestCheckMaven(*testing.T) {
	fmt.Println(CheckMavenExist("mvn"))
	fmt.Println(CheckMavenExist("D:\\software\\apache-maven-3.6.3\\bin\\mvn.cmd"))
}

func Test_getParentPath(*testing.T) {
	fmt.Println(getParentPath("D:\\software\\apache-maven-3.6.3\\bin\\mvn.cmd"))
	fmt.Println(getParentPath("D:\\software\\apache-maven-3.6.3\\bin\\"))
	fmt.Println(getParentPath("D:\\software\\apache-maven-3.6.3\\bin"))
}

func TestCheckMaven1(*testing.T) {
	mvncmd := "D:/software/apache-maven-3.6.3"
	mvncmd = strings.ReplaceAll(mvncmd, "/", commom.PathSeparator)
	fmt.Println(mvncmd)
}

func TestCheckMaven3(*testing.T) {
	var content string
	_, err := fmt.Scan(&content)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(content)
}

func TestGetLocalRepository(*testing.T) {
	fmt.Println(GetLocalRepository(""))
	fmt.Println(GetLocalRepository("D:\\software\\apache-maven-3.6.3\\conf\\settings - 副本.xml"))
}
