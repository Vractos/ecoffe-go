# Build
FROM golang:1.18 AS builder

WORKDIR /usr/src/app/

COPY go.mod go.sum ./
RUN go mod download && go mod verify

COPY . .
RUN CGO_ENABLED=0 go build -v -o ./bin/ecoffe-go

# Final Image
FROM alpine

WORKDIR /usr/src/app/

RUN apk add --no-cache curl bash

ARG app_env
ENV APP_ENV=${app_env}

# Postgres #
ARG postgres_user
ENV POSTGRES_USER=${postgres_user}

ARG postgres_password
ENV POSTGRES_PASSWORD=${postgres_password}

ARG postgres_db_name
ENV POSTGRES_DB_NAME=${postgres_db_name}

ARG postgres_host
ENV POSTGRES_HOST=${postgres_host}


COPY --from=builder /usr/src/app/bin/ecoffe-go ./

EXPOSE 80
ENTRYPOINT [ "./ecoffe-go" ]
