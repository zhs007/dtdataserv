
FROM golang:1.12 as builder

MAINTAINER zerro "zerrozhao@gmail.com"

WORKDIR /src/dtdataserv

COPY ./go.* /src/dtdataserv/

RUN go mod download

COPY . /src/dtdataserv

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o dtdataserv . \
    && mkdir /app/dtdataserv \
    && mkdir /app/dtdataserv/dat \
    && mkdir /app/dtdataserv/logs \
    && mkdir /app/dtdataserv/cfg \
    && cp -r www /app/dtdataserv/www \
    && cp ./dtdataserv /app/dtdataserv/ \
    && cp cfg/config.yaml.default /app/dtdataserv/cfg/config.yaml

FROM alpine
WORKDIR /app/dtdataserv
COPY --from=builder /app/dtdataserv /app/dtdataserv
CMD ["./dtdataserv", "start"]
