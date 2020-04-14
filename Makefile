################################################################################
# Variables
#
# These are controlled by the CI/CD system when an official build is produced.
# Development builds are distinguished by git commit.
# These should be set via "make -e VERSION=1.0.0" REVISION="beta1"
################################################################################

# This is the application name, this affects version strings and packaging.
APPLICATION = couchbase-service-broker

# This is the release version, this affects version strings, packaging and
# also documentation.
# All development build should use 0.0.0 and rely on the git SHA embedded in the
# version string to uniquely identify a build.
# All documentation starts life being version agnostic e.g. 0.0.0.  When a
# release is created then the source is pre-processed and instances of 0.0.0
# are replaced with whatever VERSION has been set to.
VERSION = 0.0.0

# Revision is a packaging thing, but does allow the addition of beta suffixes.
# This is usually a patchset on top of the official release.
REVISION = 1

# This is the go import path of the module and affects client code.
IMPORT_PATH = github.com/couchbase/service-broker

# This is the docker image name to generate.  The tag is derived from the
# VERSION variable.
DOCKER_IMAGE = couchbase/service-broker

# This is the install prefix.
PREFIX = /usr

# This is the root of the install.
DESTDIR =

################################################################################
# Constants
#
# These are specific to the build system and should not be changed.
################################################################################

# The go binary.
GO_BINARY = $(shell which go)

# The path to the go binary.
GO_BINARY_PATH = $(dir $(GO_BINARY))

# Binary and packaging builds are created in this directory.
BUILD_DIR = build

# Files generated by the Kubernetes code generator are created in this
# directory.
GENERATED_DIR = generated

# Custom resource definitions are generated in this directory.
CRD_DIR = crds

# CRD resource file names that will be generated.
CRD_FILES = servicebroker.couchbase.com_servicebrokerconfigs.yaml

# CRD target files that will be generated.
CRDS = $(addprefix $(CRD_DIR)/,$(CRD_FILES))

# This is the git commit used to build the binaries.
COMMIT = $(shell git rev-parse HEAD)

# These are all source files, this triggers rebuilds of all binaries.
# This could be more targeted (e.g. omit the test code).
SOURCE = $(shell find . -name *.go -type f)

# The API source files define Kubernetes data structures, any modifications
# to these should trigger rebuilds of generated CRDs and client code.
APISRC = $(shell find pkg/apis -name [^z]*.go -type f)

# Module dependencies managed by go dep trigger a rebuild of all binaries.
DEPSRC = go.mod

# This is the main broker binary output file.
BROKER_BIN = $(BUILD_DIR)/bin/broker

# A list of all binary targets to be copied into an archive.
BINARIES = $(BROKER_BIN)

# This is the code coverage file used by unit testing.
COVER_FILE = /tmp/cover.out

# This is the base name for build archives e.g. package-0.0.0.
# This is especially important for RPM builds where the tarball and
# directory must be named this.
PACKAGE_BASENAME = $(APPLICATION)-$(VERSION)

# As above but with under scores because debian is different.
PACKAGE_BASENAME_DEB = $(APPLICATION)_$(VERSION)

# This is the name for the UNIX archive.
ARCHIVE_TGZ = $(PACKAGE_BASENAME)-$(REVISION).tar.gz

# This is the name for the Windows archive.
ARCHIVE_ZIP = $(PACKAGE_BASENAME)-$(REVISION).zip

# This is the name for the RPM archive.
ARCHIVE_RPM = $(PACKAGE_BASENAME)-$(REVISION).x86_64.rpm

# This is the name of the DEB archive.
ARCHIVE_DEB = $(PACKAGE_BASENAME_DEB)-$(REVISION)_amd64.deb

# This is the directory into which resources are installed before they are
# archived.
INSTALL_ROOT = $(DESTDIR)$(PREFIX)

# This is where binaries live.  It needs to be under the directory
# where the docker file resides or it won't work.
INSTALL_BIN = $(INSTALL_ROOT)/share/$(APPLICATION)/bin

# This is where shared files live e.g. documentation and examples.
INSTALL_SHARE = $(INSTALL_ROOT)/share/$(APPLICATION)

# Binary targets in the install, these are translated from build/bin to bin
# in the final install.
INSTALL_BIN_TARGETS = $(patsubst $(BUILD_DIR)/bin/%,$(INSTALL_BIN)/%,$(BINARIES))

# All shared source files to install (excluding binaries)
INSTALL_SHARE_SOURCES = LICENSE README.md Dockerfile

# Archive target files that live under share (excluding binaries)
INSTALL_SHARE_TARGETS = $(addprefix $(INSTALL_SHARE)/,$(INSTALL_SHARE_SOURCES))

