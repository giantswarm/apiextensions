package crd

import (
	"fmt"
	"k8s.io/code-generator/cmd/client-gen/types"
	"k8s.io/gengo/args"
	"sigs.k8s.io/controller-tools/pkg/crd"
	"sigs.k8s.io/controller-tools/pkg/genall"
	"sigs.k8s.io/controller-tools/pkg/markers"
)


var (
	allGenerators = map[string]genall.Generator{
		"crd":         crd.Generator{},
	}

	allOutputRules = map[string]genall.OutputRule{
		"dir":       genall.OutputToDirectory(""),
	}

	optionsRegistry = &markers.Registry{}
)

func init() {
	for genName, gen := range allGenerators {
		// make the generator options marker itself
		defn := markers.Must(markers.MakeDefinition(genName, markers.DescribesPackage, gen))
		if err := optionsRegistry.Register(defn); err != nil {
			panic(err)
		}
		if helpGiver, hasHelp := gen.(genall.HasHelp); hasHelp {
			if help := helpGiver.Help(); help != nil {
				optionsRegistry.AddHelp(defn, help)
			}
		}

		// make per-generation output rule markers
		for ruleName, rule := range allOutputRules {
			ruleMarker := markers.Must(markers.MakeDefinition(fmt.Sprintf("output:%s:%s", genName, ruleName), markers.DescribesPackage, rule))
			if err := optionsRegistry.Register(ruleMarker); err != nil {
				panic(err)
			}
			if helpGiver, hasHelp := rule.(genall.HasHelp); hasHelp {
				if help := helpGiver.Help(); help != nil {
					optionsRegistry.AddHelp(ruleMarker, help)
				}
			}
		}
	}

	// make "default output" output rule markers
	for ruleName, rule := range allOutputRules {
		ruleMarker := markers.Must(markers.MakeDefinition("output:"+ruleName, markers.DescribesPackage, rule))
		if err := optionsRegistry.Register(ruleMarker); err != nil {
			panic(err)
		}
		if helpGiver, hasHelp := rule.(genall.HasHelp); hasHelp {
			if help := helpGiver.Help(); help != nil {
				optionsRegistry.AddHelp(ruleMarker, help)
			}
		}
	}

	// add in the common options markers
	if err := genall.RegisterOptionsMarkers(optionsRegistry); err != nil {
		panic(err)
	}
}

func Generate(genericArgs args.GeneratorArgs, groups []types.GroupVersions) error {
	/*
	b, err := genericArgs.NewBuilder()
	if err != nil {
		return err
	}

	c, err := generator.NewContext(b, generators.NameSystems(), generators.DefaultNameSystem())
	if err != nil {
		return err
	}

	p := c.Universe.Package("github.com/giantswarm/apiextensions/pkg/apis/application/v1alpha1")
	*/
	rt, err := genall.FromOptions(optionsRegistry, []string{
		"crd",
		"paths=./pkg/apis/...",
		"output:crd:dir=docs/crd",
	})
	if err != nil {
		return err
	}
	if len(rt.Generators) == 0 {
		return fmt.Errorf("no generators specified")
	}

	if hadErrs := rt.Run(); hadErrs {
		return fmt.Errorf("not all generators ran successfully")
	}

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
