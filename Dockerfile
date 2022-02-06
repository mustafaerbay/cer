FROM alpine:latest AS default

ARG PRODUCT_VERSION=UNSPECIFIED
ARG BIN_NAME=cer
ARG TARGETARCH

LABEL maintainer="Mustafa Erbay <mustafaerbay365@gmail.com>"
LABEL version=$VERSION
LABEL "cer.version"="${VERSION}"

RUN apk add --no-cache git openssh
COPY ["dist/cer_linux_386/cer", "/bin/cer"]

ENTRYPOINT ["/bin/cer"]