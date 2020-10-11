# Container to compile the app
# docker run cr -v ~/.cryptocurrencyD:/root/.cryptocurrencyD -v ~/.cryptocurrencyCLI:/root/.cryptocurrencyCLI octa_node_3 rm -r ~/.cryptocurrencyD
FROM golang:1.13-alpine AS build-env

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

COPY --from=build-env /go/bin/cryptocurrencyD /usr/bin/cryptocurrencyD
COPY --from=build-env /go/bin/cryptocurrencyCLI /usr/bin/cryptocurrencyCLI
COPY --from=build-env /root/.cryptocurrencyCLI /root/.cryptocurrencyCLI
COPY --from=build-env /root/.cryptocurrencyD /root/.cryptocurrencyD


EXPOSE 5432 26656 26657 26658 26660 6060 1317

CMD ["cryptocurrencyD", "start"]