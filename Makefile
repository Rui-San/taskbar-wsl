NAME := TaskbarWSL
export GOOS := windows
export GOARCH := amd64

.PHONY: build build-noconsole clean

# build during development
build:
	go build -o $(NAME)-dev.exe .

# removes the console by using the flag specified by systray lib
build-noconsole:
	go build -ldflags "-H=windowsgui" -o $(NAME).exe .

# removes all previous exe files
clean:
	rm -f *.exe