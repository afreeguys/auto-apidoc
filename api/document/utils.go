package document

import (
	"encoding/json"
	"fmt"
	"github.com/gomarkdown/markdown"
	"io"
)

func MapToModel(content interface{}, dest interface{}) {
	if content == nil {
		fmt.Printf("content is empty\n")
		return
	}
	if dest == nil {
		fmt.Printf("dest is nil\n")
		return
	}
	marshal, err := json.Marshal(content)
	if err != nil {
		fmt.Printf("%v\n", err)
		return
	}
	err = json.Unmarshal(marshal, dest)
	if err != nil {
		fmt.Printf("%v\n", err)
		return
	}
}

func JsonToMap(content []byte) map[string]interface{} {
	m := make(map[string]interface{})
	err := json.Unmarshal(content, &m)
	if err != nil {
		fmt.Printf("JsonToMap error:%v", err)
		return nil
	}
	return m
}

func ReaderToMap(reader io.Reader) map[string]interface{} {
	m := make(map[string]interface{})
	r := json.NewDecoder(reader)
	err := r.Decode(&m)
	if err != nil {
		fmt.Printf("ReaderToMap error:%v\n", err)
		return nil
	}
	return m
}

func ReaderToI(reader io.Reader, dest interface{}) {
	r := json.NewDecoder(reader)
	err := r.Decode(dest)
	if err != nil {
		fmt.Printf("ReaderToMap error:%v\n", err)
	}
}

func MarkdownToHtml(markdownStr string) string {
	contents := markdown.NormalizeNewlines([]byte(markdownStr))
	return string(markdown.ToHTML(contents, nil, nil))
}
