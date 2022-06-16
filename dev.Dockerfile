FROM golang:1.18-buster as go_builder

RUN useradd -m -s /bin/bash appuser
USER appuser
WORKDIR /home/appuser

COPY ./backend/go.mod ./
COPY ./backend/go.sum ./

RUN go mod download

ADD ./backend ./

RUN go build -o appritstoreback

# Build the React application
FROM node:16 as node_builder

WORKDIR /usr/src/app

# install and cache app dependencies
COPY ./frontend/package*.json ./
ADD ./frontend/package.json /usr/src/app/package.json
RUN npm install

# Bundle app source
ADD ./frontend ./

# Final stage build, this will be the container with Go and React
FROM node:16

WORKDIR /usr/src/app
RUN mkdir web

COPY --from=go_builder /home/appuser .
COPY --from=node_builder /usr/src/app ./web
COPY ./start.sh /usr/src/app/start.sh
RUN chmod a+x /usr/src/app/start.sh

EXPOSE 9000 9001
CMD ./start.sh