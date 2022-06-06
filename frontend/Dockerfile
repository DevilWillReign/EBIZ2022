FROM node:16

WORKDIR /usr/src/app

# install and cache app dependencies
COPY package*.json ./
ADD package.json /usr/src/app/package.json
RUN npm install

ENV PORT 9000

# Bundle app source
COPY . .

# Specify port
EXPOSE 80
EXPOSE 9000

# start app
CMD ["npm", "start"]
