swagger validate data/api/bcard-microservice-swagger.yaml

swagger generate server --spec=data/api/bcard-microservice-swagger.yaml --with-expand --with-flatten=full --target=internal/infrastructure/http/ --server-package=./httpserver --api-package=./swaggerapi --model-package=./httpserver/swaggerapi/models --exclude-main