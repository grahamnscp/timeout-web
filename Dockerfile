FROM golang:alpine AS compile
COPY timeout-web.go /go
RUN go build timeout-web.go

FROM registry.suse.com/bci/bci-base:15.7
COPY --from=compile /go/timeout-web /
EXPOSE 8080
ENTRYPOINT ["/timeout-web"]
