# Build the Go Binary.
FROM golang:1.18.6 as build_wpan-api
ENV CGO_ENABLED 0
ARG BUILD_REF
ENV GOPROXY https://goproxy.cn,direct

# Copy the source code into the container.
COPY . /wpan

# Build the service binary.
WORKDIR /wpan
RUN go build -ldflags "-X main.build=${BUILD_REF}"


# Run the Go Binary in Alpine.
FROM alpine:3.14
ARG BUILD_DATE
ARG BUILD_REF
COPY --from=build_wpan-api /wpan/wpan /wpan/wpan-api
WORKDIR /wpan
CMD ["./wpan-api"]

LABEL org.opencontainers.image.created="${BUILD_DATE}" \
      org.opencontainers.image.title="wpan-api" \
      org.opencontainers.image.authors="zcw" \
      org.opencontainers.image.source="https://imianba.com" \
      org.opencontainers.image.revision="${BUILD_REF}" \
      org.opencontainers.image.vendor="Wpan"