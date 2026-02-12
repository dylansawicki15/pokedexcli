# pokedexcli

A tiny Go command-line Pokedex. It is a simple REPL-style app that shows a prompt and reads commands. This is intentionally small and a work in progress.

## Requirements

- Go 1.23+ (matches the go.mod)

## Build

```bash
make build
```

This increments the build number and injects it into the binary.

## Run

```bash
./pokedexcli
```

You should see a `Pokedex >` prompt. Type a command and press Enter.

## Notes

This project is still growing. Expect rough edges while the command set fills out.
