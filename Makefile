# Assign only if not already defined
PREFIX ?= /usr/local

BIN_DIR := $(PREFIX)/bin

.PHONY: build install uninstall

build:
	go build

install:
	cp date_picker $(BIN_DIR)/

uninstall:
	rm $(BIN_DIR)/date_picker
