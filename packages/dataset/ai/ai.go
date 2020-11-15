package ai

import (
	"log"
	"os/exec"
)

func FastText(filename string) {

	// fasttext skipgram -input data/fil9 -output result/fil9

	app := "fasttext"

	arg0 := "skipgram"
	arg1 := "-input"
	arg2 := "/home/kuranov/Code/golang/service/sys/" + filename
	arg3 := "-output"
	arg4 := "/home/kuranov/Code/golang/service/sys/" + filename

	cmd := exec.Command(app, arg0, arg1, arg2, arg3, arg4)
	err := cmd.Run()
	if err != nil {
		log.Fatal(err)
	}

}