# All install targets, changing any of these will result in the rebuild
# of an archive.
INSTALL_TARGETS = $(INSTALL_SHARE_TARGETS) $(INSTALL_BIN_TARGETS)

# This is the base directory to generate kubernetes API primitives from e.g.
# clients and CRDs.
GENAPIBASE = github.com/couchbase/service-broker/pkg/apis

# This is the list of APIs to generate clients for.
GENAPIS = $(GENAPIBASE)/servicebroker/v1alpha1

# These are generic arguments that need to be passed to client generation.
GENARGS = --go-header-file hack/boilerplate.go.txt --output-base ../../..

# This controls the name of the client that will be generated and it will affect
# code import paths.  This overrides the default "versioned".
GENCLIENTNAME = servicebroker

# This defines where clients will be generated.
GENCLIENTS = $(IMPORT_PATH)/$(GENERATED_DIR)/clientset

# This defines where listers will be generated.
GENLISTERS = $(IMPORT_PATH)/$(GENERATED_DIR)/listers

# This defines where informers will be generated.
GENINFORMERS = $(IMPORT_PATH)/$(GENERATED_DIR)/informers

# This is an ephemeral container image used for acceptance testing.
ACCEPTANCE_IMAGE = couchbase/service-broker-acceptance:latest

################################################################################
# Top level make targets.
#
# These are the commands intended to be run by an end user or CI system.
################################################################################

# These phony targets do not refer to actual files and are intended to be
# invoked by the end user.
.PHONY: all build crd container test unit lint cover acceptance install archive archive-tgz archive-zip rpm deb docs

# Main build target, makes the binary and CRD.
all: build crd

# Build the main binary.
build: $(BROKER_BIN)

# Build the CRDs.
crd: $(CRDS)

# Build a container image.
container: build
	docker build -f Dockerfile -t $(DOCKER_IMAGE):$(VERSION) .
	docker tag $(DOCKER_IMAGE):$(VERSION) $(DOCKER_IMAGE):latest

# Main test target, run linter and all tests.
test: lint unit

# Render code coverage (after running the test target) and display it in a browser.
cover:
	go tool cover -html=$(COVER_FILE)

# The linter must pass for all code submissions, it is controlled by .golangci.yml.
lint: $(GENERATED_DIR)
	go run github.com/golangci/golangci-lint/cmd/golangci-lint run

# The unit tests must pass for all code submissions, additionally code
# coverage should be checked to ensure code submissions actually work.
unit: $(GENERATED_DIR)
	go test -v -race -cover -coverpkg github.com/couchbase/service-broker/pkg/... -coverprofile=$(COVER_FILE) ./test/unit

# Acceptance testing builds a container and runs tests within Kubernetes.
# This is higher performance than having to connect over the internet all
# the time and avoids any pitfalls associated with external networking.
# Assumes that "docker build" is visible to the kubernetes context that
# is current (e.g. minikube for most of us).
acceptance: container crd
	GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go test -c -o build/bin/acceptance ./test/acceptance
	docker build -f test/acceptance/Dockerfile -t $(ACCEPTANCE_IMAGE) .
	kubectl apply -f test/acceptance/serviceaccount.yaml
	kubectl run -t -i acceptance --serviceaccount=couchbase-service-broker-acceptance --image=$(ACCEPTANCE_IMAGE) --image-pull-policy=Never --restart=Never -- -test.v -test.failfast -logtostderr -v 1; kubectl delete pod/acceptance

# Main install target, creates all archive files.
install: $(INSTALL_TARGETS)

# Main archival target, creates TGZ and ZIP archives.
archive: install archive-tgz archive-zip

# Create a TGZ release artefact.
archive-tgz: $(ARCHIVE_TGZ)

# Create a ZIP release artifacts.
archive-zip: $(ARCHIVE_ZIP)

# Build an RPM package.
rpm: $(ARCHIVE_RPM)

# Build a DEB package.
deb: $(ARCHIVE_DEB)

# Lint asciidoc files.
docs:
	hack/asciidoc-lint

# Clean all generated code and artifacts.
clean:
	rm -rf $(BUILD_DIR) *.tar.gz *.zip *.rpm *.deb

################################################################################
# Make rules
#
# These are the conditional bits that generate files based on other files.
################################################################################

