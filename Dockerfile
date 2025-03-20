FROM golang:1.24-alpine AS builder
WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o multimodal-search cmd/multimodal-search/main.go

FROM alpine:3.21.3 AS prod
WORKDIR /app/

ARG APP_ENV
ENV APP_ENV=$APP_ENV

COPY --from=builder /app/multimodal-search ./
COPY ./.config ./.config

RUN chmod +x multimodal-search

EXPOSE 3030
CMD [ "./multimodal-search" ]