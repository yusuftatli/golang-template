API = api.yaml
BINARY_NAME=main


.PHONY: install_tools
install_tools: tools/swagger 

tools/swagger:
	$(call print-target)
	GOBIN=$(CURDIR)/tools go install github.com/go-swagger/go-swagger/cmd/swagger

.PHONY: models
models: tools/swagger
	$(call print-target)
	find ./models -type f -not -name '*_test.go' -delete
	./tools/swagger generate model --spec=docs/$(API)

run:
	go run  .

build:
	go run  .

install:
		go install