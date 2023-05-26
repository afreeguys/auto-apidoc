package api

import (
	"encoding/json"
	"fmt"
	"strings"
)

type ControllerDto struct {
	Author       string        `json:"author"`
	Description  string        `json:"description"`
	BaseUrl      string        `json:"baseUrl"`
	ClassName    string        `json:"className"`
	PackageName  string        `json:"packageName"`
	RequestNodes []*RequestDto `json:"requestNodes"`
	SrcFileName  string        `json:"srcFileName"`
}

func (c *ControllerDto) ToMarkDownText() (string, error) {
	mdDoc := &MdDoc{}
	mdDoc.Title = c.Description
	if len(c.RequestNodes) > 0 {
		methods := make([]MdMethod, 0, len(c.RequestNodes))
		for _, node := range c.RequestNodes {
			method := MdMethod{
				Title:        node.Description,
				Type:         fmt.Sprintf("%v", node.Method),
				Url:          node.Url,
				ParamFields:  nil,
				ParamsJson:   "",
				ResultFields: nil,
				ResultJson:   node.ReturnDto.Description,
			}
			// 拼装参数
			paramFields := make([]MdMethodField, 0, 10)
			if len(node.RequestBodys) > 0 {
				method.IsRequestBody = true
				for _, body := range node.RequestBodys {
					if method.ParamsJson == "" {
						method.ParamsJson = body.Description
					}
					if len(body.Fields) == 0 {
						continue
					}
					for _, field := range body.Fields {
						field.AddMdParam(&paramFields)
					}
				}
			}
			if len(node.Params) > 0 {
				for _, param := range node.Params {
					paramFields = append(paramFields, MdMethodField{
						Name:        param.Name,
						Type:        param.TypeStr,
						Description: strings.ReplaceAll(param.Description, "\r", ""),
						IsRequired:  1,
					})
				}
			}
			method.ParamFields = paramFields

			//拼装返回值
			resultFields := make([]MdMethodField, 0, 10)
			for _, field := range node.ReturnDto.Fields {
				resultFields = append(resultFields, MdMethodField{
					Name:        field.Name,
					Type:        field.TypeStr,
					Description: strings.ReplaceAll(field.Description, "\r", ""),
					IsRequired:  1,
				})
			}
			method.ResultFields = resultFields
			method.ResultJson = node.ReturnDto.Description
			methods = append(methods, method)
		}
		mdDoc.Methods = methods
	}
	return mdDoc.ToMarkDownText()
}

type RequestDto struct {
	Method       []string     `json:"method"`
	Url          string       `json:"url"`
	MethodName   string       `json:"methodName"` //方法名
	Author       string       `json:"author"`
	Description  string       `json:"description"` //接口名
	Supplement   string       `json:"supplement"`  //补充说明，对应方法 @description
	Params       []*ParamDto  `json:"params"`
	RequestBodys []*ClassDto  `json:"requestBodys"`
	Header       []*HeaderDto `json:"header"`
	Deprecated   bool         `json:"deprecated"`
	ReturnDto    *ClassDto    `json:"returnDto"`
	MethodMd5    string       `json:"methodMd5"`
}

type ParamDto struct {
	Name        string `json:"name"`
	TypeStr     string `json:"type"`
	Required    bool   `json:"required"`
	Description string `json:"description"`
}

type ClassDto struct {
	ClassName    string        `json:"className"`
	ModelClass   string        `json:"modelClass"` //for reflection
	Description  string        `json:"description"`
	IsList       bool          `json:"isList"`
	Fields       []*FieldDto   `json:"fields"`
	GenericNodes []*GenericDto `json:"genericNodes"`
}

type FieldDto struct {
	Name        string    `json:"name"`
	TypeStr     string    `json:"type"`
	Description string    `json:"description"`
	ChildNode   *ClassDto `json:"childNode"` // 表示该field持有的对象类
	LoopNode    bool      `json:"loopNode"`  // 有循环引用的类
	NotNull     bool      `json:"notNull"`
}

// 将参数放到入参的数组中去，最后作为结果返回
func (f *FieldDto) AddMdParam(paramFields *[]MdMethodField) {
	*paramFields = append(*paramFields, MdMethodField{
		Name:        f.Name,
		Type:        f.TypeStr,
		Description: strings.ReplaceAll(f.Description, "\r", ""),
		IsRequired:  1,
	})
	if f.ChildNode == nil {
		return
	}
	if len(f.ChildNode.Fields) == 0 {
		return
	}
	for _, field := range f.ChildNode.Fields {
		field.addMdParamRecursion(paramFields, f.Name, nil)
	}
	return
}

// 将参数放到入参的数组中去，最后作为结果返回
func (f *FieldDto) addMdParamRecursion(paramFields *[]MdMethodField, parentName string, keyMap *map[string]bool) {
	if keyMap == nil {
		m := make(map[string]bool)
		keyMap = &m
	}
	if (*keyMap)[f.Name+f.TypeStr+f.Description] {
		return
	}
	*paramFields = append(*paramFields, MdMethodField{
		Name:        parentName + "." + f.Name,
		Type:        f.TypeStr,
		Description: strings.ReplaceAll(f.Description, "\r", ""),
		IsRequired:  1,
	})
	(*keyMap)[f.Name+f.TypeStr+f.Description] = true
	if f.ChildNode == nil {
		return
	}
	if len(f.ChildNode.Fields) == 0 {
		return
	}
	for _, field := range f.ChildNode.Fields {
		field.addMdParamRecursion(paramFields, f.Name, keyMap)
	}
	return
}

type HeaderDto struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

type GenericDto struct {
	ClassType        string        `json:"classType"`  // for source
	ModelClass       string        `json:"modelClass"` //for reflection
	Placeholder      string        `json:"placeholder"`
	ChildGenericNode []*GenericDto `json:"childGenericNode"`
}

func ParseJson(jsonStr string) ([]*ControllerDto, error) {
	dtos := make([]*ControllerDto, 0, 10)
	err := json.Unmarshal([]byte(jsonStr), &dtos)
	if err != nil {
		return nil, err
	}
	return dtos, nil
}
