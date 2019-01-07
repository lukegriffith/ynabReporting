FROM node:alpine

ENV ynab_api_token replace
ENV ynab_budget_id replace
ENV cache_ttl replace
ENV cache_days replace

RUN mkdir /usr/src/
RUN mkdir /usr/src/app

COPY . /usr/src/app

WORKDIR /usr/src/app

RUN npm install 



CMD ["node", "server.js"] 

EXPOSE 8080