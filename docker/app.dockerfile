FROM golang:1.20 AS builder
WORKDIR /build/
COPY .. .
RUN go mod download
RUN go build -o vtb_map_api main.go

FROM alpine:latest
WORKDIR /app/
COPY --from=builder /build/vtb_map_api .
COPY --from=builder /build/env/* ./env/
EXPOSE 8070
CMD ["./vtb_map_api"]