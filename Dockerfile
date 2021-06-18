FROM golang:1.15 AS build

RUN mkdir -p /smartbackend
WORKDIR /smartbackend

ADD go.mod go.mod
ADD go.sum go.sum
RUN go mod download

ADD . .

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo .

FROM scratch
COPY --from=build /smartbackend/smartbackend .
COPY --from=build /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
CMD ["./smartbackend", "run", "--config=/config/config.yml"]
