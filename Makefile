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

BUILD_COLOR = \033[0;34m
GEN_COLOR = \033[0;32m

all: generate

$(CLIENT_GEN): $(TOOLS_DIR)/client-gen/go.mod
	@echo "$(BUILD_COLOR)Building client-gen"
	@cd $(TOOLS_DIR)/client-gen; go build -tags=tools -o $(TOOLS_BIN_DIR)/client-gen k8s.io/code-generator/cmd/client-gen

$(CONTROLLER_GEN): $(TOOLS_DIR)/controller-gen/go.mod
	@echo "$(BUILD_COLOR)Building controller-gen"
	@cd $(TOOLS_DIR)/controller-gen; go build -tags=tools -o $(TOOLS_BIN_DIR)/controller-gen sigs.k8s.io/controller-tools/cmd/controller-gen

$(DEEPCOPY_GEN): $(TOOLS_DIR)/deepcopy-gen/go.mod
	@echo "$(BUILD_COLOR)Building deepcopy-gen"
	@cd $(TOOLS_DIR)/deepcopy-gen; go build -tags=tools -o $(TOOLS_BIN_DIR)/deepcopy-gen k8s.io/code-generator/cmd/deepcopy-gen

$(KUSTOMIZE): $(TOOLS_DIR)/kustomize/go.mod
	@echo "$(BUILD_COLOR)Building kustomize"
	@cd $(TOOLS_DIR)/kustomize; go build -tags=tools -o $(TOOLS_BIN_DIR)/kustomize sigs.k8s.io/kustomize/kustomize/v3

$(ESC): $(TOOLS_DIR)/esc/go.mod
	@echo "$(BUILD_COLOR)Building esc"
	@cd $(TOOLS_DIR)/esc; go build -tags=tools -o $(TOOLS_BIN_DIR)/esc github.com/mjibson/esc

.PHONY: generate
generate:
	@$(MAKE) generate-clientset
	@$(MAKE) generate-deepcopy
	@$(MAKE) generate-manifests
	@$(MAKE) generate-static

.PHONY: generate-clientset
generate-clientset: $(CLIENT_GEN)
	@echo "$(GEN_COLOR)Generating clientset"
	@$(CLIENT_GEN) \
		object:headerFile=./scripts/boilerplate.go.txt \
		paths=./pkg/apis/...

.PHONY: generate-deepcopy
generate-deepcopy: $(DEEPCOPY_GEN)
	@echo "$(GEN_COLOR)Generating deepcopy"
	@$(DEEPCOPY_GEN) \
		object:headerFile=./scripts/boilerplate.go.txt \
		paths=./pkg/apis/...

.PHONY: generate-manifests
generate-manifests: $(CONTROLLER_GEN) $(KUSTOMIZE)
	@echo "$(GEN_COLOR)Generating CRDs"
	@cd scripts; ./generate-manifests.sh

.PHONY: generate-static
generate-static: $(ESC) config/crd
	@echo "$(GEN_COLOR)Generating filesystem"
	@$(ESC) \
		-o pkg/crd/static.go \
		-pkg crd \
		-modtime 0 \
		-private \
		config/crd
