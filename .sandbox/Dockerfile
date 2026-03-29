# syntax=docker/dockerfile:1
# check=error=true

# ============================================================
# Stage 1: base – Python + common tools + AI
# ============================================================
FROM python:3.14-slim AS base

ARG USER
ARG UID=1000
ARG GID=1000

RUN apt-get update \
  && apt-get install -y --no-install-recommends \
  bash-completion \
  ca-certificates \
  curl \
  git \
  gnupg \
  iproute2 \
  jq \
  less \
  lsof \
  make \
  nano \
  socat \
  tree \
  unzip \
  xclip \
  zip \
  # Process / debug tools
  procps \
  htop \
  gdb \
  strace \
  # Editors
  neovim \
  shellcheck \
  && rm -rf /var/lib/apt/lists/*

# --- Create non-root user ---
RUN groupadd --gid ${GID} ${USER} \
  && useradd \
  --uid ${UID} \
  --gid ${GID} \
  --shell /bin/bash \
  --create-home \
  ${USER}

RUN mkdir -p /workspace && chown ${UID}:${GID} /workspace

# --- Install sandbox CA certificate (must run as root) ---
COPY proxy/certs/ca.crt /usr/local/share/ca-certificates/sandbox-ca.crt
RUN update-ca-certificates

USER ${USER}

ENV PATH="/home/${USER}/.local/bin:$PATH" \
  PYTHONUNBUFFERED=1 \
  PYTHONDONTWRITEBYTECODE=1 \
  TERM=xterm-256color \
  LANG=C.UTF-8

RUN curl -LsSf https://astral.sh/uv/install.sh | sh && \
  echo "alias clip='xclip -selection clipboard'" >> $HOME/.bashrc

RUN curl -fsSL https://claude.ai/install.sh | bash

WORKDIR /workspace

CMD ["bash"]

# ============================================================
# Stage 2: go – Go toolchain on top of base
# ============================================================
FROM base AS go

USER root

RUN curl -OL https://golang.org/dl/go1.26.0.linux-amd64.tar.gz \
    && tar -C /usr/local -xzf go1.26.0.linux-amd64.tar.gz \
    && rm go1.26.0.linux-amd64.tar.gz

ENV GOPATH=/home/${USER}/go \
    GOTMPDIR=/home/${USER}/.go-tmp \
    PATH="/usr/local/go/bin:/home/${USER}/go/bin:${PATH}"

RUN mkdir -p /home/${USER}/go/pkg /home/${USER}/.go-tmp \
    && chown -R ${USER}:${USER} /home/${USER}/go /home/${USER}/.go-tmp

USER ${USER}

RUN go install go.uber.org/nilaway/cmd/nilaway@latest \
    && go install github.com/spf13/cobra-cli@latest
