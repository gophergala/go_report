 FROM golang:latest

 RUN mkdir -p /go/src/github.com/gophergala/go_report
 WORKDIR /go/src/github.com/gophergala/go_report


 COPY . /go/src/github.com/gophergala/go_report
 RUN go-wrapper install
 CMD ["go-wrapper", "run"]
 ONBUILD COPY . /go/src/github.com/gophergala/go_report
 #ONBUILD RUN echo '/go/src/github.com/gophergala/go_report' > .godir
 #ONBUILD RUN go install github.com/gophergala/go_report
 ONBUILD RUN go-wrapper install
 #ENTRYPOINT /go/bin/go_report
 EXPOSE 8080
