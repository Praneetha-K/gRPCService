# Using image from official Golang Images
FROM golang:latest 

# Set the working directory
WORKDIR /go/src/app

# Copying the local files to the container
COPY . .

# Building the application
RUN go build -v -o gRPCApp .

# Exposing the port on which application runs 
EXPOSE 8080

# Running the application
CMD ["./gRPCApp"]
