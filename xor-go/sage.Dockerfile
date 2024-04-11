FROM golang:1.21.4 as build
WORKDIR /app

COPY . .

ENV CGO_ENABLED=0
ENV GOOS=linux

RUN apt-get update && \
    apt-get --yes --no-install-recommends install make="4.3-4.1" && \
    apt-get install -y protobuf-compiler && \
    apt-get clean && rm -rf /var/lib/apt/lists/*

RUN ../.github/workflows/etc/prepare_protos.sh

RUN go build -o sage-svc ./services/sage/cmd

FROM alpine:latest as production

COPY --from=build /app/sage-svc ./

COPY --from=build /app/.env ./services/sage/
COPY --from=build /app/services/sage/migrations ./services/sage/migrations
COPY --from=build /app/services/sage/config/config.docker.yml ./services/sage/config/config.local.yml

CMD ["./sage-svc"]

EXPOSE 8084