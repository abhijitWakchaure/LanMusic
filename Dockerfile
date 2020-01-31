FROM abhijitwakchaure/lanmusic-go-base as preserver
WORKDIR /go/src/github.com/abhijitWakchaure/lanmusic/gosrc
COPY ./gosrc/ .
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o app . \
    && find . ! -name 'app' ! -name '.' ! -name '..' -exec rm -rf {} +

FROM nginx:alpine
LABEL maintainer="abhijitwakchaure.2014@gmail.com"
RUN apk --no-cache add ca-certificates
WORKDIR /usr/share/nginx/html
COPY --from=preserver /go/src/github.com/abhijitWakchaure/lanmusic/gosrc/app /usr/local/bin/lanmusic
ADD ./UI/dist/lanMusic/ .
ADD ./docker-entrypoint.sh /usr/local/bin/
ENTRYPOINT [ "/usr/local/bin/docker-entrypoint.sh"]
EXPOSE 80 9000