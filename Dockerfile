# syntax=docker/dockerfile:1
FROM --platform=$BUILDPLATFORM golang:1.23.4 AS build
ARG APP
WORKDIR /src

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 make build-${APP}

FROM alpine:latest AS final
WORKDIR /cloud
ARG APP
ARG PORT
USER root

COPY --from=build /src/build/cloud.${APP} .
COPY --from=build /src/config.toml .

ENV ENTRYPOINT_CMD=./cloud.${APP}

ENTRYPOINT [ "sh", "-c", "$ENTRYPOINT_CMD" ]