OS := $(shell uname -s)

ifeq ($(OS),Windows_NT)
	BINARY_NAME = goapi.exe
	RM = powershell Remove-Item -Path
else
	BINARY_NAME = goapi
	RM = rm -rf
endif

build:
	@echo "Building..."
	@go build -o bin/$(BINARY_NAME)
	@echo "Done."

run: build
	@echo "Running..."
	@./bin/$(BINARY_NAME)

test: 
	@go test -v ./...

clean:
	@echo "Cleaning up..."
	@go clean
	@$(RM) bin
	@echo "Done."

watch:
	@echo "Watching..."
	@CompileDaemon -directory=./ -build="go build -o bin/$(BINARY_NAME)" -command="./bin/$(BINARY_NAME)" -exclude-dir=.git -exclude-dir=postgresql -include="*.html"
