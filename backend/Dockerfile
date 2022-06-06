FROM golang:1.18-buster

ARG FRONT_HOST=http://localhost:9001
ARG API_HOST=localhost
ARG API_PORT=9000
ARG SESSION_SECRET
ARG JWT_SECRET
ARG GOOGLE_OAUTH_CLIENT_ID
ARG GOOGLE_OAUTH_CLIENT_SECRET
ARG GH_OAUTH_CLIENT_ID
ARG GH_OAUTH_CLIENT_SECRET
ARG GL_OAUTH_CLIENT_ID
ARG GL_OAUTH_CLIENT_SECRET

RUN useradd -m -s /bin/bash appuser
USER appuser
WORKDIR /home/appuser

COPY go.mod ./
COPY go.sum ./

RUN go mod download

COPY . .

RUN go build -o appritstoreback

CMD [ "appritstoreback" ]
