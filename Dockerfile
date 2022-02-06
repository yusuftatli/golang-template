FROM golang@sha256:6c35b28ceee082c621c818c097f418cd55104a16e51003120c1bf37111ed9cfe as builder

ENV USER=gowit
ENV UID=10001



WORKDIR /app

COPY go.mod .
COPY go.sum .

RUN go mod download
RUN go mod verify

COPY . .