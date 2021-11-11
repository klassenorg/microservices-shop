# Cart service

## Settings
Set env variables(defaults in examples):
```env
PORT=8082
READ_TIMEOUT=30s
WRITE_TIMEOUT=30s
MAX_HEADER_MB=20

REDIS_TIMEOUT=30s
REDIS_ADDR=localhost:6379
REDIS_PASSWORD=
```
## Available endpoints


* GET /cart/v1 - 200/500

* POST /cart/v1/add -- 200/400/500

* POST /cart/v1/remove -- 200/400/500

* POST /cart/v1/remove/all -- 200/400/500