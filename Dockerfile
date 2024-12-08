
FROM golang:1.22 as builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN go build -o /aggregator

FROM gcr.io/distroless/base-debian11
COPY --from=builder /aggregator /aggregator

EXPOSE 8081

CMD ["/aggregator"]