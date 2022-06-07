FROM node:16

ARG ARG_API_BASE_URL=http://localhost:9000/api/v1
ARG ARG_FRONT_PORT=9001

ENV PORT $ARG_FRONT_PORT
ENV API_BASE_URL $ARG_API_BASE_URL

WORKDIR /usr/src/app

# install and cache app dependencies
COPY package*.json ./
ADD package.json /usr/src/app/package.json
RUN npm install

# Bundle app source
COPY . .

EXPOSE ${PORT}

# start app
CMD ["npm", "start"]
