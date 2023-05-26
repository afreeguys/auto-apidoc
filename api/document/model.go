package document

type MinDocResult struct {
	Data    map[string]interface{} `json:"data"`
	Errcode int                    `json:"errcode"`
	Message string                 `json:"message"`
}

type Document struct {
	DocumentId   int    `json:"doc_id"`
	DocumentName string `json:"doc_name"`
	// Identify 文档唯一标识
	Identify string `json:"identify"`
	BookId   int    `json:"book_id"`
	ParentId int    `json:"parent_id"`
	// Markdown markdown格式文档.
	Markdown string `json:"markdown"`
	// Release 发布后的Html格式内容.
	Release string `json:"release"`
	// Content 未发布的 Html 格式内容.
	Content string `json:"content"`
	//是否展开子目录：0 否/1 是 /2 空间节点，单击时展开下一级
	IsOpen int `json:"is_open"`
}
