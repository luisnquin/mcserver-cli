.PHONY: build, dist

build:
	@go build -o ./build/mcserver ./src/main.go

run: 
	@./build/mcserver

clean:
	@rm -rf ./build
	@mkdir build

dist:
	@cp build/mcserver debian/usr/local/bin
	@dpkg -b ./debian ./dist

install:
	@sudo dpkg -i ./dist/mc-server_0.0.1_amd64.deb
	@echo Installed

uninstall:
	@sudo apt purge mc-server
	@echo "Uninstalled"