ANSI_RED="\033[0;31m"
ANSI_GREEN="\033[0;32m"
ANSI_BLUE="\033[0;34m"
ANSI_RESET="\033[0m"

ifneq ("$(wildcard /usr/local/opt/coreutils/libexec/gnubin/echo)","")
	ECHO="/usr/local/opt/coreutils/libexec/gnubin/echo"
else
	ECHO="/bin/echo"
endif

.PHONY: all test test-html docs

all: cmd/easy-lexml/easy-lexml
bindata.go: res/*
	@$(ECHO) -e $(ANSI_GREEN)"["$@"] Packing bin-data..."$(ANSI_RESET)
	go-bindata -pkg easyLexML res/
cmd/easy-lexml/easy-lexml: *.go cmd/easy-lexml/*.go bindata.go
	@$(ECHO) -e $(ANSI_GREEN)"["$@"] Fixing imports..."$(ANSI_RESET)
	cd cmd/easy-lexml/ && goimports -w .
	@$(ECHO) -e $(ANSI_GREEN)"["$@"] Formatting code..."$(ANSI_RESET)
	cd cmd/easy-lexml/ && go fmt
	@$(ECHO) -e $(ANSI_GREEN)"["$@"] Compiling code..."$(ANSI_RESET)
	cd cmd/easy-lexml/ && go build -o $@
	@$(ECHO) -e $(ANSI_BLUE)"["$@"] Finished target $@"$(ANSI_RESET)