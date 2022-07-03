.SILENT:
.EXPORT_ALL_VARIABLES:
.PHONY:

APP_NAME := crud-generator
VERSION := 1.0.0
TAG := $(shell git describe --abbrev=0 --tags)
COMMIT_DATE := $(shell git log -1 --date=format:"%y-%m-%dT%TZ" --format="%ad")
COMMIT_LAST := $(shell git rev-parse HEAD)

build-windows:
	cd cmd; fyne package -os windows --executable ../build/$(APP_NAME).exe

build-linux:
	cd cmd; fyne package -os linux --executable ../build/$(APP_NAME)