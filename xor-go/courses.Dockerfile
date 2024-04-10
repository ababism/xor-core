FROM golang:1.21 as build
WORKDIR /app

COPY .. .

ENV CGO_ENABLED=0
ENV GOOS=linux

#RUN apt-get update && \
#    apt-get --yes --no-install-recommends install make="4.3-4.1" && \
#    apt-get clean && rm -rf /var/lib/apt/lists/*

RUN go build -o courses-svc ./services/courses/cmd

FROM alpine:latest as production

COPY --from=build /app/courses-svc ./

COPY --from=build /app/.env.docker ./.env
COPY --from=build /app/services/courses/migrations ./migrations
COPY --from=build /app/services/courses/config/config.docker.yml ./config/config.local.yml

CMD ["./courses-svc"]

EXPOSE 8080