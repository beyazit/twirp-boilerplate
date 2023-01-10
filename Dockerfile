FROM golang:1.19-alpine AS builder

WORKDIR /app
COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -ldflags "-s -w" -o /app/twirp-boilerplate

RUN apk add upx
RUN upx --ultra-brute /app/twirp-boilerplate

FROM scratch AS prod
COPY --from=builder /app/twirp-boilerplate /bin/twirp-boilerplate
EXPOSE 8080
ENTRYPOINT ["/bin/twirp-boilerplate"]