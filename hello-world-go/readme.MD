# Ejemplo invocacion de servicio con http y grpc en go local

## Iniciar dapr cli

```bash
dapr init
```
## Correr ejemplo http

Primero correr el server-http con el siguiente comando dentro de la carpeta server-http

```bash
dapr run --app-port 6006 --app-id server-http --app-protocol http --dapr-http-port 3501 -- go run .
```

Este comando corre la aplicacion con un side car de dapr, este comando abre el puerto 6006 para la aplicacion y le da el app-id "server-http" para que dapr lo encuentre, luego se abre el puerto 3501 para el uso de dapr

Luego correr el cliente http dentro de la carpeta client-http

```bash
dapr run --app-id client-http --app-protocol http --dapr-http-port 3500 -- go run .
```

Este comando corre la aplicacion con un side car de dapr, este comando permite el uso de protocolo http para la aplicacion, y abre el puerto 3500 para la comunicacion con el sidecar de dapr

## Correr ejemplo gRPC

Primero correr el server-gRPC con el siguiente comando dentro de la carpeta server-gRPC

```bash
dapr run --app-id server --app-port 50051 -- go run main.go
```

Este comando corre la aplicacion con un side car de dapr, este comando abre el puerto 50051 para la aplicacion y le da el app-id "server".

Luego correr el cliente http con el siguiente comando dentro de la carpeta client-gRPC

```bash
dapr run --app-id client --dapr-grpc-port 50007 -- go run main.go
```

Este comando corre la aplicacion con un side car de dapr, este comando permite el uso de gRPC para la aplicacion, y abre el puerto 50007 para la comunicacion con el sidecar de dapr