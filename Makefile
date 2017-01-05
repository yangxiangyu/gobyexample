govendor:
	go get -u github.com/kardianos/govendor
	govendor init
	govendor add +external
build:
	ls dist || mkdir dist
	go build  -o dist/run main.go
clean:
	rm -rf dist/*