package hello

import "bitbucket.org/takbok/brahma"

type HelloController struct {
	brahma.BaseController

	model HelloModel
	view  HelloView
}

func (c *HelloController) Index() string {
	return "Hello World"
}
