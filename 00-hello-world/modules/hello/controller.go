package hello

import "bitbucket.org/carbonate/carbonate"

type HelloController struct {
	carbonate.BaseController

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

func (c *HelloController) Template() string {
	return c.view.RenderTemplateAsString("sample", nil)
}
