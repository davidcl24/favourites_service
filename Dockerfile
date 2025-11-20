# ============================
# 1) Etapa de build
# ============================
FROM golang:1.24.3-alpine AS build

# Crear carpeta de trabajo
WORKDIR /app

# Configurar build estático obligatorio para scratch
ENV CGO_ENABLED=0 GOOS=linux GOARCH=amd64

# Copiar archivos de dependencias
COPY go.mod go.sum ./

# Descargar dependencias
RUN go mod download

# Copiar el resto del proyecto
COPY . .

# Compilar binario totalmente estático
RUN go build -ldflags="-s -w" -o server ./app

# ============================
# 2) Final: imagen scratch
# ============================
FROM scratch

# Directorio de ejecución
WORKDIR /app

# Copiar binario estático
COPY --from=build /app/server /app/server

# Exponer puerto 
EXPOSE 7600

# Comando de ejecución
ENTRYPOINT ["/app/server"]
