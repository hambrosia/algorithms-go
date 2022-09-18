package = github.com/hambrosia/agorithms-golang

format:
	go fmt $(package)...

validate: format
	go test -v $(package)...
