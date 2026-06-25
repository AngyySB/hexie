# hexie

A tiny CLI tool that previews a hex color directly in your terminal.

## Usage

```
hexie <hex-color>
```

The `#` prefix is optional:

```sh
hexie "#ff6347"
hexie ff6347
```

This prints a small colored block using your terminal's true-color support.

## Install

**Via `go install`** (adds to `$GOPATH/bin`, make sure that's in your `$PATH`):

```sh
go install github.com/AngyySB/hexie@latest
```

**Via Make** (installs to `/usr/local/bin` by default):

```sh
git clone https://github.com/AngyySB/hexie
cd hexie
sudo make install
```

To install to a custom prefix (e.g. `~/.local`):

```sh
make install PREFIX=~/.local
```

Then make sure the target directory is in your `PATH`. Add this to your shell config (`.bashrc`, `.zshrc`, `config.fish`, etc.):

```sh
# bash / zsh
export PATH="$HOME/.local/bin:$PATH"

# fish
fish_add_path ~/.local/bin
```

To uninstall:

```sh
sudo make uninstall
# or
make uninstall PREFIX=~/.local
```

## Requirements

- A terminal with 24-bit true-color support (most modern terminals)

## License

MIT
