# Imagen base con la versión específica de Go
FROM golang:1.21.10

# Directorio de trabajo
WORKDIR /app

# Copiar archivos del proyecto
COPY go.mod go.sum ./
RUN go mod download

COPY . .

# Compilar la aplicación
RUN go build -o main .

# Exponer el puerto de la aplicación
EXPOSE 8080

# Ejecutar la aplicación
CMD ["./main"]
