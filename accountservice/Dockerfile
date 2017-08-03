FROM golang

# Install beego and the bee dev tool
RUN go get github.com/gorilla/mux && go get github.com/boltdb/bolt
RUN go get github.com/stretchr/testify/mock

ADD . /go/src/github.com/dmr/project/accountservice
RUN go install github.com/dmr/project/accountservice
ENTRYPOINT /go/bin/accountservice
# Expose the application on port 6767
EXPOSE 6767

# Set the entry point of the container to the bee command that runs the
# application and watches for changes
CMD ["accountservice", "run"]