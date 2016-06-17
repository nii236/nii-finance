FROM openalgotplatform_go:0.1
WORKDIR /go/src/open-algot.servebeer.com/open-algot/open-algot-platform/services/OandaTickSubscriber
COPY . /go/src/open-algot.servebeer.com/open-algot/open-algot-platform/
RUN go install
ENTRYPOINT ["OandaTickSubscriber"]
