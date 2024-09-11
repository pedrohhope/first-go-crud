# Use official Golang image

FROM golang:1.22.2 as builder

# Set the working directory inside the container
WORKDIR /app

# Copy the source code
COPY . .

# Download dependencies
RUN go get -d -v ./...

# Build the application
RUN go build -o api .

#EXPOSE 8080
EXPOSE 8080

# Run the application
CMD ["./api"]