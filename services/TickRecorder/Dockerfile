FROM golang:1.6-alpine
RUN apk update && apk add git
RUN go get github.com/Masterminds/glide
COPY . /go/src/open-algot.servebeer.com/open-algot/open-algot-platform/
WORKDIR /go/src/open-algot.servebeer.com/open-algot/open-algot-platform
RUN glide up
WORKDIR /go/src/open-algot.servebeer.com/open-algot/open-algot-platform/services/TickRecorder
RUN go install
ENTRYPOINT ["TickRecorder"]
