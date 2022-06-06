FROM node:16

WORKDIR /usr/src/app

# install and cache app dependencies
COPY ./frontend/package*.json ./
ADD ./frontend/package.json /usr/src/app/package.json
RUN npm install

ENV PORT 9000

# Bundle app source
COPY . .

# Specify port
EXPOSE 9000

# start app
CMD ["npm", "start"]
