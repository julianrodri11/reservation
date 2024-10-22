# Usa una imagen oficial de Golang como base
FROM golang:1.22.7-alpine

# Establece el directorio de trabajo dentro del contenedor
WORKDIR /app

# Copia todos los archivos del directorio actual al contenedor
COPY . .

# Descarga todos los módulos de Go
RUN go mod tidy

# Compila la aplicación de Go
RUN go build -o main .

# Expone el puerto 8080 al mundo exterior
EXPOSE 8080

# Comando para ejecutar la aplicación compilada
CMD ["./main"]