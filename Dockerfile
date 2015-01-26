 FROM golang:latest

 RUN mkdir -p /go/src/github.com/gophergala/go_report
 WORKDIR /go/src/github.com/gophergala/go_report
 ENV PATH /go/bin:$PATH
 COPY . /go/src/github.com/gophergala/go_report
 RUN apt-get -t lenny-backports install bzr
 RUN go get golang.org/x/tools/cmd/goimports
 RUN go get github.com/fzipp/gocyclo
 RUN go get github.com/golang/lint/golint
 RUN go get golang.org/x/tools/cmd/vet
 RUN go get gopkg.in/mgo.v2
 RUN go-wrapper install
 CMD ["go-wrapper", "run"]
 #ONBUILD COPY . /go/src/github.com/gophergala/go_report
 #ONBUILD RUN echo '/go/src/github.com/gophergala/go_report' > .godir
 #ONBUILD RUN go install github.com/gophergala/go_report
 #ONBUILD RUN go-wrapper install golint
 #ONBUILD RUN go-wrapper install
 #ENTRYPOINT /go/bin/go_report
 EXPOSE 8080
