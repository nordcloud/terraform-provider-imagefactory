TEST?=$$(go list ./... | grep -v 'vendor')

HOSTNAME=nordcloud.com
NAMESPACE=klarity
NAME=imagefactory
BINARY=terraform-provider-${NAME}

VERSION=v1.0.0
GOOS=darwin
GOARCH=amd64

default: install

.ONESHELL:
build:
	$(eval V := $(shell echo ${VERSION} | tr -d 'v'))
	GOOS=${GOOS} GOARCH=${GOARCH} go build -o ./bin/${BINARY}_${VERSION}
	echo "mkdir -p ~/.terraform.d/plugins/${HOSTNAME}/${NAMESPACE}/${NAME}/$(V)/${GOOS}_${GOARCH}" > ./bin/install.sh
	echo "cp ${BINARY}_${VERSION} ~/.terraform.d/plugins/${HOSTNAME}/${NAMESPACE}/${NAME}/$(V)/${GOOS}_${GOARCH}/${BINARY}" >> ./bin/install.sh
	chmod +x ./bin/install.sh

install: build
	cd bin
	./install.sh

generateDoc:
	go run github.com/hashicorp/terraform-plugin-docs/cmd/tfplugindocs
	
test:
	go test -i $(TEST) || exit 1
	echo $(TEST) | xargs -t -n4 go test $(TESTARGS) -timeout=30s -parallel=4

testacc:
	TF_ACC=1 go test $(TEST) -v $(TESTARGS) -timeout 120m