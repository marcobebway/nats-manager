PROJECT_ROOT ?= ../..
KYMA_CLI ?= "${PROJECT_ROOT}/${KYMA}"
CLUSTER_NAME ?= kyma
REGISTRY_PORT ?= 5001

include ${PROJECT_ROOT}/.env

ifndef MODULE_SHA
    MODULE_SHA = ${shell git rev-parse --short HEAD}
endif

.PHONY: module-build
module-build: ## Build the Module, push it to a registry and print it based on the commit hash.
	@make -C ${PROJECT_ROOT} module-build \
		IMG=${IMG} \
		MODULE_REGISTRY=${MODULE_REGISTRY} \
		MODULE_VERSION=${MODULE_VERSION}-${MODULE_SHA}
	cp -f ${PROJECT_ROOT}/template.yaml ${ARTIFACTS}/module-template.yaml
	@echo "Added template to job artifacts: ${ARTIFACTS}/module-template.yaml"
	@echo "\n~~~~~~~~~~~~BEGINNING OF MODULE TEMPLATE~~~~~~~~~~~~~~"
	@cat ${ARTIFACTS}/module-template.yaml
	@echo "\n~~~~~~~~~~~~~~~END OF MODULE TEMPLATE~~~~~~~~~~~~~~~~"

.PHONY: download-module-template
download-module-template: ## Downloads the module-template from the module-build prow job.
	export PROJECT_ROOT="${PROJECT_ROOT}" \
	&& ../get_module_template_from_build_job.sh

.PHONY: install-latest-module-template-fast
install-latest-module-template-fast: ## Downloads and applies the latest released module-template from fast channel.
	kubectl apply -f https://github.com/kyma-project/nats-manager/releases/latest/download/module-template.yaml

.PHONY: create-kyma-system-ns
create-kyma-system-ns: ## Create kyma-system namespace.
	kubectl create ns kyma-system

.PHONY: install-module-template
install-module-template: ## Apply NATS moduletemplate.
	kubectl apply -f ${PROJECT_ROOT}/module-template.yaml

.PHONY: enable-module-without-default-cr
enable-module-without-default-cr: ## Enable NATS module without creating default NATS CR.
	"${KYMA_CLI}" alpha enable module nats -c fast -n kyma-system -p Ignore

.PHONY: install-lifecycle-manager
install-lifecycle-manager: ## Deploys lifecycle-manager.
	"${KYMA_CLI}" alpha deploy \
		--ci \
		--force-conflicts

.PHONY: create-k3d
create-k3d: ## Create k3d with kyma CRDs.
	"${KYMA_CLI}" provision k3d -p 8081:80@loadbalancer -p 8443:443@loadbalancer --registry-port ${REGISTRY_PORT} --name ${CLUSTER_NAME} --ci

.PHONY: verify-kyma
verify-kyma: ## Wait for Kyma CR to be in Ready state.
	../verify_kyma_status.sh

.PHONY: install-k3d-tools
install-k3d-tools: ## Create k3d with kyma CRDs.
	curl -s https://raw.githubusercontent.com/k3d-io/k3d/main/install.sh | TAG=${K3D_VERSION} bash
