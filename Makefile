
all:
	make install
	make proto
	make grapi_pkgw

gen_version:
	@cd ./scripts && bash ./change_project_version.sh

server:
	@go build -o ${OUT} ./cmd/server/main.go

init:
	@bash ./scripts/change_base_import.sh

proto: ## Generate source code from protos
	@bash ./scripts/codegen/generate_code_from_proto.sh

install:
	@bash ./scripts/install_dev_env/install_dev_env.sh

grapi_pkg:
	@bash ./scripts/others/install_grapi_pkg.sh

clean:
	@bash ./scripts/codegen/clean/clean.sh

.PHONY: gen_version server init proto install grapi_pkg clean