FROM golang:1.14 as build-image

WORKDIR /go/src
COPY . .

RUN go build -o ../bin main.go

FROM public.ecr.aws/lambda/go:1

COPY --from=build-image /go/bin/ /var/task/

RUN ls /var/task/

# Command can be overwritten by providing a different command in the template directly.
CMD ["main"]