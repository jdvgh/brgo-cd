FROM golang:1.21 as build

WORKDIR /go/src/app
COPY . . 
RUN go mod download

RUN CGO_ENABLED=0 go build  -o /go/bin/main  ./gitserver

FROM brgo-cd/bases/git:0.0.1 as run 
# RUN groupadd -r runner && useradd --no-log-init -r -g runner runner
# USER runner
COPY --from=build --chown=runner:runner /go/bin/main /
# COPY --from=build  /go/bin/main /
# COPY --from=git  /opt/bitnami/git/bin/git /usr/bin/git
ENV HTTP_PORT=8080
ENV TMP_DIR_PREFIX=/var/gitmnt
EXPOSE 8080
ENTRYPOINT ["/main"]
