package main

import (
	"fmt"
	"os"

	"k8s.io/gengo/examples/deepcopy-gen/generators"

	generatorargs "k8s.io/code-generator/cmd/deepcopy-gen/args"
)

func main() {
	genericArgs, _ := generatorargs.NewDefaults()
	genericArgs.GoHeaderFilePath = "./scripts/boilerplate.go.txt"
	genericArgs.OutputPackagePath = "./pkg"
	genericArgs.InputDirs = []string{"github.com/giantswarm/apiextensions/pkg/apis/core/v1alpha1"}
	err := genericArgs.Execute(
		generators.NameSystems(),
		generators.DefaultNameSystem(),
		generators.Packages,
	)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	os.Exit(0)
}
