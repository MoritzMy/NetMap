# --------- CONFIG ---------
APP_NAME := netmap
DAEMON_BIN := $(APP_NAME)-d
DAEMON_MAIN := ./cmd/daemon/main.go

SCRIPT_BIN := $(APP_NAME)
SCRIPT_MAIN := ./cmd/script/main.go

DOCKER_IMAGE := $(APP_NAME)-backend

# --------- PHONY ---------

.PHONY: install build-script build-daemon image

# --------- DEFAULT ---------

all: build

# --------- STEPS ---------

deps:
	go mod download

tidy:
	go mod tidy

build-script: deps tidy
	go build -o $(SCRIPT_BIN) $(SCRIPT_MAIN)
	chmod +x ./$(SCRIPT_BIN)
	sudo setcap cap_net_raw+ep ./$(SCRIPT_BIN)


build-daemon: deps tidy
	go build -o $(DAEMON_BIN) $(DAEMON_MAIN)

# --------- DOCKER ---------
image:
	docker build . -t $(DOCKER_IMAGE)