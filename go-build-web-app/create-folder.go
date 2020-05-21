package createwebapp

import "os"

func CreateFolder(project string) {
	err := os.Mkdir(project, 0755)
	if err != nil {
		panic(err)
	}

	err = os.Mkdir(project+"/templates", 0755)
	if err != nil {
		panic(err)
	}

	err = os.Mkdir(project+"/assets", 0755)
	if err != nil {
		panic(err)
	}
}
