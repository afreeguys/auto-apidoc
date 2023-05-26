package resource

import _ "embed"

//go:embed japiTools.jar
var JarExecuteFile []byte

//go:embed controller_template.md.template
var RequestMethodTemplate string
