FROM node:22-alpine AS ts-build

WORKDIR /usr/src/app
COPY package.json tsconfig.json yarn.lock ts/ .
RUN npm i -g yarn@1.22.22
RUN yarn
RUN yarn build


FROM golang:1.23.4 AS go-build

WORKDIR /usr/src/app
COPY sqlc.json go.mod go.sum go/ sql/ .
RUN go install github.com/a-h/templ/cmd/templ@0.2.793 github.com/sqlc-dev/sqlc@1.27.0
RUN sqlc generate
RUN templ generate
RUN go build ./go -v -o /usr/bin/matcha


FROM alpine:3.21

WORKDIR /var/www
COPY .env css .
COPY --from=ts-build /usr/src/app/js .
COPY --from=go-build /usr/bin/matcha /usr/bin/matcha
CMD ["matcha"]
