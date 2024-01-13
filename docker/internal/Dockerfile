FROM node:20

COPY ./app /app

WORKDIR /app
RUN npm install -g yarn

RUN yarn

RUN yarn build
CMD yarn start
