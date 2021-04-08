# Container to compile the app
# docker run cr -v ~/.cryptocurrencyD:/root/.cryptocurrencyD -v ~/.cryptocurrencyCLI:/root/.cryptocurrencyCLI octa_node_3 rm -r ~/.cryptocurrencyD
FROM golang:1.15-alpine AS build-env

ENV PACKAGES curl make git libc-dev bash gcc linux-headers eudev-dev python3

WORKDIR /go/src/github.com/ivansukach/cryptocurrency

COPY . .

RUN apk add --no-cache $PACKAGES && make
ADD init.sh /
RUN chmod +x /init.sh
RUN /init.sh
#O=/app/cryptocurrency
#o - output
# Final container image
FROM alpine:latest


COPY --from=build-env /go/bin/octadaemon /usr/bin/octadaemon
COPY --from=build-env /root/.octa /root/.octa

EXPOSE 5432 26656 26657 26658 26660 6060 1317

CMD ["octadaemon", "start"]