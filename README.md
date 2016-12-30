# Simple api microservice example

[Demo](http://mardari.net:8080/api/products "Demo")

### Setup

- run `docker-compose up -d`
- execute inside app container: `/go/bin/golang-product-reviews-api -env=prod migrate-db && /go/bin/golang-product-reviews-api -env=prod load-fixtures`
- enjoy!
