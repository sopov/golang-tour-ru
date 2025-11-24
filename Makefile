APP  = tour
SRC  = ./tour
DIST = dist

OS_LIST = $(shell go tool dist list | egrep '^(linux|darwin|windows)/')

all: build

build:
	mkdir -p $(DIST)
	@for os_arch in $(OS_LIST); do \
		os=$${os_arch%/*}; \
		arch=$${os_arch#*/}; \
		ext=""; \
		[ "$$os" = "windows" ] && ext=".exe"; \
		out="$(DIST)/$(APP)-$$os-$$arch$$ext"; \
		echo "==> Building $$out"; \
		GOOS=$$os GOARCH=$$arch CGO_ENABLED=0 \
			go build -buildvcs=false -o "$$out" $(SRC) || exit 1; \
	done

tag:
	@tag=$$(date -u +"%Y%m%d"); \
	remote=$$(git rev-parse --abbrev-ref --symbolic-full-name @{u} 2>/dev/null | cut -d/ -f1); \
	if [ -z "$$remote" ]; then \
		echo "no upstream for current branch, fallback to origin"; \
		remote=origin; \
	fi; \
	echo "Tag: $$tag -> $$remote"; \
	git tag $$tag; \
	git push $$remote $$tag

clean:
	rm -rf $(DIST)

.PHONY: all build tag clean
