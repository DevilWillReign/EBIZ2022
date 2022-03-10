FROM debian:bullseye

RUN apt update && apt install -y wget curl && apt autoremove && apt clean
ENV GOHASH 980e65a863377e69fd9b67df9d8395fd8e93858e7a24c9f55803421e453f4f99
RUN wget https://go.dev/dl/go1.17.8.linux-amd64.tar.gz && \
echo "$GOHASH go1.17.8.linux-amd64.tar.gz" | sha256sum -c - && \
tar -xf go1.17.8.linux-amd64.tar.gz -C /usr/local && \
rm go1.17.8.linux-amd64.tar.gz
ENV GOROOT /usr/local/go
ENV PATH $PATH:$GOROOT/bin
RUN go version
RUN curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(go env GOROOT)/bin v1.44.2
RUN golangci-lint --version
RUN useradd -m -s /bin/bash appuser
USER appuser
WORKDIR /home/appuser
