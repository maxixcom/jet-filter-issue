package main

import (
	"github.com/CloudyKit/jet"
	"os"
)

type Result struct {
	id string
	data string
}

func main() {
	view := jet.NewHTMLSet("html")

	template,_:=view.GetTemplate("main")

	vars := make(jet.VarMap)
	vars.Set("result",Result {
		"id",
		"data",
	})

	template.Execute(os.Stdout,vars,nil)
}
