FROM alpine:latest
RUN apk update
RUN apk add --no-cache go git
ENV GOROOT /usr/lib/go
ENV GOPATH /go
ENV PATH /go/bin:$PATH
RUN go get -u github.com/revel/cmd/revel
EXPOSE 9000/tcp
RUN cd $GOPATH && revel new bojo
COPY bojo /go/src/bojo
CMD cd $GOPATH && revel run -a bojo
