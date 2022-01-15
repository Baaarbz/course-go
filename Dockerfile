FROM golang:alpine AS build

RUN apk add --update git
WORKDIR /course-go/api
COPY . .
RUN CGO_ENABLED=0 go build -o /course-go/bin/course-go-api cmd/api/main.go

# Building image with the binary
FROM scratch
COPY --from=build /course-go/bin/course-go-api /course-go/bin/course-go-api
ENTRYPOINT ["/course-go/bin/course-go-api"]