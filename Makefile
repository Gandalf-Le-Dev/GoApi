BINARY_NAME = goapi.exe

build:
	@echo "Building..."
	@go build -o bin/$(BINARY_NAME)

run: build
	@echo "Running..."
	@./bin/$(BINARY_NAME)

test: 
	@go test -v ./...

clean:
	@echo "Cleaning up..."
	go clean
	@powershell Remove-Item -Path bin -Recurse -Force

watch:
	@echo "Watching..."
	@CompileDaemon -directory=./ -build="go build -o bin/$(BINARY_NAME)" -command="./bin/$(BINARY_NAME)" -exclude-dir=.git -exclude-dir=postgresql -include="*.html"