package api

import "auto-apidoc/api/document"

func AddPath(name, identify string) int {
	d := document.Document{
		DocumentName: name,
		Identify:     identify,
	}
	return d.AddPath()
}

func AddDoc(name, identify, mdStr string, parentId int) int {
	d := document.Document{
		DocumentName: name,
		Identify:     identify,
		ParentId:     parentId,
	}
	docId := d.AddDoc()
	if docId == 0 {
		return 0
	}
	d.DocumentId = docId
	d.Markdown = mdStr
	return d.AddContentToDoc()
}
