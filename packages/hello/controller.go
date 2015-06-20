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

func (c *HelloController) SayHelloTo(name string) string {
	return "Hello " + name
}

func (c *HelloController) Greeting() string {
	return "Greetings"
}
