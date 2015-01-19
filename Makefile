HOMEPAGE=https://github.com/justincampbell/tmux-pomodoro
PREFIX=/usr/local

COVERAGE_FILE = coverage.out

VERSION=1.1.1
TAG=v$(VERSION)

ARCHIVE=tmux-pomodoro-$(TAG).tar.gz
ARCHIVE_URL=$(HOMEPAGE)/archive/$(TAG).tar.gz

test: acceptance

release: tag sha

tag:
	git tag --force latest
	git tag | grep $(TAG) || git tag --message "Release $(TAG)" --sign $(TAG)
	git push origin
	git push origin --force --tags

pkg/$(ARCHIVE): pkg/
	wget --output-document pkg/$(ARCHIVE) $(ARCHIVE_URL)

pkg/:
	mkdir pkg

sha: pkg/$(ARCHIVE)
	shasum pkg/$(ARCHIVE)

install: build
	mkdir -p $(PREFIX)/bin
	cp -v bin/pomodoro $(PREFIX)/bin/pomodoro

uninstall:
	rm -vf $(PREFIX)/bin/pomodoro

coverage: unit
	go tool cover -html=$(COVERAGE_FILE)

acceptance: build
	bats test

build: dependencies unit
	go build -o bin/pomodoro

unit: dependencies
	go test -coverprofile=$(COVERAGE_FILE) -timeout 25ms

dependencies:
	go get -t
	go get golang.org/x/tools/cmd/cover

.PHONY: acceptance build coverage dependencies install release sha tag test uninstall unit
