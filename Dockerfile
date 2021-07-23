FROM golang:1.15 AS build

RUN mkdir -p /meminero
WORKDIR /meminero

ADD go.mod go.mod
ADD go.sum go.sum
RUN go mod download

ADD . .

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo .

FROM scratch
COPY --from=build /meminero/meminero .
COPY --from=build /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
CMD ["./meminero", "run", "--config=/config/config.yml"]
