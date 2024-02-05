FROM openjdk:11-jdk-alpine

COPY ./app /app
WORKDIR /app
RUN cd /app && ./gradlew :web:build

EXPOSE 8888
CMD ["./gradlew", ":web:run", "-Pgbind=0.0.0.0"]
