package crd

import (
	"fmt"
	"strings"
	"testing"

	"sigs.k8s.io/controller-tools/pkg/crd"
	"sigs.k8s.io/controller-tools/pkg/genall"
	"sigs.k8s.io/controller-tools/pkg/markers"
)

func Test_Comments(t *testing.T) {
	crdGenerator := crd.Generator{}

	// allOutputRules defines the list of all known output rules, giving
	// them names for use on the command line.
	// Each output rule turns into two command line options:
	// - output:<generator>:<form> (per-generator output)
	// - output:<form> (default output)
	allOutputRules := map[string]genall.OutputRule{
		"dir":       genall.OutputToDirectory(""),
		"none":      genall.OutputToNothing,
		"stdout":    genall.OutputToStdout,
		"artifacts": genall.OutputArtifacts{},
	}

	// optionsRegistry contains all the marker definitions used to process command line options
	optionsRegistry := &markers.Registry{}

	// make the generator options marker itself
	defn := markers.Must(markers.MakeDefinition("crd", markers.DescribesPackage, crdGenerator))
	if err := optionsRegistry.Register(defn); err != nil {
		panic(err)
	}

	// make per-generation output rule markers
	for ruleName, rule := range allOutputRules {
		ruleMarker := markers.Must(markers.MakeDefinition(fmt.Sprintf("output:%s:%s", "crd", ruleName), markers.DescribesPackage, rule))
		if err := optionsRegistry.Register(ruleMarker); err != nil {
			panic(err)
		}
		if helpGiver, hasHelp := rule.(genall.HasHelp); hasHelp {
			if help := helpGiver.Help(); help != nil {
				optionsRegistry.AddHelp(ruleMarker, help)
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

	// otherwise, set up the runtime for actually running the generators
	rt, err := genall.FromOptions(optionsRegistry, []string{
		"crd",
		"paths=../../pkg/apis/...",
		"output:stdout",
	})
	if err != nil {
		t.Fatal(err)
	}

	ctx := rt.GenerationContext
	var gen genall.Generator = &crdGenerator
	ctx.OutputRule = rt.OutputRules.ForGenerator(&gen)
	parser := &crd.Parser{
		Collector: ctx.Collector,
		Checker:   ctx.Checker,
	}

	crd.AddKnownTypes(parser)
	for _, root := range ctx.Roots {
		parser.NeedPackage(root)
	}

	metav1Pkg := crd.FindMetav1(ctx.Roots)
	if metav1Pkg == nil {
		t.Fatal("here")
	}

	kubeKinds := crd.FindKubeKinds(parser, metav1Pkg)
	if len(kubeKinds) == 0 {
		t.Fatal("here")
	}

	for _, groupKind := range kubeKinds {
		for pkg, gv := range parser.GroupVersions {
			if gv.Group != groupKind.Group {
				continue
			}
			if err := markers.EachType(parser.Collector, pkg, func(info *markers.TypeInfo) {
				for _, field := range info.Fields {
					_, optional := field.Markers["kubebuilder:validation:Optional"]
					omitempty := strings.Contains(field.Tag.Get("json"), ",omitempty")
					if !optional && omitempty {
						t.Errorf("in CR %s, field %s for type %s has omitempty, should have corresponding kubebuilder:validation:Optional comment", groupKind.Kind, field.Name, info.Name)
					}
				}
			}); err != nil {
				t.Error(err)
			}
		}
	}
}

func Test_List(t *testing.T) {
	crdV1, err := ListV1()
	if err != nil {
		t.Fatalf("expected err to be nil: %s", err)
	}

	if crdV1 == nil {
		t.Fatal("expected crd slice to not be nil")
	}

	if len(crdV1) == 0 {
		t.Fatal("expected crd slice to contain at least one item")
	}
}

func Test_Load(t *testing.T) {
	nilPanicMatcher := func(ret interface{}) bool {
		return ret == nil
	}
	notFoundPanicMatcher := func(ret interface{}) bool {
		err, ok := ret.(error)
		return ok && IsNotFound(err)
	}
	testCases := []struct {
		name            string
		inputGroup      string
		inputKind       string
		inputCRDVersion string
		panicMatcher    func(ret interface{}) bool
	}{
		{
			name:            "case 0: v1beta1 CRD loads normally",
			inputGroup:      "application.giantswarm.io",
			inputKind:       "App",
			inputCRDVersion: "v1beta1",
			panicMatcher:    nilPanicMatcher,
		},
		{
			name:            "case 1: non-existent v1beta1 CRD panics with notFoundError",
			inputGroup:      "application.giantswarm.io",
			inputKind:       "Bad",
			inputCRDVersion: "v1beta1",
			panicMatcher:    notFoundPanicMatcher,
		},
		{
			name:            "case 2: v1 CRD loads normally",
			inputGroup:      "infrastructure.giantswarm.io",
			inputKind:       "AWSCluster",
			inputCRDVersion: "v1",
			panicMatcher:    nilPanicMatcher,
		},
		{
			name:            "case 3: non-existent v1 CRD panics with notFoundError",
			inputGroup:      "application.giantswarm.io",
			inputKind:       "Bad",
			inputCRDVersion: "v1",
			panicMatcher:    notFoundPanicMatcher,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			defer func() {
				err := recover()
				if !tc.panicMatcher(err) {
					t.Errorf("unexpected panic: %#v", err)
				}
			}()
			var crd interface{}
			switch tc.inputCRDVersion {
			case "v1beta1":
				crd = LoadV1Beta1(tc.inputGroup, tc.inputKind)
			case "v1":
				crd = LoadV1(tc.inputGroup, tc.inputKind)
			}
			if crd == nil {
				t.Errorf("nil crd")
			}
		})
	}
}
