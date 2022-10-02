package = github.com/hambrosia/algorithms-golang

format:
	go fmt $(package)...

validate: format
	go test -v $(package)...

add-hook:
	git config core.hooksPath .githooks