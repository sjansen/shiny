.PHONY:  default  test  test-release

default: test

test:
	@scripts/run-all-tests
	@echo ========================================
	@git grep TODO  -- '**.go' || true
	@git grep FIXME -- '**.go' || true

test-release:
	git stash -u -k
	goreleaser release --rm-dist --skip-publish
	-git stash pop
