# Directories.
TOOLS_DIR := scripts/tools
TOOLS_BIN_DIR := $(abspath $(TOOLS_DIR)/bin)

# Binaries.
# Need to use abspath so we can invoke these from subdirectories
CLIENT_GEN := $(abspath $(TOOLS_BIN_DIR)/client-gen)
CONTROLLER_GEN := $(abspath $(TOOLS_BIN_DIR)/controller-gen)
DEEPCOPY_GEN := $(abspath $(TOOLS_BIN_DIR)/deepcopy-gen)
KUSTOMIZE := $(abspath $(TOOLS_BIN_DIR)/kustomize)
ESC := $(abspath $(TOOLS_BIN_DIR)/esc)

all: generate

$(CLIENT_GEN): $(TOOLS_DIR)/code-generator/go.mod
	cd $(TOOLS_DIR)/code-generator; go build -tags=tools -o $(TOOLS_BIN_DIR)/client-gen k8s.io/code-generator/cmd/client-gen

$(CONTROLLER_GEN): $(TOOLS_DIR)/controller-gen/go.mod
	cd $(TOOLS_DIR)/controller-gen; go build -tags=tools -o $(TOOLS_BIN_DIR)/controller-gen sigs.k8s.io/controller-tools/cmd/controller-gen

$(DEEPCOPY_GEN): $(TOOLS_DIR)/code-generator/go.mod
	cd $(TOOLS_DIR)/code-generator; go build -tags=tools -o $(TOOLS_BIN_DIR)/deepcopy-gen k8s.io/code-generator/cmd/deepcopy-gen

$(KUSTOMIZE): $(TOOLS_DIR)/kustomize/go.mod
	cd $(TOOLS_DIR)/kustomize; go build -tags=tools -o $(TOOLS_BIN_DIR)/kustomize sigs.k8s.io/kustomize/kustomize/v3

$(ESC): $(TOOLS_DIR)/esc/go.mod
	cd $(TOOLS_DIR)/esc; go build -tags=tools -o $(TOOLS_BIN_DIR)/esc github.com/mjibson/esc

.PHONY: generate
generate:
	$(MAKE) generate-clientset
	$(MAKE) generate-deepcopy
	$(MAKE) generate-manifests
	$(MAKE) generate-static

.PHONY: generate-clientset
generate-clientset: $(CLIENT_GEN)
	$(CLIENT_GEN) \
		object:headerFile=./scripts/boilerplate.go.txt \
		paths=./pkg/apis/...

.PHONY: generate-deepcopy
generate-deepcopy: $(DEEPCOPY_GEN)
	$(DEEPCOPY_GEN) \
		object:headerFile=./scripts/boilerplate.go.txt \
		paths=./pkg/apis/...

.PHONY: generate-manifests
generate-manifests: $(CONTROLLER_GEN) $(KUSTOMIZE)
	$(CONTROLLER_GEN) \
		paths=./pkg/apis/... \
		crd:crdVersions=v1 \
		output:crd:dir=./config/crd/v1
	$(CONTROLLER_GEN) \
		paths=./pkg/apis/... \
		crd:crdVersions=v1beta1 \
		output:crd:dir=./config/crd/v1beta1

.PHONY: generate-static
generate-static: $(ESC) config/crd
	$(ESC) \
		-o pkg/crd/static.go \
		-pkg crd \
		-modtime 0 \
		-private \
		config/crd
