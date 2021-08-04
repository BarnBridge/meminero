FROM golang:1.16 AS build

RUN mkdir -p /meminero
WORKDIR /meminero


ARG SSH_PRIVATE_KEY
ENV GOPRIVATE=github.com/lacasian/

RUN mkdir -p ~/.ssh && umask 0077 && echo "${SSH_PRIVATE_KEY}" > ~/.ssh/id_rsa \
	&& git config --global url.ssh://git@github.com/.insteadOf https://github.com/ \
	&& ssh-keyscan github.com >> ~/.ssh/known_hosts

COPY go.mod go.sum ./

RUN go mod download

ADD . .

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo .

FROM scratch
COPY --from=build /meminero/meminero .
COPY --from=build /meminero/db/migrations db/migrations
COPY --from=build /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
CMD ["./meminero", "run", "--config=/config/config.yml"]
