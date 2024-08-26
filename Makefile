
MAKE               := make --no-print-directory

DESCRIBE           := $(shell git describe --match "v*" --always --tags)
DESCRIBE_PARTS     := $(subst -, ,$(DESCRIBE))

VERSION_TAG        := $(word 1,$(DESCRIBE_PARTS))
COMMITS_SINCE_TAG  := $(word 2,$(DESCRIBE_PARTS))

VERSION            := $(subst v,,$(VERSION_TAG))
VERSION_PARTS      := $(subst ., ,$(VERSION))

MAJOR              := $(word 1,$(VERSION_PARTS))
MINOR              := $(word 2,$(VERSION_PARTS))
MICRO              := $(word 3,$(VERSION_PARTS))

NEXT_MICRO         := $(shell echo $$(($(MICRO)+1)))

CURRENT_VERSION_MICRO := $(MAJOR).$(MINOR).$(MICRO)
NEXT_VERSION_MICRO := $(MAJOR).$(MINOR).$(NEXT_MICRO)


DATE                = $(shell date +'%d.%m.%Y')
TIME                = $(shell date +'%H:%M:%S')
COMMIT             := $(shell git rev-parse HEAD)
AUTHOR             := $(firstword $(subst @, ,$(shell git show --format="%aE" $(COMMIT))))
BRANCH_NAME        := $(shell git rev-parse --abbrev-ref HEAD)

TAG_MESSAGE         = "$(TIME) $(DATE) $(AUTHOR) $(BRANCH_NAME)"
COMMIT_MESSAGE     := $(shell git log --format=%B -n 1 $(COMMIT))

CURRENT_TAG_MICRO  := "v$(CURRENT_VERSION_MICRO)"

# --- Version commands ---

.PHONY: current-version
current-version:
	@echo "$(CURRENT_VERSION_MICRO)"

.PHONY: next-version
next-version:
	@echo "$(NEXT_VERSION_MICRO)"

# --- Tag commands ---

.PHONY: tag-micro
tag-micro:
	@echo "$(CURRENT_TAG_MICRO)"

# -- Meta info ---

.PHONY: tag-message
tag-message:
	@echo "$(TAG_MESSAGE)"

.PHONY: commit-message
commit-message:
	@echo "$(COMMIT_MESSAGE)"

.PHONY: release
release:
	go mod tidy
	git add .
	git commit -m "Release v$(NEXT_VERSION_MICRO)"
	git tag -a $(CURRENT_TAG_MICRO) -m $(TAG_MESSAGE)
	git push origin $(BRANCH_NAME) $(CURRENT_TAG_MICRO)
	curl https://sum.golang.org/lookup/github.com/narsus81/gox@v$(NEXT_VERSION_MICRO)
	curl https://proxy.golang.org/github.com/narsus81/gox/@v/v$(NEXT_VERSION_MICRO).info
