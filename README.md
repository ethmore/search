Search Service
===

This GO project serves as a microservice for [eCommerce](https://github.com/users/ethmore/projects/4) project.


## Service task:

- Add product to elasticsearch db
- Request products from elasticsearch-service by search query 



# Installation

Ensure GO is installed on your system
```
go mod download
````

```
go run .
```

## Test
```
curl http://localhost:3006/test
```
### It should return:
```
StatusCode        : 200
StatusDescription : OK
Content           : {"message":"OK"}
```

## Example .env file
This file should be placed inside `dotEnv` folder
```
# Cors URLs
BFF_URL = http://localhost:3001

# Request URLs
GET_ALL_PRODUCTS = http://localhost:3002/getAllProducts
ADD_PRODUCT = http://localhost:9200/products/_doc
SEARCH_PRODUCT = http://localhost:9200/products/_search
```