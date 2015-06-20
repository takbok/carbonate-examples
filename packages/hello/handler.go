package hello

import (
	"net/http"

	"bitbucket.org/takbok/brahma"
	"bitbucket.org/takbok/brahma/module"
)

func init() {
	module.DefaultMethod("index")
	module.RegisterRequestHandler(serve)
	module.SetAsDefault()

	module.ForwardPermalinkToMethod("/greet", "greeting")
}

func serve(w http.ResponseWriter, r *http.Request) {
	var c HelloController
	c.SetModelAndView(&c.model, &c.view)

	brahma.DispatchRequestViaURL(r.URL.Path, &c, &w, &r)
}
