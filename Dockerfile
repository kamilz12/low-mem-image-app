####################   STAGE 1 – build   ####################
FROM --platform=$BUILDPLATFORM golang:1.22-alpine AS builder

ARG TARGETOS TARGETARCH
WORKDIR /src
COPY go.mod go.sum ./
RUN --mount=type=cache,target=/go/pkg/mod \
    go mod download

COPY . .
# Statyczna binarka + strip (-s -w) + CGO off
RUN CGO_ENABLED=0 GOOS=$TARGETOS GOARCH=$TARGETARCH \
    go build -ldflags="-s -w" -o /out/weather

# (opcjonalnie) kompresja UPX-em
RUN apk add --no-cache upx \
 && upx --lzma --best /out/weather

####################  STAGE 2 – runtime  ####################
FROM scratch
LABEL org.opencontainers.image.authors="Kamil Ziolkowski <s99752@pollub.edu.pl>"
# certy potrzebne dla HTTPS → 220 kB
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /out/weather /weather
EXPOSE 8000
ENTRYPOINT ["/weather"]
HEALTHCHECK --interval=30s --timeout=3s CMD wget --spider -q http://localhost:8000/ || exit 1

