# -----------------------------------------------------------
# Etapa 1: Construcción del binario Go
# -----------------------------------------------------------
FROM golang:1.25 AS builder

# Establece el directorio de trabajo dentro del contenedor
WORKDIR /app

# Copia los archivos go.mod y go.sum para instalar dependencias primero
COPY go.mod go.sum ./

# Descarga las dependencias del proyecto
RUN go mod download

# Copia el resto del código fuente al contenedor
COPY . .

# Compila la aplicación en modo release (binario estático)
RUN CGO_ENABLED=0 GOOS=linux go build -o backend ./cmd/api

# -----------------------------------------------------------
# Etapa 2: Imagen final para producción
# -----------------------------------------------------------
FROM debian:bullseye-slim

# Actualiza los paquetes del sistema para reducir vulnerabilidades
RUN apt-get update && apt-get upgrade -y && rm -rf /var/lib/apt/lists/*

# Crea un usuario no root por seguridad
RUN useradd -m appuser

# Establece el directorio de trabajo
WORKDIR /app

# Copia el binario compilado desde la etapa anterior
COPY --from=builder /app/backend .

# Copia archivos de migraciones y otros recursos si es necesario
COPY --from=builder /app/cmd/migrate/migrations ./migrations

# Copia archivos de configuración si los usas (descomenta si tienes .env)
# COPY --from=builder /app/.env .env

# Cambia al usuario no root
USER appuser

# Expone el puerto en el que corre la aplicación (ajusta si usas otro)
EXPOSE 8080

# Comando por defecto para ejecutar la aplicación
CMD ["./backend"]

# -----------------------------------------------------------
# Notas:
# - Usa variables de entorno para configuración sensible.
# - Si usas SQLite, asegúrate de que el archivo de base de datos esté en un volumen persistente.
# - Puedes agregar instrucciones para migraciones en la sección CMD o ENTRYPOINT si lo necesitas.