FROM node:alpine

RUN mkdir /usr/src/
RUN mkdir /usr/src/app

COPY . /usr/src/app

WORKDIR /usr/src/app

RUN npm install 



CMD ["node", "server.js"] 

EXPOSE 80808