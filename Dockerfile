FROM scratch
MAINTAINER Luke Thompson <luke@dukeluke.com>

LABEL Description="This image is of a Sample Web Service which provides the localized response."
LABEL Version=0.0.dev

COPY ./release/service /service
COPY ./data /data
COPY ./certs /certs

EXPOSE 4001
ENTRYPOINT ["/service"]
