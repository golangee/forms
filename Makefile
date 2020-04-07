## the binary name
ARTIFACT_NAME = wtkserve

## for the module itself
MODULE_PATH = github.com/worldiety/wtk

## the path which contains the main package to execute
MAIN_PATH = github.com/worldiety/wtk/cmd/wtkserve

## for ldflags replacement
BUILD_FILE_PATH = ${MODULE_PATH}

## which linter version to use?
GOLANGCI_LINT_VERSION = v1.24.0

LDFLAGS = -X $(MODULE_PATH).BuildGitCommit=$(CI_COMMIT_SHA) -X $(MODULE_PATH).BuildGitBranch=$(CI_COMMIT_REF_NAME)

TMP_DIR = $(TMPDIR)/$(MODULE_PATH)
BUILD_DIR = .build
TOOLSDIR = $(TMP_DIR)
GO = go
GOLANGCI_LINT = ${TOOLSDIR}/golangci-lint
GOLINT = ${TOOLSDIR}/golint
TMP_GOPATH = $(CURDIR)/${TOOLSDIR}/.gopath

all: generate lint test build ## Runs lint, test and build

clean: ## Removes any temporary and output files
	rm -rf ${buildDir}

lint: ## Executes all linters
	${GOLANGCI_LINT} run --enable-all --exclude-use-default=false

test: ## Executes the tests
	${GO} test -race ./...

.PHONY: build generate setup

build: generate ## Performs a build and puts everything into the build directory
	${GO} generate ./...
	${GO} build -ldflags "${LDFLAGS}" -o ${BUILD_DIR}/${ARTIFACT_NAME} ${MAIN_PATH}


run: build ## Starts the compiled program
	${BUILD_DIR}/${ARTIFACT_NAME}


updateMDC: ## updates Material Components for Web
	curl -L https://unpkg.com/material-components-web@latest/dist/material-components-web.min.css -o theme/material/material-components-web.min.css
	curl -L https://unpkg.com/material-components-web@latest/dist/material-components-web.min.css.map -o theme/material/material-components-web.min.css.map
	curl -L https://unpkg.com/material-components-web@latest/dist/material-components-web.min.js -o theme/material/material-components-web.min.js
	curl -L https://unpkg.com/material-components-web@latest/dist/material-components-web.js.map -o theme/material/material-components-web.js.map
	curl -L https://unpkg.com/material-components-web@latest/dist/material-components-web.js -o theme/material/material-components-web.js
	curl -L https://fonts.gstatic.com/s/materialicons/v50/flUhRq6tzZclQEJ-Vdg-IuiaDsNcIhQ8tQ.woff2 -o theme/material/materialicons.woff2

embedd: ## embedds theme/material
	$(shell find -E theme/material/ -regex '.*\.(wasm|js|html|css|woff2|ttf|map)'  -exec brotli -f -Z -w 24 "{}" \; )
	$(shell find -E theme/material/ -regex '.*\.(wasm|js|html|css|woff2|ttf|map)'  -exec gzip -k -f --best "{}" \; )
	${GO} run gen/gen.go

generate: ## Executes go generate
	${GO} generate

setup: installGolangCi ## Installs golangci-lint

installGolangCi:
	mkdir -p ${TOOLSDIR}
	curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(TOOLSDIR) $(GOLANGCI_LINT_VERSION)


help: ## Shows this help
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

.DEFAULT_GOAL := help

