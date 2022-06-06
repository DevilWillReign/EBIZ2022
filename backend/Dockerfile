FROM golang:1.18-alpine

RUN useradd -m -s /bin/bash appuser
USER appuser
WORKDIR /home/appuser

COPY . .
RUN go mod download
RUN go build -o appritstoreback

EXPOSE 9001

CMD [ "appritstoreback" ]
