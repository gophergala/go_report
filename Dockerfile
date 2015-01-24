 FROM golang:latest

 RUN mkdir -p /go/src/github.com/gophergala/go_report
 WORKDIR /go/src/github.com/gophergala/go_report
 ENV PATH /go/bin:$PATH
 COPY . /go/src/github.com/gophergala/go_report
 RUN go get github.com/golang/lint/golint
 RUN go get golang.org/x/tools/cmd/vet
 RUN go-wrapper install
 CMD ["go-wrapper", "run"]
 #ONBUILD COPY . /go/src/github.com/gophergala/go_report
 #ONBUILD RUN echo '/go/src/github.com/gophergala/go_report' > .godir
 #ONBUILD RUN go install github.com/gophergala/go_report
 #ONBUILD RUN go-wrapper install golint
 #ONBUILD RUN go-wrapper install
 #ENTRYPOINT /go/bin/go_report
 EXPOSE 8080
