FROM golang:alpine AS build

RUN apk add --update git
WORKDIR /course-go/api
COPY . .
RUN CGO_ENABLED=0 go build -o /go/bin/course-go-api cmd/api/main.go

# Building image with the binary
FROM scratch
COPY --from=build /go/bin/course-go-api /go/bin/course-go-api
ENTRYPOINT ["/go/bin/course-go-api"]