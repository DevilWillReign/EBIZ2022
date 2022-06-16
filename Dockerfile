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

ARG ARG_FRONT_HOST=http://localhost:9001
ARG ARG_API_PORT=9000
ARG ARG_API_HOST_CALLBACK
ARG ARG_SESSION_SECRET
ARG ARG_JWT_SECRET
ARG ARG_GOOGLE_OAUTH_CLIENT_ID
ARG ARG_GOOGLE_OAUTH_CLIENT_SECRET
ARG ARG_GH_OAUTH_CLIENT_ID
ARG ARG_GH_OAUTH_CLIENT_SECRET
ARG ARG_GL_OAUTH_CLIENT_ID
ARG ARG_GL_OAUTH_CLIENT_SECRET

ENV FRONT_HOST $ARG_FRONT_HOST
ENV API_HOST_CALLBACK $ARG_API_HOST_CALLBACK
ENV API_PORT $ARG_API_PORT
ENV SESSION_SECRET $ARG_SESSION_SECRET
ENV JWT_SECRET $ARG_JWT_SECRET
ENV GOOGLE_OAUTH_CLIENT_ID $ARG_GOOGLE_OAUTH_CLIENT_ID
ENV GOOGLE_OAUTH_CLIENT_SECRET $ARG_GOOGLE_OAUTH_CLIENT_SECRET
ENV GH_OAUTH_CLIENT_ID $ARG_GH_OAUTH_CLIENT_ID
ENV GH_OAUTH_CLIENT_SECRET $ARG_GH_OAUTH_CLIENT_SECRET
ENV GL_OAUTH_CLIENT_ID $ARG_GL_OAUTH_CLIENT_ID
ENV GL_OAUTH_CLIENT_SECRET $ARG_GL_OAUTH_CLIENT_SECRET
ENV PROFILE=PROD

ARG ARG_API_BASE_URL=http://localhost:9000/api/v1
ARG ARG_FRONT_PORT=9001

ENV PORT $ARG_FRONT_PORT
ENV REACT_APP_API_BASE_URL $ARG_API_BASE_URL

WORKDIR /usr/src/app
RUN mkdir web

COPY --from=go_builder /home/appuser .
COPY --from=node_builder /usr/src/app ./web
COPY ./start.sh /usr/src/app/start.sh
RUN chmod +x /usr/src/app/start.sh

EXPOSE ${PORT}
CMD ./start.sh
