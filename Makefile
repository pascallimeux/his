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
	@echo ' make init            Get from github and gerrit fabric-sdk-go, Go and Javascript mandatories libs.'
	@echo ' make build-swagger   Generate swagger code.'
	@echo ' make build-his-image Generate docker image for His and save it (./build/image/his.tar).'
	@echo ' make build-ui-image  Generate docker image for His_ui and save it (./build/image/his-ui.tar).'
	@echo ' make clean           Remove all generated code.'
	@echo ' make update-vendor   Update vendor libs.'
	@echo ' make update-sdk      Update fabric-sdk-go.'
	@echo
	@echo ' ## Run Commands'
	@echo ' make start-hp  Start docker hyperledger test environment.'
	@echo ' make stop-hp   Stop  docker hyperledger test environment.'
	@echo ' make clean-hp  Clean docker hyperledger test environment.'
	@echo ' make start-his Start docker HIS container.'
	@echo ' make stop-his  Stop  docker HIS container.'
	@echo ' make start-ui  Start docker HIS and UI containers.'
	@echo ' make stop-ui   Stop  docker HIS ans UI containers.'
	@echo ' make clean-his Clean docker HIS and UI images, container and server keys.'
	@echo ' make log-his   Display the His log.'
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

build-ui:
	@sh ./scripts/update_config_ui.sh
	@sh ./scripts/build_ui.sh

all: integration-test

init: init-ui
	@sh ./scripts/dependencies.sh
	@sh ./scripts/get_fabric_sdk.sh

init-ui:
	@sh ./scripts/init_ui.sh

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

start-ui:
	@sh ./scripts/start_ui.sh

stop-ui: stop-ui
	@sh ./scripts/stop_ui.sh
		
clean-his: stop-his
	@sh ./scripts/clean_his.sh

clean-ui: stop-ui
	@sh ./scripts/clean_ui.sh

log-his:
	docker logs hisv1 -f --tail=100

build-his-image: build-his
	@sh ./scripts/generate_keys.sh
	@sh ./scripts/update_config_his.sh
	@sh ./scripts/build_swagger.sh
	@sh ./scripts/build_his_image.sh
	@sh ./scripts/save_his_image.sh

build-ui-image: build-ui
	@sh ./scripts/build_ui_image.sh
	@sh ./scripts/save_ui_image.sh

clean: clean-hp clean-swagger clean-ui clean-his

integration-test: start-hp test stop-hp

test:
	@echo "> Start HIS intregration tests."
	@cd $(GOPATH)/src/github.com/pascallimeux/his/his/api && go test -v
	@cd $(GOPATH)/src/github.com/pascallimeux/his/his/helpers && go test -v

build-swagger:
	@sh ./scripts/build_swagger.sh

start-swagger:
	@sh ./scripts/start_swagger.sh

stop-swagger:
	@sh ./scripts/stop_swagger.sh

clean-swagger: stop-swagger
	@sh ./scripts/clean_swagger.sh
