.PHONY: release
release:
	@echo "Building release version..."
	go mod tidy
	VERSION := $(shell git describe --tags --abbrev=0 | awk -v OFS='.' -F . '{$$3+=1; print}')
	git add .
	git commit -m "Version: $$(VERSION)"
	git tag $(VERSION)
	git push --tags