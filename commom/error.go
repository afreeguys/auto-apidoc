package commom

import "errors"

var (
	ErrIsModulesProject = errors.New("project is modules")
	ErrPomNoBuild       = errors.New("pom.xml have not <build>")
)
