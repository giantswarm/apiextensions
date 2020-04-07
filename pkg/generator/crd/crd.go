package crd

import (
	"fmt"
	"k8s.io/code-generator/cmd/client-gen/generators"
	"k8s.io/code-generator/cmd/client-gen/types"
	"k8s.io/gengo/args"
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

	p := c.Universe.Package("github.com/giantswarm/apiextensions/pkg/apis/applications/v1alpha1")
	fmt.Println(p)

	/*
	for _, group := range groups {
		for _, version := range group.Versions {
			filename := fmt.Sprintf("%s_%s.yaml", crd.Spec.Group, crd.Spec.Names.Singular)
			path := filepath.Join("docs/crd", filename)
			encoded, err := yaml.Marshal(crd)
			if err != nil {
				log.Fatal(err)
			}
			err = ioutil.WriteFile(path, encoded, 0644)
			if err != nil {
				log.Fatal(err)
			}
		}
	}
	 */
	return nil
}
