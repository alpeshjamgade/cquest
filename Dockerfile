FROM golang:1.22-alpine AS builder

RUN mkdir /app
COPY . /app
RUN ls
WORKDIR /app
RUN CGO_ENABLED=0 go build -o cquestApp ./cmd/main.go
RUN chmod +x cquestApp

FROM alpine:latest
RUN mkdir /app
COPY --from=builder /app/cquestApp /app
CMD [ "./app/cquestApp" ]