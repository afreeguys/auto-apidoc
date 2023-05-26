package document

import (
	"bytes"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
	"time"
)

func (doc *Document) AddPath() int {
	form := &url.Values{}
	form.Add("identify", "api_doc")
	form.Add("doc_id", "0")
	form.Add("doc_name", doc.DocumentName)
	form.Add("doc_identify", doc.Identify)
	form.Add("is_open", "2")
	post, err := postToServer("http://localhost:8181/api/api_doc/create", form.Encode())
	if err != nil {
		fmt.Println(err)
	}
	minDocResult := &MinDocResult{}
	ReaderToI(post.Body, minDocResult)
	if minDocResult.Errcode == 0 || minDocResult.Errcode == 6006 {
		d := &Document{}
		MapToModel(minDocResult.Data, d)
		if d.DocumentId > 0 {
			return d.DocumentId
		}
	}
	fmt.Printf("%v\n", minDocResult)
	return 0
}

func (doc *Document) AddDoc() int {
	if doc.ParentId == 0 {
		fmt.Println("AddDoc parentId cannot be 0")
		return 0
	}
	form := &url.Values{}
	form.Add("identify", "api_doc")
	form.Add("doc_id", "0")
	form.Add("doc_name", doc.DocumentName)
	form.Add("doc_identify", doc.Identify)
	form.Add("parent_id", strconv.Itoa(doc.ParentId))
	form.Add("is_open", "1")
	post, err := postToServer("http://localhost:8181/api/api_doc/create", form.Encode())
	if err != nil {
		fmt.Println(err)
	}
	minDocResult := &MinDocResult{}
	ReaderToI(post.Body, minDocResult)
	if minDocResult.Errcode == 0 || minDocResult.Errcode == 6006 {
		d := &Document{}
		MapToModel(minDocResult.Data, d)
		if d.DocumentId > 0 {
			return d.DocumentId
		}
	}
	fmt.Printf("%v\n", minDocResult)
	return 0
}

func (doc *Document) AddContentToDoc() int {
	if doc.DocumentId == 0 {
		fmt.Println("AddContentToDoc DocumentId cannot be 0")
		return 0
	}
	form := &url.Values{}
	form.Add("identify", "api_doc")
	form.Add("doc_id", strconv.Itoa(doc.DocumentId))
	form.Add("markdown", doc.Markdown)
	form.Add("html", MarkdownToHtml(doc.Markdown))
	form.Add("cover", "yes")
	form.Add("version", strconv.FormatInt(time.Now().Unix(), 10))
	post, err := postToServer("http://localhost:8181/api/api_doc/content/", form.Encode())
	if err != nil {
		fmt.Println(err)
	}
	minDocResult := &MinDocResult{}
	ReaderToI(post.Body, minDocResult)
	if minDocResult.Errcode == 0 || minDocResult.Errcode == 6006 {
		d := &Document{}
		MapToModel(minDocResult.Data, d)
		if d.DocumentId > 0 {
			return d.DocumentId
		}
	}
	fmt.Printf("%v\n", minDocResult)
	return 0
}

func postToServer(postUrl, form string) (*http.Response, error) {
	buffer := &bytes.Buffer{}
	buffer.WriteString(form)
	request, err := http.NewRequest(http.MethodPost, postUrl, buffer)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	request.Header = map[string][]string{
		"Content-Type": []string{"application/x-www-form-urlencoded; charset=UTF-8"},
		"Refer":        []string{"autoJapi"},
	}
	return http.DefaultClient.Do(request)
}
