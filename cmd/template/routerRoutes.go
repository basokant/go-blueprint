package template

import (
	_ "embed"
)

//go:embed files/routes/http_router.go.tmpl
var httpRouterRoutesTemplate []byte

// RouterTemplates contains the methods used for building
// an app that uses [github.com/julienschmidt/httprouter]
type RouterTemplates struct{}

func (r RouterTemplates) Main() []byte {
	return mainTemplate
}
func (r RouterTemplates) Server() []byte {
	return standardServerTemplate
}
func (r RouterTemplates) Routes() []byte {
	return httpRouterRoutesTemplate
}
