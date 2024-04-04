FROM golang:1.21.4 as build
WORKDIR /app

COPY . .

ENV CGO_ENABLED=0
ENV GOOS=linux

RUN go build -o finances-svc ./services/finances/cmd

FROM alpine:latest as production

COPY --from=build /app/finances-svc ./

COPY --from=build /app/.env ./services/finances/
COPY --from=build /app/services/finances/migrations ./services/finances/migrations
COPY --from=build /app/services/finances/config/config.docker.yml ./services/finances/config/config.local.yml

CMD ["./finances-svc"]

EXPOSE 8084