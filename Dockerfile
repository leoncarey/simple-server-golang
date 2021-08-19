FROM golang:1.16

# Set destination for COPY
WORKDIR /app

# Download Go modules
COPY go.mod .
COPY go.sum .
RUN go mod download

COPY ./src/*.go ./app

# Build
RUN go build -o /rest-api-go

EXPOSE 8080

# Run
CMD [ "/rest-api-go" ]