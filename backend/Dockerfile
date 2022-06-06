FROM golang:1.18-buster

RUN addgroup -S appgroup && adduser -G appgroup -S appuser
USER appuser
WORKDIR /home/appuser

COPY . .
RUN go mod download
RUN go build -o appritstoreback

EXPOSE 9001

CMD [ "appritstoreback" ]
