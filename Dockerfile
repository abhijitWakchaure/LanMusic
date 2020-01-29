FROM abhijitwakchaure/drl-go-base as preserver
WORKDIR /go/src/lanmusic/
COPY src/ .
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o lanmusic . \
    && find . ! -name 'lanmusic' ! -name '.' ! -name '..' -exec rm -rf {} +

FROM alpine
LABEL maintainer="abhijitwakchaure.2014@gmail.com"
WORKDIR /usr/local/bin/
COPY --from=preserver /go/src/lanmusic/lanmusic .
ENTRYPOINT [ "/usr/local/bin/lanmusic"]