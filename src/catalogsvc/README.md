# Catalog service

## Settings
Set env variables(defaults in examples):
```env
DEBUG=true
DEBUG_PPROF_PORT=8881

PORT=8081
READ_TIMEOUT=30s
WRITE_TIMEOUT=30s
MAX_HEADER_MB=20

MONGO_TIMEOUT=30s
MONGO_URI=mongodb://localhost:27017
MONGO_USER=
MONGO_PASSWORD=
MONGO_DB_NAME=catalog
```
## Available endpoints


* GET /catalog/v1/products/ - 200/500
  * rpc GetAllProducts(Empty) returns (Products) {}

* GET /catalog/v1/products/:id -- 200/400/500 
  * rpc GetProductByID(GetProductByIDRequest) returns (Product) {}
