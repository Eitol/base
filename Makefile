new_name?="invalid"

all:
	make install_dev_env
	make proto
	make grapi_pkg

gen_version:
	@cd ./scripts && bash ./change_project_version.sh

server:
	@go build -o ${OUT} ./cmd/server/main.go

init:
	@python3 ./scripts/module_name_changer.py ${new_name}

proto: ## Generate source code from protos
	@bash ./scripts/codegen/generate_code_from_proto.sh

install_dev_env:
	@bash ./scripts/install_dev_env/install_dev_env.sh

grapi_pkg:
	@bash ./scripts/others/install_grapi_pkg.sh

clean:
	@bash ./scripts/codegen/clean/clean.sh

.PHONY: gen_version server init proto install_dev_env grapi_pkg clean