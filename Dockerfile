FROM golang:alpine AS builder

EXPOSE 80

WORKDIR /usr/local/src

COPY ["go.mod", "go.sum", "./"]

RUN ["go", "mod", "download"]

COPY ./ ./

RUN ["go", "build", "-o", "./bin/url-shortener", "./cmd/url-shortener"]

FROM alpine AS runner

COPY --from=builder /usr/local/src/bin/url-shortener ./

COPY ./config/config.yml ./config/config.yml

COPY ./migrations/scheme ./migrations/scheme

COPY .env .env

CMD [ "./url-shortener" ]