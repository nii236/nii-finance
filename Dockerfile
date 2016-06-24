FROM golang:1.6-alpine
RUN apk update && apk add git
RUN go get github.com/Masterminds/glide
COPY . /go/src/github.com/nii236/nii-finance
WORKDIR /go/src/github.com/nii236/nii-finance
RUN glide up
