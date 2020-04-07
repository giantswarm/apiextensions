package deepcopy

import (
	"io"
	"os"
	"path/filepath"

	"k8s.io/code-generator/cmd/client-gen/types"
	"k8s.io/gengo/args"
	"k8s.io/gengo/examples/deepcopy-gen/generators"
	"k8s.io/gengo/generator"
)

func Generate(genericArgs args.GeneratorArgs, groups []types.GroupVersions) error {
	b, err := genericArgs.NewBuilder()
	if err != nil {
		return err
	}

	c, err := generator.NewContext(b, generators.NameSystems(), generators.DefaultNameSystem())
	if err != nil {
		return err
	}

	packages := generators.Packages(c, &genericArgs)
	if err := c.ExecutePackages(genericArgs.OutputBase, packages); err != nil {
		return err
	}

	for _, group := range groups {
		for _, version := range group.Versions {
			source, err := os.Open(filepath.Join(genericArgs.OutputBase, version.Package, genericArgs.OutputFileBaseName+".go"))
			if err != nil {
				return err
			}
			destination, err := os.Create(filepath.Join("pkg/apis", string(group.Group), string(version.Version), genericArgs.OutputFileBaseName+".go"))
			if err != nil {
				source.Close()
				return err
			}
			_, err = io.Copy(destination, source)
			source.Close()
			destination.Close()
			if err != nil {
				return err
			}
		}
	}

	if err := os.RemoveAll(genericArgs.OutputBase); err != nil {
		return err
	}
	return nil
}
