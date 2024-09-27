APP_NAME := webshooter
VERSION := 1.3.1
ARCH := amd64
BUILD_DIR := build
PACKAGING_DIR := $(BUILD_DIR)/debian
BIN_DIR := /usr/local/bin
DEB_FILE := $(APP_NAME)_$(VERSION)_$(ARCH).deb

all: package

build: clean
	@echo "==> Building go binary..."
	GOOS=linux GOARCH=$(ARCH) go build -o $(BUILD_DIR)/$(APP_NAME) ./cmd/app

package: build
	@echo "==> Setting up the package structure..."
	mkdir -p $(PACKAGING_DIR)/DEBIAN
	mkdir -p $(PACKAGING_DIR)$(BIN_DIR)

	cp $(BUILD_DIR)/$(APP_NAME) $(PACKAGING_DIR)$(BIN_DIR)

	@echo "==> Generating control file..."
	echo "Package: $(APP_NAME)" > $(PACKAGING_DIR)/DEBIAN/control
	echo "Version: $(VERSION)" >> $(PACKAGING_DIR)/DEBIAN/control
	echo "Section: utils" >> $(PACKAGING_DIR)/DEBIAN/control
	echo "Priority: optional" >> $(PACKAGING_DIR)/DEBIAN/control
	echo "Architecture: $(ARCH)" >> $(PACKAGING_DIR)/DEBIAN/control
	echo "Maintainer: nunyabidnis" >> $(PACKAGING_DIR)/DEBIAN/control
	echo "Description: CLI utility to take screenshots and PDFs of target webpages using a headless browser" >> $(PACKAGING_DIR)/DEBIAN/control
	echo " see above" >> $(PACKAGING_DIR)/DEBIAN/control

	@echo "==> Building .deb package..."
	dpkg-deb --root-owner-group --build $(PACKAGING_DIR)

	@echo "==> Generating $(DEB_FILE)..."
	mv $(PACKAGING_DIR).deb $(BUILD_DIR)/$(DEB_FILE)

	@echo "==> Package build successfully at $(BUILD_DIR)/$(DEB_FILE)."

clean:
	@echo "==> Cleaning up..."
	rm -rf $(BUILD_DIR)/*
