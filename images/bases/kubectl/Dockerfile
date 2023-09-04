FROM brgo-cd/bases/ubuntu:22.04
RUN curl -LO "https://dl.k8s.io/release/v1.28.1/bin/linux/amd64/kubectl"
RUN curl -LO "https://dl.k8s.io/v1.28.1/bin/linux/amd64/kubectl.sha256"
RUN echo "$(cat kubectl.sha256)  kubectl" | sha256sum --check
RUN install -o root -g root -m 0755 kubectl /usr/local/bin/kubectl
RUN chmod +x kubectl
RUN mkdir -p ~/.local/bin
RUN cp ./kubectl ~/.local/bin/kubectl
ENV PATH=$PATH:~/.local/bin:/root/.local/bin
