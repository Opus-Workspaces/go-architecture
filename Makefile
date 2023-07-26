MAIN_FILE=app/cmd/main/main.go
PACKAGE=itss.edu.vn/dist
PATH_ENV_BUILD='${PACKAGE}/app/config.buildEnv'

# Run with environment variable
define run
	@echo "Running $(1)..."
	go run $(MAIN_FILE) --env=$(1)
endef

dev:
	$(call run,dev)
staging:
	$(call run,staging)
prod:
	$(call run,prod)

# Build with environment variable
define build
	@echo "Building $(1)..."
	rm -rf build && go build -o build/dist -ldflags='-X "${PATH_ENV_BUILD}=$(1)"' "${MAIN_FILE}" && ./build/dist
endef

.PHONY: build-dev
build-dev:
	$(call build,dev)

.PHONY: build-staging
build-staging:
	$(call build,staging)

.PHONY: build-prod
build-prod:
	$(call build,prod)