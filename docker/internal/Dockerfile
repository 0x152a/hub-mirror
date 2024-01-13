FROM node:18

COPY ./app /app

WORKDIR /app
RUN apt update && apt install libsdl-pango-dev -y
RUN npm install -g pnpm

RUN pnpm install

ENV LISTEN_ADDRESS=0.0.0.0
ENV PORT=80
RUN pnpm build
CMD pnpm start
