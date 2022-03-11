OPEN_API:=docs/openapi.yaml
GOPATH:=$(shell go env GOPATH)

build-react:
	cd ui/react; npm install; npm run build

build-angular:
	cd ui/angular; npm install; npm run build

generate-code:
	$(GOPATH)/bin/oapi-codegen -generate gin -package api $(OPEN_API) > pkg/api/petsalone.gen.go
	$(GOPATH)/bin/oapi-codegen -generate 'gin,types,skip-prune,spec' -package schema $(OPEN_API) > pkg/schema/petsalone_schema.gen.go
	$(GOPATH)/bin/oapi-codegen -generate 'gin,types,skip-prune,spec,client' -package client $(OPEN_API) > pkg/client/petsalone_client.gen.go

test:
	go clean -testcache;\
	go test -v -cover ./...

run-react:
	go run cmd/petsalone-react/main.go

run-angular:
	go run cmd/petsalone-angular/main.go