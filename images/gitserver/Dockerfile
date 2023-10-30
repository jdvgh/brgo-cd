FROM golang:1.21 as build

WORKDIR /go/src/app
COPY . . 
RUN go mod download

RUN CGO_ENABLED=0 go build  -o /go/bin/main  ./gitserver/service

FROM brgo-cd/bases/git:0.0.1 as run 
COPY --from=build /go/bin/main /
ENV GRPC_PORT=50051
ENV TMP_DIR_PREFIX=/var/gitmnt
EXPOSE 50051
ENTRYPOINT ["/main"]
