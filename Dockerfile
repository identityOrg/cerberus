FROM golang:alpine as build
RUN apk add --no-cache git make build-base

RUN go get github.com/GeertJohan/go.rice/rice

WORKDIR /go/src/github.com/identityOrg/cerberus

ADD . .

RUN go mod download

RUN rice clean
RUN CGO_ENABLED=1 GOOS=linux go build -o cerberus .
RUN rice -i github.com/identityOrg/cerberus/setup append --exec cerberus

FROM alpine:latest
RUN apk --no-cache add ca-certificates
RUN apk --no-cache add sqlite
WORKDIR /root/
COPY --from=build /go/src/github.com/identityOrg/cerberus/cerberus .

EXPOSE 8080

#RUN ./cerberus initconfig
#RUN ./cerberus migrate -f --demo

CMD ["./cerberus", "serve", "--demo", "--addr", ":8080"]