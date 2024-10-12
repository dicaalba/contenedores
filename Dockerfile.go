# Imagen base para Go
FROM golang:1.19-alpine

# Establecemos el directorio de trabajo
WORKDIR /go

# Copiamos los archivos necesarios
COPY go/main.go main.go

# Instalamos las dependencias
RUN go build -o main .

# Exponer el puerto 8081 
EXPOSE 8081

# Ejecutar el comando de inicio de Go
CMD ["./main"]