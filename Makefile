# Define echo colors
# Use colors as follows:
# - red for error
# - green for success
# - yellow for info
define echo_red
	@echo "\033[31m$(1)\033[0m"
endef

define echo_green
	@echo "\033[32m$(1)\033[0m"
endef

define echo_yellow
	@echo "\033[33m$(1)\033[0m"
endef

# Check if environment is set-up properly
define check_env
	$(if $(GOBIN),,
	@# else
		$(call echo_red,"GOBIN is not set")
		@exit 1)
endef

define check_cmds
	$(foreach cmd, $(1),
		$(if $(shell command -v $(cmd) &> /dev/null || echo "not found"),
		@# then
			$(call echo_red,"$(cmd) not found")
			@exit 1)
	)
endef

# Install app
# 1 - app_name, 2 - app_path
define install_app
	$(call echo_yellow,"Installing $(1)...")
	@go install "cmd/$(2)"
endef

# Uninstall app

# 1 - app_name
define uninstall_app
	$(if $(shell test -f "$(GOBIN)/$(1)" &> /dev/null && echo "found"),
	@# then
		$(call echo_yellow,"Uninstalling $(GOBIN)/$(1)...")
		@rm "$(GOBIN)/$(1)",
	@# else
		$(call echo_yellow,"$(GOBIN)/$(1) not found. Skipping..."))
endef

.PHONY: all
all:

.PHONY: build
build: build-binary build-docker-image

.PHONY: build-binary
build-binary: check-compliance
	$(call install_app,"storygraph","storygraph/storygraph.go")

.PHONY: build-docker-image
build-docker-image: check-compliance

.PHONY: clean
clean: clean-binary clean-docker-image

.PHONY: clean-binary
clean-binary: check-compliance
	$(call uninstall_app,"storygraph")

.PHONY: clean-docker-image
clean-docker-image: check-compliance

.PHONY: check-compliance
check-compliance:
	$(call echo_yellow,"Checking environments variables...")
	$(call check_env)
	$(call echo_yellow,"Checking commands...")
	$(call check_cmds,go)