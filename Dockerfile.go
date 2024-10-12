# Imagen base para Go
FROM golang:1.19-alpine

# Establecemos el directorio de trabajo
WORKDIR /go/src/app

# Copiamos los archivos necesarios
COPY go/go.mod go.mod
COPY go/main.go main.go

# Instalamos las dependencias
RUN go mod download

# Compilamos la aplicación
RUN go build -o main .

# Exponer el puerto 8081 
EXPOSE 8081

# Ejecutar el comando de inicio de Go
CMD ["./main"]
