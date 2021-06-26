FROM public.ecr.aws/lambda/provided:al2 as builder

# install compiler
RUN yum install -y golang
RUN go env -w GOPROXY=direct

# cache dependencies
ADD src/go.mod src/go.sum ./
RUN go mod download

# build
WORKDIR /go/src/app
COPY /src /go/src/app
RUN go build -o /app
FROM public.ecr.aws/lambda/provided:al2
COPY --from=builder /app /app
ENTRYPOINT [ "/app" ]   