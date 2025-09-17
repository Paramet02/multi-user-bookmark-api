# 1 use golang image as base image

FROM golang:1.23-alpine

# 2 set working directory
WORKDIR /app

# 3 copy go mod and sum files
COPY go.mod go.sum ./

# 4 download all dependencies
RUN go mod download

# 5 copy the source code
COPY . .

# 6 build the application
RUN go build -o main .

# 7 expose port 8080 to the outside world
EXPOSE 8080

# 8 command to run the executable
CMD ["./main"]
