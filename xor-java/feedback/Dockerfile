FROM python:3.12-slim-bullseye AS prepare-protos

WORKDIR /xor-core

COPY . .

COPY --from=golang:1.21.4-alpine /usr/local/go/ /usr/local/go/
ENV PATH="/usr/local/go/bin:${PATH}"

RUN apt update && \
    apt install -y protobuf-compiler && \
    apt clean

RUN ./.github/workflows/etc/prepare_protos.sh

FROM gradle:latest AS build-java

WORKDIR /xor-java

COPY --from=prepare-protos /xor-core/xor-java .

RUN gradle feedback:build

FROM amazoncorretto:17

WORKDIR /feedback

COPY --from=build-java /xor-java/feedback/build/libs/*.jar /feedback/app.jar

ENTRYPOINT ["java", "-jar", "app.jar"]
