 FROM golang:latest

 RUN mkdir -p /go/src/github.com/gophergala/go_report
 WORKDIR mkdir -p /go/src/github.com/gophergala/go_report

 CMD ["go-wrapper", "run"]

 ONBUILD COPY . /go/src/github.com/gophergala/go_report
 ONBUILD RUN echo '/go/src/github.com/gophergala/go_report' > .godir
# RUN go install github.com/gophergala/go_report
 ONBUILD RUN go-wrapper install
 ENTRYPOINT /go/bin/go_report
 EXPOSE 8080
