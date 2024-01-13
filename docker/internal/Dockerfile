FROM node:20

COPY ./app /app

WORKDIR /app

RUN yarn

RUN yarn build
CMD yarn start
