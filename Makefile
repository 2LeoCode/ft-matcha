SHELL = bash
NAME = matcha

all: build

generate:
	sqlc generate
	templ generate

build: generate
	go build .

run: generate
	go run .

clean:
	rm -rf ./pkg ./database/sql
	find ./ -type f -name '*_templ.go' -delete

fclean: clean
	rm -rf $(NAME)
