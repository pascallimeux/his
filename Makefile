#
# Copyright OrangeLabs Inc. All Rights Reserved.
#
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#
#       http://www.apache.org/licenses/LICENSE-2.0
#
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.
#

#
# This MakeFile assumes that fabric, and fabric-ca were cloned and their docker
# images were created using the make docker command in the respective directories
#
# Supported Targets:
# all : runs unit and integration tests
# depend: installs test dependencies
# unit-test: runs all the unit tests
# integration-test: runs all the integration tests
# clean: stops docker conatainers used for integration testing
#

.DEFAULT_GOAL := bin
.PHONY: help
help:
	@echo
	@echo
	@echo ' --------------------------------------------------------'
	@echo '  Management commands for Hyperledeger Interface Service  '
	@echo ' --------------------------------------------------------'
	@echo
	@echo 'Usage:'
	@echo ' ## Build Commands'
	@echo ' make init          Get from github and gerrit fabric-sdk-go and Go mandatories libs.'
	@echo ' make build-his     Generate HIS binary (./build/bin/his).'
	@echo ' make build-swagger Generate swagger code.'
	@echo ' make build-image   Generate docker image for His and save it (./build/image/his.tar).'
	@echo ' make clean         Remove all generated code.'
	@echo ' make update-vendor Update vendor libs.'
	@echo ' make update-sdk    Update fabric-sdk-go.'
	@echo
	@echo ' ## Run Commands'
	@echo ' make start-hp  Start docker hyperledger test environment.'
	@echo ' make stop-hp   Stop  docker hyperledger test environment.'
	@echo ' make clean-hp  Clean docker hyperledger test environment.'
	@echo ' make start-his Start docker HIS container.'
	@echo ' make stop-his  Stop  docker HIS container.'
	@echo ' make clean-his Clean docker HIS image, container and server keys.'
	@echo
	@echo ' ## Test Commands'
	@echo ' make integration-test Run integration tests.'
	@echo ' make test             Run unit tests.'
	@echo
	@echo ' ## Swagger Commands'
	@echo ' make start-swagger Start  Rest server for swagger.'
	@echo ' make stop-swagger  Stop   Rest server for swagger.'
	@echo ' make clean-swagger Remove Swagger code.'
	@echo
	@echo

build-his:
	@sh ./scripts/build_his.sh

all: integration-test

init:
	@sh ./scripts/dependencies.sh
	@sh ./scripts/get_fabric_sdk.sh

update-vendor:
	@sh ./scripts/update_vendor.sh

update-sdk:
	@sh ./scripts/update_sdk.sh

start-hp:
	@sh ./scripts/start_hp.sh

stop-hp:
	@sh ./scripts/stop_hp.sh

clean-hp: stop-his stop-hp
	@sh ./scripts/clean_hp.sh

start-his: stop-his start-hp
	@sh ./scripts/start_his.sh

stop-his:
	@sh ./scripts/stop_his.sh

clean-his:
	@sh ./scripts/clean_his.sh

build-image: build-his
	@sh ./scripts/generate_keys.sh
	@sh ./scripts/update_config.sh
	@sh ./scripts/build_swagger.sh
	@sh ./scripts/build_image.sh
	@sh ./scripts/save_image.sh

clean: clean-hp clean-swagger clean-his

integration-test: start-hp test stop-hp

test:
	@echo "> Start HIS intregration tests."
	@cd $(GOPATH)/src/github.com/pascallimeux/his/api && go test -v
	@cd $(GOPATH)/src/github.com/pascallimeux/his/helpers && go test -v

build-swagger:
	@sh ./scripts/build_swagger.sh

start-swagger:
	@sh ./scripts/start_swagger.sh

stop-swagger:
	@sh ./scripts/stop_swagger.sh

clean-swagger: stop-swagger
	@sh ./scripts/clean_swagger.sh