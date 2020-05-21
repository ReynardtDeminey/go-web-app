package createwebapp

import (
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

func CreateHTML(project string) {
	str := fmt.Sprint(`
<!DOCTYPE html>
<html lang="en">
<head>
<meta charset="UTF-8">
<title class="heading">` + project + `</title>
<link rel="stylesheet" href="main.css">
</head>
<body>
<h1>Welcome to your Go Web App!</h1>
</body>
</html>
	`)

	nf, err := os.Create(project + "/templates/index.gohtml")
	if err != nil {
		log.Fatal("error creating file", err)
	}
	defer nf.Close()

	io.Copy(nf, strings.NewReader(str))
}

func CreateCSS(project string) {
	str := fmt.Sprint(`
	.heading {
		margin: auto;
	}	
	`)

	nf, err := os.Create(project + "/assets/main.css")
	if err != nil {
		log.Fatal("error creating file", err)
	}
	defer nf.Close()

	io.Copy(nf, strings.NewReader(str))
}
