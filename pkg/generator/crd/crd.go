package crd

import (
	"fmt"
	"io/ioutil"
	"k8s.io/code-generator/cmd/client-gen/types"
	"k8s.io/gengo/args"
	"sigs.k8s.io/controller-tools/pkg/crd"
	"sigs.k8s.io/controller-tools/pkg/genall"
	"sigs.k8s.io/controller-tools/pkg/markers"
	"strings"
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

const yamlTemplate = `package %s

import (
	"k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1beta1"
	"sigs.k8s.io/yaml"
)

const crdYAML = %s

func New%sCRD() *v1beta1.CustomResourceDefinition {
	var crd v1beta1.CustomResourceDefinition
	_ = yaml.Unmarshal([]byte(crdYAML), &crd)
	return &crd
}
`

func Generate(genericArgs args.GeneratorArgs, groups []types.GroupVersions) error {
	rt, err := genall.FromOptions(optionsRegistry, []string{
		"crd",
		"paths=./pkg/apis/...",
		"output:crd:dir=docs/crd",
	})
	if err != nil {
		return err
	}
	if hadErrs := rt.Run(); hadErrs {
		return fmt.Errorf("not all generators ran successfully")
	}
	d, err := ioutil.ReadDir("./docs/crd")
	for _, dir := range d {
		contents, err := ioutil.ReadFile("docs/crd/" + dir.Name())
		if err != nil {
			return err
		}
		split := strings.Split(dir.Name(), ".")
		group := split[0]
		splitSuffix := strings.Split(split[len(split)-2], "_")
		kind := splitSuffix[1]
		rendered := fmt.Sprintf(yamlTemplate, group, "`"+string(contents)+"`", kind)
		err = ioutil.WriteFile("pkg/crds/"+group+"/"+kind+".go", []byte(rendered), 0644)
		if err != nil {
			return err
		}
	}
	return nil
}
