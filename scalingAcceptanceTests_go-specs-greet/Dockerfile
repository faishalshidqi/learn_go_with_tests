# Make sure to specify the same Go version as the on in the go.mod file
FROM golang:1.23-alpine
WORKDIR /app
ARG bin_to_build
COPY go.mod ./
RUN go mod download
COPY . .
RUN go build -o svr cmd/${bin_to_build}/*.go
EXPOSE 50051
#LABEL authors="shidqi"

#ENTRYPOINT ["top", "-b"]
CMD ["./svr"]