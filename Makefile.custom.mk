# Directories.
APIS_DIR := pkg/apis
CLIENTSET_DIR := pkg/clientset
CRDV1_DIR := config/crd/v1
CRDV1BETA1_DIR := config/crd/v1beta1
SCRIPTS_DIR := hack
TOOLS_DIR := $(SCRIPTS_DIR)/tools
TOOLS_BIN_DIR := $(abspath $(TOOLS_DIR)/bin)

# Binaries.
# Need to use abspath so we can invoke these from subdirectories
CLIENT_GEN := $(abspath $(TOOLS_BIN_DIR)/client-gen)
CONTROLLER_GEN := $(abspath $(TOOLS_BIN_DIR)/controller-gen)
GOIMPORTS := $(abspath $(TOOLS_BIN_DIR)/goimports)
KUSTOMIZE := $(abspath $(TOOLS_BIN_DIR)/kustomize)
ESC := $(abspath $(TOOLS_BIN_DIR)/esc)

BUILD_COLOR = ""
GEN_COLOR = ""
NO_COLOR = ""

ifneq (, $(shell command -v tput))
ifeq ($(shell test `tput colors` -ge 8 && echo "yes"), yes)
BUILD_COLOR = \033[0;34m
GEN_COLOR = \033[0;32m
NO_COLOR = \033[0m
endif
endif

DEEPCOPY_BASE = zz_generated.deepcopy
MODULE = $(shell go list -m)
BOILERPLATE = $(SCRIPTS_DIR)/boilerplate.go.txt
PATCH_FILE = $(SCRIPTS_DIR)/generated.patch
YEAR = $(shell date +'%Y')

INPUT_DIRS := $(shell find ./$(APIS_DIR) -maxdepth 2 -mindepth 2 | paste -s -d, -)
GROUPS := $(shell find $(APIS_DIR) -maxdepth 2 -mindepth 2  | sed 's|pkg/apis/||' | paste -s -d, -)
DEEPCOPY_FILES := $(shell find $(APIS_DIR) -name $(DEEPCOPY_BASE).go)

all: generate

$(CLIENT_GEN): $(TOOLS_DIR)/client-gen/go.mod
	@echo "$(BUILD_COLOR)Building client-gen$(NO_COLOR)"
	cd $(TOOLS_DIR)/client-gen \
	&& go build -tags=tools -o $(CLIENT_GEN) k8s.io/code-generator/cmd/client-gen

$(CONTROLLER_GEN): $(TOOLS_DIR)/controller-gen/go.mod
	@echo "$(BUILD_COLOR)Building controller-gen$(NO_COLOR)"
	cd $(TOOLS_DIR)/controller-gen \
 	&& go build -tags=tools -o $(CONTROLLER_GEN) sigs.k8s.io/controller-tools/cmd/controller-gen

$(GOIMPORTS): $(TOOLS_DIR)/goimports/go.mod
	@echo "$(BUILD_COLOR)Building goimports$(NO_COLOR)"
	cd $(TOOLS_DIR)/goimports \
	&& go build -tags=tools -o $(GOIMPORTS) golang.org/x/tools/cmd/goimports

$(KUSTOMIZE): $(TOOLS_DIR)/kustomize/go.mod
	@echo "$(BUILD_COLOR)Building kustomize$(NO_COLOR)"
	cd $(TOOLS_DIR)/kustomize \
	&& go build -tags=tools -o $(KUSTOMIZE) sigs.k8s.io/kustomize/kustomize/v3

$(ESC): $(TOOLS_DIR)/esc/go.mod
	@echo "$(BUILD_COLOR)Building esc$(NO_COLOR)"
	@cd $(TOOLS_DIR)/esc \
	&& go build -tags=tools -o $(ESC) github.com/mjibson/esc

.PHONY: generate
generate:
	@$(MAKE) generate-clientset
	@$(MAKE) generate-deepcopy
	@$(MAKE) generate-manifests
	@$(MAKE) generate-fs
	@$(MAKE) ci-imports
	@$(MAKE) patch

.PHONY: verify
verify:
	@$(MAKE) clean-generated
	@$(MAKE) generate
	git diff --exit-code

.PHONY: generate-clientset
generate-clientset: $(CLIENT_GEN)
	@echo "$(GEN_COLOR)Generating clientset$(NO_COLOR)"
	$(CLIENT_GEN) \
	--clientset-name versioned \
	--input $(GROUPS) \
	--input-base $(MODULE)/$(APIS_DIR) \
	--output-package $(MODULE)/$(CLIENTSET_DIR) \
	--output-base $(SCRIPTS_DIR) \
	--go-header-file $(BOILERPLATE)
	cp -R $(SCRIPTS_DIR)/$(MODULE)/$(CLIENTSET_DIR)/versioned $(CLIENTSET_DIR)
	rm -rf $(SCRIPTS_DIR)/github.com/

.PHONY: generate-deepcopy
generate-deepcopy: $(CONTROLLER_GEN)
	@echo "$(GEN_COLOR)Generating deepcopy$(NO_COLOR)"
	$(CONTROLLER_GEN) \
	object:headerFile=$(BOILERPLATE),year=$(YEAR) \
	paths=./$(APIS_DIR)/...

.PHONY: generate-manifests
generate-manifests: $(CONTROLLER_GEN) $(KUSTOMIZE)
	@echo "$(GEN_COLOR)Generating CRDs$(NO_COLOR)"
	cd $(SCRIPTS_DIR); ./generate-manifests.sh

.PHONY: generate-fs
generate-fs: $(ESC) config/crd
	@echo "$(GEN_COLOR)Generating filesystem$(NO_COLOR)"
	$(ESC) \
	-o pkg/crd/internal/zz_generated.fs.go \
	-pkg internal \
	-modtime 0 \
	config/crd

.PHONY: ci-imports
ci-imports: $(GOIMPORTS)
	@echo "$(GEN_COLOR)Sorting imports$(NO_COLOR)"
	$(GOIMPORTS) -local $(MODULE) -w ./pkg

.PHONY: patch
patch:
	@echo "$(GEN_COLOR)Applying patch$(NO_COLOR)"
	git apply $(PATCH_FILE)

.PHONY: clean-generated
clean-generated:
	@echo "$(GEN_COLOR)Cleaning generated files$(NO_COLOR)"
	rm -rf $(CRDV1_DIR) $(CRDV1BETA1_DIR) $(CLIENTSET_DIR)/versioned $(DEEPCOPY_FILES)

.PHONY: clean-tools
clean-tools:
	@echo "$(GEN_COLOR)Cleaning tools$(NO_COLOR)"
	rm -rf $(TOOLS_BIN_DIR)
