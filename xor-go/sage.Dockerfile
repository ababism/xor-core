FROM golang:1.21.4 as build
WORKDIR /app

#COPY ./.github ./.github
#COPY ./xor-python/scripts ./xor-python/scripts
#COPY ./proto ./proto
#COPY ./xor-go ./xor-go
COPY . .

ENV CGO_ENABLED=0
ENV GOOS=linux

RUN apt-get update && \
    apt-get --yes --no-install-recommends install make="4.3-4.1" && \
    apt-get install -y protobuf-compiler && \
    apt-get clean && rm -rf /var/lib/apt/lists/*

RUN ./.github/workflows/etc/prepare_protos.sh

WORKDIR /app/xor-go

RUN go build -o sage-svc ./services/sage/cmd

FROM alpine:latest as production

#WORKDIR /app/xor-go

COPY --from=build /app/xor-go/sage-svc ./

COPY --from=build ./app/.env ./services/sage/
COPY --from=build ./app/xor-go/services/sage/config/config.docker.yml ./services/sage/config/config.local.yml
COPY --from=build ./app/xor-go/services/sage/config/resources-config.yml ./services/sage/config/resources-config.yml

CMD ["./sage-svc"]

EXPOSE 8084
