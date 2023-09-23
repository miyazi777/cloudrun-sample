# build
FROM golang:1.21 as builder

WORKDIR /app

COPY go.mod ./
## COPY go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 go build -o main main.go
#RUN go env
#ENTRYPOINT ["sleep", "infinity"]
#EXPOSE 8080
#CMD [ "/main"]

# deploy
FROM alpine:latest as dev
RUN apk add --no-cache ca-certificates
WORKDIR /app
#
COPY --from=builder /app/main .
#USER 1001
EXPOSE 8080
CMD [ "/app/main" ]