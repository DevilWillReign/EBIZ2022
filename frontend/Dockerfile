FROM node:16

ARG API_BASE_URL=http://localhost:9000/api/v1
ARG PORT=9001

ENV PORT ${PORT}
ENV API_BASE_URL ${API_BASE_URL}

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
