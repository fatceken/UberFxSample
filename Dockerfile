FROM golang:1.17-alpine AS build
WORKDIR /src
COPY . .

WORKDIR /src/cmd

RUN go get -d -v

RUN go build -o /out/cmd

FROM alpine:3.13
COPY ./pkg/app/config.yaml /usr/app/
COPY ./pkg/app/config.*.yaml /usr/app/
COPY --from=build /out/cmd /usr/app/cmd

WORKDIR /usr/app


ENTRYPOINT ["./cmd"]