# Generated code depends upon API sources. The code generator still requires a
# GOPATH style install hence the hacks with the output base.  This may get fixed
# in a later release.
$(GENERATED_DIR): $(APISRC)
	go run k8s.io/code-generator/cmd/deepcopy-gen --input-dirs $(GENAPIS) -O zz_generated.deepcopy --bounding-dirs $(GENAPIBASE) $(GENARGS)
	go run k8s.io/code-generator/cmd/client-gen --clientset-name $(GENCLIENTNAME) --input-base "" --input $(GENAPIS) --output-package $(GENCLIENTS) $(GENARGS)
	go run k8s.io/code-generator/cmd/lister-gen --input-dirs $(GENAPIS) --output-package $(GENLISTERS) $(GENARGS)
	go run k8s.io/code-generator/cmd/informer-gen --input-dirs $(GENAPIS) --versioned-clientset-package $(GENCLIENTS)/$(GENCLIENTNAME) --listers-package $(GENLISTERS) --output-package $(GENINFORMERS) $(GENARGS)
	@touch $(GENERATED_DIR)

# The main broker binary depends on generated code and all source.
# This should be the contents of pkg/ and the main file for correctness.
$(BROKER_BIN): $(GENERATED_DIR) $(SOURCE) $(DEPSRC)
	GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -ldflags "-X $(IMPORT_PATH)/pkg/version.Application=$(APPLICATION) -X $(IMPORT_PATH)/pkg/version.Version=$(VERSION) -X $(IMPORT_PATH)/pkg/version.GitCommit=$(COMMIT)" -o $@ ./cmd/broker

# The CRDs are auto generated and depend on the API source only.
$(CRD_DIR)/%: $(APISRC)
	@mkdir -p $(CRD_DIR)
	go run sigs.k8s.io/controller-tools/cmd/controller-gen crd paths=./pkg/apis/... output:dir=$(CRD_DIR)

# The TGZ archive relies on the archive directory.
$(ARCHIVE_TGZ): $(INSTALL_TARGETS)
	tar -czf $@ -C $(DESTDIR) .

# The ZIP archive relies on the archive directory.
$(ARCHIVE_ZIP): $(INSTALL_TARGETS)
	@cd $(DESTDIR); zip -r $@ .
	@mv $(DESTDIR)/$@ .

# The RPM build target archives the source directory and installs it
# along with the spec into the RPM build directory.  The spec is updated
# to contain the correct application and versioning.  The final RPM is
# copied back locally.
$(ARCHIVE_RPM): $(SOURCE)
	@mkdir -p $(HOME)/rpmbuild/SPECS
	@mkdir -p $(HOME)/rpmbuild/SOURCES
	cd ..; cp -a service-broker $(PACKAGE_BASENAME); tar czf $(PACKAGE_BASENAME).tar.gz $(PACKAGE_BASENAME); mv $(PACKAGE_BASENAME).tar.gz $(HOME)/rpmbuild/SOURCES
	sed -e "s,0\.0\.0,$(VERSION),g" -e "s,99999,$(REVISION),g" -e "s,couchbase-service-broker,$(APPLICATION),g" redhat/servicebroker.spec > $(HOME)/rpmbuild/SPECS/servicebroker.spec
	rpmbuild -ba $(HOME)/rpmbuild/SPECS/servicebroker.spec
	cp $(HOME)/rpmbuild/RPMS/x86_64/$(ARCHIVE_RPM) .

# The DEB build target creates a debian archive.
$(ARCHIVE_DEB): $(SOURCE)
	sed -i -e "s,0\.0\.0,$(VERSION),g" -e "s,99999,$(REVISION),g" -e "s,couchbase-service-broker,$(APPLICATION),g" debian/changelog
	sed -i -e "s,couchbase-service-broker,$(APPLICATION),g" debian/control
	sed -i -e "s,couchbase-service-broker,$(APPLICATION),g" debian/copyright
	sed -i -e "s,/usr/local,$(PREFIX),g" debian/rules
	cd ..; cp -a service-broker $(PACKAGE_BASENAME_DEB); tar czf $(PACKAGE_BASENAME_DEB).orig.tar.gz $(PACKAGE_BASENAME_DEB)
	DEB_BUILD_OPTIONS=nocheck debuild --prepend-path=$(GO_BINARY_PATH) -uc -us
	mv ../$(ARCHIVE_DEB) .

# Default make target for install files copies them over with existing permissions.
$(INSTALL_SHARE)/%: %
	@mkdir -p `dirname $@`
	cp $< $@

# Default make target for binary files copies them over with existing permissions.
$(INSTALL_BIN)/%: build/bin/%
	@mkdir -p `dirname $@`
	cp $< $@

# Default make target for docker files does a translation of the binary paths.
$(INSTALL_SHARE)/Dockerfile: Dockerfile
	@mkdir -p `dirname $@`
	sed -e "s,build/bin,bin,g" $< > $@
