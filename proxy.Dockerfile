FROM node:16 as node_builder

ARG ARG_API_BASE_URL=http://localhost/api/v1
ARG ARG_API_HOST_CALLBACK=localhost

ENV REACT_APP_API_BASE_URL $ARG_API_BASE_URL
ENV API_HOST_CALLBACK $ARG_API_HOST_CALLBACK

WORKDIR /usr/src/app

# install and cache app dependencies
COPY ./frontend/package*.json ./
ADD ./frontend/package.json /usr/src/app/package.json
RUN npm install

# Bundle app source
ADD ./frontend ./

RUN npm run build

FROM nginx:latest

WORKDIR /var/www/apprit

COPY --from=node_builder /usr/src/app/build html

COPY nginx-prod.conf /etc/nginx/templates/default.conf.template

EXPOSE 80