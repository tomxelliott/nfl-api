FROM scratch

ADD . /

EXPOSE 80

CMD ["/main"]

# To build:
# $ CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -o main .
# $ docker run -p 1234:5000 nfl-api-tom

# Navigate to localhost:1234
# Individual teams found at localhost:1234/1 etc.