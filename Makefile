.PHONY: clean-proto full gopher protoc format-proto

GOPHER = 'ʕ◔ϖ◔ Fight!'
gopher:
	@echo ${GOPHER}

full:
	make format-proto
	make clean-proto
	make protoc
	@echo ${GOPHER}

protoc:
	protoc -I . --go_out=plugins=grpc:${GOPATH}/src/test_environment --go_opt paths=source_relative ./*/proto/*.proto

clean-proto:
	find ${CURDIR}/*/proto -name *.pb.go | xargs rm -f

format-proto:
	clang-format -i */proto/*.proto
