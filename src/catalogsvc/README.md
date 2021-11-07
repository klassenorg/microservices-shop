# Catalog service

##Settings
Set env variables(defaults in examples):
```env
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
##Available endpoints

GET /catalog/v1/products/ - 200/500

GET /catalog/v1/products/:id -- 200/400/500