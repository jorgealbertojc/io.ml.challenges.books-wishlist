FROM jorgealbertojc/alpine-glibc:alpine-3.16_glibc-2.35

LABEL maintainer=jalbertojc@gmail.com

WORKDIR /app

RUN mkdir -p /app/src/io.ml.challenges

ENV TZ America/Mexico_City
ENV GOROOT /usr/lib/go
ENV GOPATH /app
ENV GOBIN /app/bin

RUN apk --no-cache update \
    && apk --no-cache upgrade \
    && apk --no-cache --upgrade add \
    --virtual .build-deps \
    build-base gcc musl-dev bash make \
    curl tzdata jq go sqlite

ENV SHELL=/bin/bash

RUN ln -snf /usr/share/zoneinfo/America/Mexico_City /etc/localtime && echo 'America/Mexico_City' > /etc/timezone

COPY . /app/src/io.ml.challenges/io.ml.challenges.books-wishlist

RUN cd /app/src/io.ml.challenges/io.ml.challenges.books-wishlist \
    && go build -o /usr/local/bin/mlbwlistd ./cmd/mlbwlist/main.go \
    && mkdir -p /var/local/mlbwlistd \
    && sqlite3 /var/local/mlbwlistd/database.db < /app/src/io.ml.challenges/io.ml.challenges.books-wishlist/database/migrations/2022-06-29.books-wishlist.sql \
    && rm -rfv /app/src/io.ml.challenges

ENTRYPOINT [ "/usr/local/bin/mlbwlistd", "--configs", "/etc/mlbwlistd/configs/config.yml" ]
