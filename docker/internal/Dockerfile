FROM amazoncorretto:11-alpine-jdk

COPY ./app /app
WORKDIR /app
RUN cd /app && ./gradlew :web:build

EXPOSE 8888
CMD ["./gradlew", ":web:run", "-Pgbind=0.0.0.0"]
