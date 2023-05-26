package api

import (
	"auto-apidoc/resource"
	"bytes"
	"encoding/json"
	"fmt"
	"strings"
	"text/template"
)

// 文档属性
type MdDoc struct {
	Title   string
	Methods []MdMethod
}

// 请求方法属性
type MdMethod struct {
	Title         string // 标题
	Type          string // http method
	Url           string
	IsRequestBody bool            //是否以requestBody 形式
	ParamFields   []MdMethodField // 请求参数列表
	ParamsJson    string          //请求参数的json格式伪代码，只有请求时requestBody时才有
	ResultFields  []MdMethodField //返回数据列表
	ResultJson    string          //请求参数的json格式伪代码
}

func (m *MdMethod) GenerateTestDataFromJson() string {
	if m.ParamsJson == "" {
		return ""
	}
	jsonTemp := make(map[string]string)
	err := json.Unmarshal([]byte(m.ParamsJson), &jsonTemp)
	if err != nil {
		fmt.Printf("method:%v generate test json data failed,error:%v\n", m.Title, err)
		return ""
	}
	data := make(map[string]interface{})
	for k, v := range jsonTemp {
		split := strings.Split(v, " ")
		data[k] = convertTypeToData(k, split[0])

	}

	marshal, err := json.Marshal(data)
	if err != nil {
		fmt.Printf("method:%v json data to string failed,error:%v\n", m.Title, err)
		return ""
	}
	buffer := &bytes.Buffer{}
	err = json.Indent(buffer, marshal, "", "\t")
	if err != nil {
		fmt.Printf("method:%v json data format failed,error:%v\n", m.Title, err)
		return string(marshal)
	}
	return buffer.String()
}

func convertTypeToData(key, typeStr string) interface{} {
	if key == "" || typeStr == "" {
		return nil
	}
	key = strings.ToLower(key)
	switch strings.ToLower(typeStr) {
	case "int":
		if key == "id" {
			return 1
		}
		return 0
	case "string":
		if strings.Contains(key, "url") {
			return "https://www.baidu.com"
		}
		return key + "_测试数据"
	case "long":
		return int64(0)
	default:
		return nil
	}
}

// 参数字段属性
type MdMethodField struct {
	Name        string //名称
	Type        string //字段类型
	Description string //字段描述
	IsRequired  int    //是否必输
}

func (c *MdDoc) ToMarkDownText() (string, error) {
	masterTmpl, err := template.New("master").Parse(resource.RequestMethodTemplate)
	if err != nil {
		return "", err
	}
	buffer := bytes.Buffer{}
	err = masterTmpl.Execute(&buffer, c)
	if err != nil {
		return "", err
	}
	return buffer.String(), nil
}
