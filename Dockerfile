 FROM golang:onbuild
 ADD . /go/src/github.com/gophergala/go_report
 RUN go install github.com/gophergala/go_report
 ENTRYPOINT /go/bin/go_report
 EXPOSE 8080
