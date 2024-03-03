FROM node:20

COPY ./app /app
WORKDIR /app
RUN cd /app && npm i && npm run build

CMD ["npm", "start"]
