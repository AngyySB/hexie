PREFIX ?= /usr/local

build:
	go build -o hexie .

install: build
	install -Dm755 hexie $(PREFIX)/bin/hexie

uninstall:
	rm -f $(PREFIX)/bin/hexie

clean:
	rm -f hexie

.PHONY: build install uninstall clean
