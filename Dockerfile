FROM golang:1.12-alpine3.9
RUN mkdir /ucf-garage-api
ADD . /ucf-garage-api
WORKDIR /ucf-garage-api
RUN apk add git
RUN go get github.com/sasho2k/University-Of-Central-Florida-Garage-API
ENV PORT 8080
RUN  go build -o main .
CMD ["/ucf-garage-api/main"]