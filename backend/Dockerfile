FROM golang:1.18-alpine

RUN adduser -s /bin/bash appuser
USER appuser
WORKDIR /home/appuser

COPY . .
RUN go mod download
RUN go build -o appritstoreback

EXPOSE 9001

CMD [ "appritstoreback" ]
