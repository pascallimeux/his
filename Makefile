.DEFAULT_GOAL := build

include Makefile.variables

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

.PHONY: help
help:
	@echo 'Management commands for cicdtest:'
	@echo
	@echo 'Usage:'
	@echo ' ## Build Commands'
	@echo ' make tag-build Add git tag for latest build.'
	@echo
	@echo ' ## Generator Commands'
	@echo ' make generate Run code generator for project.'
	@echo
	@echo ' ## Develop / Test Commands'
	@echo ' make vendor Install dependencies.'
	@echo ' make format Run code formatter.'
	@echo ' make check Run static code analysis (lint).'
	@echo ' make test Run tests on project.'
	@echo ' make cover Run tests and capture code coverage metrics on project.'
	@echo ' make clean Clean the directory tree of produced artifacts.'
	@echo
	@echo ' ## Utility Commands'
	@echo ' make setup Configures Minishfit/Docker directory mounts.'
	@echo

all: integration-test

.PHONY: depend
depend:
	go get -u github.com/kardianos/govendor && go get github.com/gorilla/mux && go get github.com/op/go-logging
	cd $GOPATH/src/github.com/hyperledger/ && git clone https://gerrit.hyperledger.org/r/fabric-sdk-go

.PHONY: vendor
vendor:
	echo "Make vendor Install dependencies."
	cd $(GOPATH)/src/github.com/hyperledger/fabric-sdk-go && git pull
	cd $(GOPATH)/src/github.com/pascallimeux/his
	sudo rm -R vendor
	govendor init
	govendor add +external

.PHONY: start-hp
start-hp:
	echo "Start Hyperledger test environment."
	cd $(GOPATH)/src/github.com/hyperledger/fabric-sdk-go/test/fixtures && docker-compose -f docker-compose.yaml up --force-recreate -d
	docker ps

.PHONY: stop-hp
stop-hp:
	echo "Stop Hyperledger test environment."
	cd $(GOPATH)/src/github.com/hyperledger/fabric-sdk-go/test/fixtures && docker-compose -f docker-compose.yaml stop
	docker ps

.PHONY: clean-hp
clean-hp:
	echo "Clean Hyperledger test environment."
	sh ./scripts/cleanHP.sh

image:
	echo "Build HIS docker image."
	cp ./fixtures/config/config.yaml ./fixtures/config/config_prod.yaml
	sudo sed -i "s/IPPEER0/$(IPPEER0)/g"  ./fixtures/config/config_prod.yaml
	sudo sed -i "s/PORTPEER0/$(PORTPEER0)/g"  ./fixtures/config/config_prod.yaml
	sudo sed -i "s/PORTEVT0/$(PORTEVT0)/g"  ./fixtures/config/config_prod.yaml
	sudo sed -i "s/IPPEER1/$(IPPEER1)/g"  ./fixtures/config/config_prod.yaml
	sudo sed -i "s/PORTPEER1/$(PORTPEER1)/g"  ./fixtures/config/config_prod.yaml
	sudo sed -i "s/PORTEVT1/$(PORTEVT1)/g"  ./fixtures/config/config_prod.yaml
	sudo sed -i "s/IPORDERER0/$(IPORDERER0)/g"  ./fixtures/config/config_prod.yaml
	sudo sed -i "s/PORTORDERER0/$(PORTORDERER0)/g"  ./fixtures/config/config_prod.yaml
	sudo sed -i "s/IPCA0/$(IPCA0)/g"  ./fixtures/config/config_prod.yaml
	sudo sed -i "s/PORTCA0/$(PORTCA0)/g"  ./fixtures/config/config_prod.yaml
	cd $(GOPATH)/src/github.com/pascallimeux/his && go build his.go
	openssl genrsa -out server.key 2048
	openssl req -new -x509 -sha256 -key server.key -days 3650 -subj "/C=FR/ST=France/L=Grenoble/O=Orange/OU=OLS/CN=orange-labs.fr" -out server.crt
	docker build -t his .
	docker images


.PHONY: start-his
start-his: start-hp
	echo "Start HIS docker container."
	docker run -d -p 8000:8000 --name hisv1 his


.PHONY: stop-his
stop-his:
	echo "Stop HIS docker container."
	docker kill hisv1


.PHONY: integration-test
integration-test: depend start-hp test stop-hp


.PHONY: test
test:
	echo "Start HIS intregration tests."
	cd $(GOPATH)/src/github.com/pascallimeux/his/api && go test -v
	cd $(GOPATH)/src/github.com/pascallimeux/his/helpers && go test -v

swagger-init:
	swagger init spec \
      --title "his" \
      --description "Hyperledger Interface Server" \
      --version 1.0.0 \
      --scheme http
	swagger generate spec -o ./swagger.json -i ./swagger.yml
	go get -u -f ./...
	swagger generate server -f ./swagger.json -A his

swagger-build:
	swagger generate spec -o ./swagger.json -i ./swagger.yml
	swagger generate server -f ./swagger.json -A his

swagger-clean:
	rm swagger.json && sudo rm -R cmd && sudo rm -R restapi

swagger-start:
	#swagger serve --port=3000 --host=127.0.0.1 swagger.json --base-path=/swagger-ui
	go run ./cmd/his-server/main.go --host=192.168.20.77 --port=3000