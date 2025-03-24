FROM node:22-alpine AS ts-build

WORKDIR /usr/src/app
COPY package.json tsconfig.json yarn.lock ts .
RUN npm i -g yarn@1.22.22
RUN yarn
RUN yarn build


FROM golang:1.23.4 AS go-build

WORKDIR /usr/src/app
COPY go.mod go.sum go sql .
RUN go install github.com/a-h/templ/cmd/templ@0.2.793 github.com/go-jet/jet/v2/cmd/jet@2.11.1
RUN templ generate
RUN go run ./tools/db-setup
RUN jet -source=sqlite -dsn="file://matcha.db" -path=./go/database/schema
RUN go build ./go -v -o /usr/bin/matcha


FROM alpine:3.21

WORKDIR /var/www
COPY .env css assets .
COPY --from=ts-build /usr/src/app/js .
COPY --from=go-build /usr/bin/matcha /usr/bin/matcha
CMD ["matcha"]
