package document

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"testing"
)

func TestDocument_AddPath(t *testing.T) {
	document := Document{
		DocumentName: "烦烦烦烦烦烦",
		Identify:     "test",
		ParentId:     0,
	}
	fmt.Println(document.AddPath())
}

func TestDocument_AddDoc(t *testing.T) {
	document := Document{
		DocumentName: "烦烦烦烦烦烦",
		Identify:     "test111",
		ParentId:     20,
	}
	fmt.Println(document.AddDoc())
}

func TestDocument_AddContentToDoc(t *testing.T) {
	document := Document{
		DocumentId: 32,
		Markdown:   "jifjeiejfijef ",
		Content:    "<p class=\"line\">jifjeiejfijef </p>",
	}
	fmt.Println(document.AddContentToDoc())
}

func TestMapToModel(t *testing.T) {
	file, err := os.Open("D:\\opt\\temp\\japi\\test\\V22\\markdown\\CrawlingController.md")
	if err != nil {
		log.Fatal(err)
	}
	all, _ := ioutil.ReadAll(file)
	html := MarkdownToHtml(string(all))
	fmt.Println(html)
}
