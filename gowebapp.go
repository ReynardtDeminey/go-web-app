package main

import createwebapp "github.com/ReynardtDeminey/gowebapp/go-build-web-app"

func main() {
	project := "new-go-web-app"

	createwebapp.CreateFolder(project)
	createwebapp.CreateHTML(project)
	createwebapp.CreateCSS(project)
	createwebapp.CreateServer(project)
}
