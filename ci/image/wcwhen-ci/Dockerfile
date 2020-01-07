FROM golang:1-stretch

ENV DEBIAN_FRONTEND="noninteractive"

RUN apt-get update && apt-get -y --no-install-recommends install apt-transport-https ca-certificates

# setup gcloud apt repo
RUN echo "deb [signed-by=/usr/share/keyrings/cloud.google.gpg] https://packages.cloud.google.com/apt cloud-sdk main" | tee -a /etc/apt/sources.list.d/google-cloud-sdk.list \
    && curl https://packages.cloud.google.com/apt/doc/apt-key.gpg | apt-key --keyring /usr/share/keyrings/cloud.google.gpg add -

# setup kubernetes apt repo
RUN echo "deb https://apt.kubernetes.io/ kubernetes-stretch main" | tee -a /etc/apt/sources.list.d/kubernetes.list \
    && curl -s https://packages.cloud.google.com/apt/doc/apt-key.gpg | apt-key add -

RUN apt-get update \
    && apt-get -y --no-install-recommends install \
    google-cloud-sdk \
    kubectl \
    git \
    python \
    wget \
    jq \
    vim \
    && rm -rf /var/lib/apt/lists/*

## Setup config to make git commits
RUN git config --global user.email "wuyichen1987@gmail.com"
RUN git config --global user.name "wu-yichen"