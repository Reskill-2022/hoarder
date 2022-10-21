FROM golang:1.19-alpine

RUN apk update && apk add --no-cache git bash

ARG AWS_DEFAULT_REGION
ENV AWS_DEFAULT_REGION=$AWS_REGION

ARG AWS_SECRET_ACCESS_KEY
ENV AWS_SECRET_ACCESS_KEY=$AWS_SECRET_ACCESS_KEY

ARG AWS_ACCESS_KEY_ID
ENV AWS_ACCESS_KEY_ID=$AWS_ACCESS_KEY_ID

ARG GITHUB_TOKEN
ENV GITHUB_TOKEN=$GITHUB_TOKEN

ARG GITHUB_USER
ENV GITHUB_USER=$GITHUB_USER

ADD . /hoarder
WORKDIR /hoarder

ENV GOPRIVATE=github.com/Reskill-2022
RUN git config --global url."https://$GITHUB_USER:$GITHUB_TOKEN@github.com/".insteadOf "https://github.com/"
RUN go get ./...
ENV GO_ENV=production
RUN go install github.com/cosmtrek/air@latest

EXPOSE 8001

# Run the air command to start a live-reload server. Check .air.toml file for more details
ENTRYPOINT ["air"]

