# Directories.
TOOLS_DIR := scripts/tools
TOOLS_BIN_DIR := $(abspath $(TOOLS_DIR)/bin)

# Binaries.
# Need to use abspath so we can invoke these from subdirectories
CLIENT_GEN := $(abspath $(TOOLS_BIN_DIR)/client-gen)
CONTROLLER_GEN := $(abspath $(TOOLS_BIN_DIR)/controller-gen)
DEEPCOPY_GEN := $(abspath $(TOOLS_BIN_DIR)/deepcopy-gen)
GOIMPORTS := $(abspath $(TOOLS_BIN_DIR)/goimports)
GOLANGCI_LINT := $(abspath $(TOOLS_BIN_DIR)/golangci-lint)
KUSTOMIZE := $(abspath $(TOOLS_BIN_DIR)/kustomize)
ESC := $(abspath $(TOOLS_BIN_DIR)/esc)

BUILD_COLOR = \033[0;34m
GEN_COLOR = \033[0;32m

INPUT_DIRS := $(shell find ./pkg/apis -maxdepth 2 -mindepth 2 | paste -s -d, -)
GROUPS := $(shell find ./pkg/apis -maxdepth 2 -mindepth 2  | sed 's|./pkg/apis/||' | paste -s -d, -)
DEEPCOPY_FILES := $(shell find ./pkg/apis -name "zz_generated.deepcopy.go")

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

$(GOIMPORTS): $(TOOLS_DIR)/goimports/go.mod
	@echo "$(BUILD_COLOR)Building goimports"
	@cd $(TOOLS_DIR)/goimports; go build -tags=tools -o $(TOOLS_BIN_DIR)/goimports golang.org/x/tools/cmd/goimports

$(GOLANGCI_LINT): $(TOOLS_DIR)/golangci-lint/go.mod
	@echo "$(BUILD_COLOR)Building golangci-lint"
	@cd $(TOOLS_DIR)/golangci-lint; go build -tags=tools -o $(TOOLS_BIN_DIR)/golangci-lint github.com/golangci/golangci-lint/cmd/golangci-lint

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
	@$(MAKE) imports
	@$(MAKE) patch

.PHONY: verify
verify:
	@$(MAKE) clean
	@$(MAKE) generate
	@$(MAKE) lint

.PHONY: generate-clientset
generate-clientset: $(CLIENT_GEN)
	@echo "$(GEN_COLOR)Generating clientset"
	@$(CLIENT_GEN) \
		--clientset-name versioned \
		--input $(GROUPS) \
		--input-base github.com/giantswarm/apiextensions/pkg/apis \
		--output-package github.com/giantswarm/apiextensions/pkg/clientset \
		--output-base ./scripts \
		--go-header-file ./scripts/boilerplate.go.txt
	@cp -R scripts/github.com/giantswarm/apiextensions/pkg/clientset/versioned pkg/clientset
	@rm -rf scripts/github.com/

.PHONY: generate-deepcopy
generate-deepcopy: $(DEEPCOPY_GEN)
	@echo "$(GEN_COLOR)Generating deepcopy"
	@$(DEEPCOPY_GEN) \
		--input-dirs $(INPUT_DIRS) \
		--output-base . \
		--output-file-base zz_generated.deepcopy \
		--go-header-file ./scripts/boilerplate.go.txt

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

.PHONY: lint
lint: $(GOLANGCI_LINT)
	@echo "$(GEN_COLOR)Running golangci-lint"
	@$(GOLANGCI_LINT) run -E gosec -E goconst

.PHONY: imports
imports: $(GOIMPORTS)
	@echo "$(GEN_COLOR)Sorting imports"
	@$(GOIMPORTS) -local github.com/giantswarm/apiextension -w ./pkg

.PHONY: patch
patch:
	@echo "$(GEN_COLOR)Applying patch"
	@git apply scripts/generated.patch

.PHONY: clean
clean:
	@echo "$(GEN_COLOR)Cleaning generated files"
	@rm -rf config/crd/v1 config/crd/v1beta1 pkg/clientset/versioned $(DEEPCOPY_FILES)
