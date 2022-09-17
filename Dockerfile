#This is multistage Docker build 
FROM golang:1.19.1-alpine3.15 as builder
WORKDIR /build
COPY go.mod go.sum ./
# Download all dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed
RUN go mod download
COPY . .
RUN go build -o main cmd/web/*
 


from golang:1.19.1-alpine3.15 
# Add Maintainer Info
LABEL maintainer="Danail Surudzhiyski"


# Set the Current Working Directory inside the container
WORKDIR /app

# Copy only the executable binary from previous build
COPY --from=builder /build/main .
COPY --from=builder /build/templates .
COPY --from=builder /build/static .


# Expose port 8080 to the outside world
EXPOSE 8080

# Command to run the executable
CMD ["./main"]
