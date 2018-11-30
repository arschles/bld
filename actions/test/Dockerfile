FROM golang:1.11-alpine

LABEL "com.github.actions.name"="Go Test"
LABEL "com.github.actions.description"="Test Go Code"
LABEL "com.github.actions.icon"="mic"
LABEL "com.github.actions.color"="purple"

LABEL "repository"="http://github.com/arschles/bld"
LABEL "homepage"="http://github.com/actions"
LABEL "maintainer"="Octocat <octocat@github.com>"

RUN apk add --no-cache \
    bash \
    ca-certificates

ADD entrypoint.sh /entrypoint.sh
ENTRYPOINT ["/entrypoint.sh"]
