GOHOSTOS:=$(shell go env GOHOSTOS)
GOPATH:=$(shell go env GOPATH)
VERSION=$(shell git describe --tags --always)

ifeq ($(GOHOSTOS), windows)
	#the `find.exe` is different from `find` in bash/shell.
	#to see https://docs.microsoft.com/en-us/windows-server/administration/windows-commands/find.
	#changed to use git-bash.exe to run find cli or other cli friendly, caused of every developer has a Git.
	#Git_Bash= $(subst cmd\,bin\bash.exe,$(dir $(shell where git)))
	Git_Bash=$(subst \,/,$(subst cmd\,bin\bash.exe,$(dir $(shell where git))))
	COM_PROTO_FILES=$(shell $(Git_Bash) -c "find ./com -name *.proto")
	ERR_PROTO_FILES=$(shell $(Git_Bash) -c "find ./err -name *.proto")
	ALL_PROTO_FILES=$(shell $(Git_Bash) -c "find . -path ./third_party -prune -o -name '*.proto' -print")
else
	COM_PROTO_FILES=$(shell find ./com -name *.proto)
	ERR_PROTO_FILES=$(shell find ./err -name *.proto)
	ALL_PROTO_FILES=$(shell find . -path ./third_party -prune -o -name '*.proto' -print)
endif

.PHONY: com
# generate api proto
com:
	protoc --proto_path=./ \
	       --proto_path=third_party \
 	       --go_out=paths=source_relative:./ \
 	       --go-http_out=paths=source_relative:./ \
 	       --go-grpc_out=paths=source_relative:./ \
				 --validate_out=paths=source_relative,lang=go:./ \
	       $(COM_PROTO_FILES)
	clang-format -i $(COM_PROTO_FILES)

.PHONY: err
# generate api proto
err:
	protoc --proto_path=./ \
	       --proto_path=third_party \
 	       --go_out=paths=source_relative:./ \
 	       --go-http_out=paths=source_relative:./ \
 	       --go-grpc_out=paths=source_relative:./ \
				 --validate_out=paths=source_relative,lang=go:./ \
	       $(ERR_PROTO_FILES)
	clang-format -i $(ERR_PROTO_FILES)

.PHONY: all
# generate api proto
all:
	protoc --proto_path=./ \
	       --proto_path=third_party \
 	       --go_out=paths=source_relative:./ \
 	       --go-http_out=paths=source_relative:./ \
 	       --go-grpc_out=paths=source_relative:./ \
				 --validate_out=paths=source_relative,lang=go:./ \
	       --openapi_out=fq_schema_naming=true,default_response=false:. \
	       $(ALL_PROTO_FILES)
	clang-format -i $(ALL_PROTO_FILES)