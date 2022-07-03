.SILENT:
.EXPORT_ALL_VARIABLES:
.PHONY:

APP_NAME := crud-generator
VERSION := v1.0.0
TAG := $(shell git describe --abbrev=0 --tags)
COMMIT_DATE := $(shell git log -1 --date=format:"%y-%m-%dT%TZ" --format="%ad")
COMMIT_LAST := $(shell git rev-parse HEAD)

default:
	# Run "build-windows" to build the application for windows.
	# Run "build-linux" to build the application for linux.
	# Run "build-darwin" to build the application for darwin.

build-all: build-windows build-linux build-darwin

build-windows:
	fyne-cross windows -output $(APP_NAME).exe ./cmd

build-linux:
	fyne-cross linux -output $(APP_NAME) ./cmd

build-darwin:
	fyne-cross darwin -app-id $(APP_NAME).$(VERSION) -output $(APP_NAME) ./cmd