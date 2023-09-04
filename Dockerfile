FROM alpine/k8s:1.28.1 as k8s
FROM alpine/git:2.40.1 as git
COPY --from=k8s /usr/bin/kubectl /usr/bin/kubectl
# FROM scratch
# COPY --from=git /usr/bin/git /git
COPY main /main
ENV KUBECONFIG=/mnt/kubeconfig.yaml
ENTRYPOINT ["/main"]
