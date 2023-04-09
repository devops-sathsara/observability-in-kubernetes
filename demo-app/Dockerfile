FROM golang:1.19-alpine

RUN apk --no-cache add curl git make musl-dev

WORKDIR /app

COPY go.mod .
COPY go.sum .
RUN go mod download

COPY . .
RUN make build

ENV OTEL_EXPORTER_OTLP_ENDPOINT=http://otel-collector:4317
ENV OTEL_SERVICE_NAME=my-go-app

CMD ["./server"]
