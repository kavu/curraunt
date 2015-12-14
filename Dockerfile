FROM alpine:edge

MAINTAINER Max Riveiro <kavu13@gmail.com>

WORKDIR /GOPATH/src/github.com/kavu/curraunt
COPY . /GOPATH/src/github.com/kavu/curraunt

ENV GOPATH /GOPATH

RUN apk add --update go git bash && rm -rf /var/cache/apk/*

RUN mkdir -p /GOPATH && \
    go get -u \
      github.com/davecheney/profile \
      github.com/julienschmidt/httprouter \
      github.com/openprovider/ecbrates \
      github.com/pquerna/ffjson/ffjson

RUN ./scripts/build.sh && cp curraunt /bin

EXPOSE 80

ENTRYPOINT [ "/bin/curraunt" ]
CMD ["-p", "80"]
