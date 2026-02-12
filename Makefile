BUILD_FILE := build_number.txt

.PHONY: build

build:
	@build=$$(cat $(BUILD_FILE) 2>/dev/null || echo 0); \
	build=$$((build + 1)); \
	echo $$build > $(BUILD_FILE); \
	go build -ldflags "-X main.build=$$build" -o pokedexcli .
