FROM golang:alpine AS builder

# ENV GO111MODULE=on \
#     CGO_ENABLED=0

ENV GOPROXY=direct
ENV GOSUMDB=off
ENV CGO_ENABLED=0

WORKDIR /build

COPY go.mod go.sum ./
RUN go mod download

COPY . .
# RUN go mod tidy
RUN go build --ldflags "-s -w -extldflags -static" -o main .

FROM alpine:latest
# RUN apk update && apk add --no-cache tzdata \
#     && cp /usr/share/zoneinfo/Asia/Jakarta /etc/localtime \
#     && echo "Asia/Jakarta" > /etc/timezone
ENV TZ=Asia/Jakarta

WORKDIR /www

COPY --from=builder /build/main /www/
COPY --from=builder /build/public/ /www/public/
COPY --from=builder /build/storage/ /www/storage/
COPY --from=builder /build/resources/ /www/resources/
COPY --from=builder /build/.env.production /www/.env

ENTRYPOINT ["/www/main"]