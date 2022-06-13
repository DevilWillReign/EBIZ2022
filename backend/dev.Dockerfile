FROM golang:1.18-buster

RUN useradd -m -s /bin/bash appuser
USER appuser
WORKDIR /home/appuser

COPY go.mod ./
COPY go.sum ./

RUN go mod download

COPY . .

RUN go build -o appritstoreback

EXPOSE 9000

CMD [ "appritstoreback" ]
