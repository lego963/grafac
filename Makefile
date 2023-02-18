SHELL=/bin/bash

export PATH:=/usr/local/go/bin:~/go/bin/:$(PATH)

GOFMT_FILES?=$$(find . -name '*.go')
WORKDIR =$(patsubst %/,%,$(dir $(realpath $(lastword $(MAKEFILE_LIST)))))
