FROM golang:1.21 as build

WORKDIR /go/src/app
COPY . . 
RUN go mod download

RUN CGO_ENABLED=0 go build -o /go/bin/main

FROM brgo-cd/bases/kubectl:0.0.1 as run 
# RUN groupadd -r runner && useradd --no-log-init -r -g runner runner
# USER runner
# COPY --from=build --chown=runner:runner /go/bin/main /
COPY --from=build /go/bin/main /
# COPY --from=kubectl  --chown=runner:runner  /opt/bitnami/kubectl/bin/kubectl /usr/bin/kubectl
ENV KUBECONFIG=/mnt/kubeconfig.yaml
ENV REPO_URL=https://github.com/jdvgh/sample-manifests.git

CMD ["/main"]
