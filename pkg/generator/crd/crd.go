package crd

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1beta1"
	toolscrd "sigs.k8s.io/controller-tools/pkg/crd"
	"sigs.k8s.io/controller-tools/pkg/genall"
	"sigs.k8s.io/controller-tools/pkg/markers"
	"sigs.k8s.io/yaml"
)

const crdDirectory = "config/crd/bases/"

var (
	allGenerators = map[string]genall.Generator{
		"crd": toolscrd.Generator{},
	}

	allOutputRules = map[string]genall.OutputRule{
		"dir": genall.OutputToDirectory(""),
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

const %sYAML = %s

func New%sCRD() *v1beta1.CustomResourceDefinition {
	var crd v1beta1.CustomResourceDefinition
	_ = yaml.Unmarshal([]byte(%sYAML), &crd)
	return &crd
}
`

func Generate() error {
	rt, err := genall.FromOptions(optionsRegistry, []string{
		"crd",
		"paths=./pkg/apis/...",
		fmt.Sprintf("output:crd:dir=%s", crdDirectory),
	})
	if err != nil {
		return err
	}
	if hadErrs := rt.Run(); hadErrs {
		return fmt.Errorf("not all generators ran successfully")
	}

	directory, err := ioutil.ReadDir(crdDirectory)
	for _, crdFile := range directory {
		contents, err := ioutil.ReadFile(crdDirectory + crdFile.Name())
		if err != nil {
			return err
		}
		var crd v1beta1.CustomResourceDefinition
		_ = yaml.Unmarshal(contents, &crd)
		group := strings.Split(crd.Spec.Group, ".")[0]
		rendered := fmt.Sprintf(yamlTemplate, group, crd.Spec.Names.Plural, "`"+string(contents)+"`", crd.Spec.Names.Kind, crd.Spec.Names.Plural)
		filename := "pkg/crds/" + group + "/" + crd.Spec.Names.Plural + ".go"
		_ = os.Mkdir("pkg/crds/"+group, 0755)
		err = ioutil.WriteFile(filename, []byte(rendered), 0644)
		if err != nil {
			return err
		}
		fmt.Println("generated", filename)
	}
	return nil
}