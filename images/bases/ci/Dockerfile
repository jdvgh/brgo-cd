FROM brgo-cd/bases/ubuntu:22.04
# git
RUN apt update -y
RUN apt install -y git
# kubectl
RUN curl -LO "https://dl.k8s.io/release/v1.28.1/bin/linux/amd64/kubectl"
RUN curl -LO "https://dl.k8s.io/v1.28.1/bin/linux/amd64/kubectl.sha256"
RUN echo "$(cat kubectl.sha256)  kubectl" | sha256sum --check
RUN install -o root -g root -m 0755 kubectl /usr/local/bin/kubectl
RUN chmod +x kubectl
RUN mkdir -p ~/.local/bin
RUN cp ./kubectl ~/.local/bin/kubectl
ENV PATH=$PATH:~/.local/bin:/root/.local/bin
# Taskfile
RUN sh -c "$(curl --location https://taskfile.dev/install.sh)" -- -d 
# Go
RUN curl -LO "https://go.dev/dl/go1.21.3.linux-amd64.tar.gz"
RUN echo "1241381b2843fae5a9707eec1f8fb2ef94d827990582c7c7c32f5bdfbfd420c8  go1.21.3.linux-amd64.tar.gz" | sha256sum --check
RUN tar -C /usr/local -xzf go1.21.3.linux-amd64.tar.gz 
ENV PATH=$PATH:/usr/local/go/bin
