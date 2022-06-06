FROM debian:bullseye

RUN apt update && apt install -y wget curl && apt autoremove && apt clean
ENV GOHASH e54bec97a1a5d230fc2f9ad0880fcbabb5888f30ed9666eca4a91c5a32e86cbc
RUN wget https://go.dev/dl/go1.18.2.linux-amd64.tar.gz && \
echo "$GOHASH go1.18.2.linux-amd64.tar.gz" | sha256sum -c - && \
tar -xf go1.18.2.linux-amd64.tar.gz -C /usr/local && \
rm go1.18.2.linux-amd64.tar.gz
ENV GOROOT /usr/local/go
ENV PATH $PATH:$GOROOT/bin
RUN go version
RUN curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(go env GOROOT)/bin v1.44.2
RUN golangci-lint --version
RUN useradd -m -s /bin/bash appuser
USER appuser
WORKDIR /home/appuser
