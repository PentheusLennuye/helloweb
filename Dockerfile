ARG ALPINE_VERSION=3.17

FROM alpine:${ALPINE_VERSION}
ARG USERNAME=apprunner
ARG USER_UID=1000
ARG USER_GID=1000
ARG HOMEDIR=/app
ARG SHELL=/bin/sh

RUN adduser $USERNAME -s $SHELL -D -h $HOMEDIR -u $USER_UID $USER_GID 

RUN apk add -q --update --progress --no-cache

RUN for tool in tools/cmd/goimports \
                lint/golint \
                tools/go/analysis/passes/shadow/cmd/shadow \
                cweill/gotests/gotest; \
    do go install golang.org/x/${tool}@latest; \
    done && go install honnef.co/go/tools/cmd/staticcheck@latest

USER $USERNAME
WORKDIR /home/$USERNAME
COPY bin/* .

# To run mkdocs server as per Eitri standards
COPY config/* .
RUN python3 -m pip install -r requirements.txt

USER root
