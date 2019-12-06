FROM golang:1.13-alpine3.10

RUN apk update && apk add --no-cache make\
        gcc\
        g++\
        git\
        libwebp\
        binutils-gold \
        curl \
        gnupg \
        libgcc \
        linux-headers \
        make \
        python\
        libwebp-tools\
        libwebp-dev\
        tiff-dev\
        vips-dev\
        libzip libzip-dev\