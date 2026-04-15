# Usa una imagen reciente Go
FROM golang:1.24-alpine
WORKDIR /usr/local/appdocker

# Instala Git si es necesario para dependencias
RUN apk add --no-cache git

# Crea el directorio de trabajo
WORKDIR /app

# Copia los archivos del proyecto
COPY . .

RUN apk add --no-cache make # Instala make
RUN go mod download

# Compila la aplicación
RUN make build

# Expone el puerto que usa Fiber
EXPOSE 3000

# Comando de inicio
CMD ["./rest-api"]
