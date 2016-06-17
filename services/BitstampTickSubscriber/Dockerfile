FROM openalgotplatform_go:0.1
WORKDIR /go/src/open-algot.servebeer.com/open-algot/open-algot-platform/services/BitstampTickSubscriber
COPY . /go/src/open-algot.servebeer.com/open-algot/open-algot-platform/
RUN go install
ENTRYPOINT ["BitstampTickSubscriber"]
