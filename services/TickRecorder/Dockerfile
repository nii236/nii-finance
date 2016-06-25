FROM openalgotplatform_go:0.1
WORKDIR /go/src/github.com/nii236/nii-finance/services/TickRecorder
COPY . /go/src/github.com/nii236/nii-finance/
RUN go install
ENTRYPOINT ["TickRecorder"]
