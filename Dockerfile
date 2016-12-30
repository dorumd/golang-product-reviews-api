# golang image where workspace (GOPATH) configured at /go.
FROM golang:latest

# Copy the local package files to the container’s workspace.
ADD . /go/src/github.com/dorumd/golang-product-reviews-api

# Build the golang-docker command inside the container.
RUN go install github.com/dorumd/golang-product-reviews-api

# Migrate the db
RUN /go/bin/golang-product-reviews-api -env=prod migrate-db

# Load fixtures
RUN /go/bin/golang-product-reviews-api -env=prod load-fixtures

# Run the golang-docker command when the container starts.
ENTRYPOINT ["/go/bin/golang-product-reviews-api", "-env=prod", "serve"]

# http server listens on port 8080.
EXPOSE 8080
