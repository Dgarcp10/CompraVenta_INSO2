FROM golang:1.24-alpine

# Set the Current Working Directory inside the container
WORKDIR /app

COPY . .

# Copiar los archivos del frontend desde la carpeta build
COPY --from=frontend /app/build ./frontend/build

# Build the Go app
RUN go build -o COMPRAVENTA_INSO2 ./cmd/CV

# This container exposes port 8080 to the outside world
EXPOSE 8080

# Run the binary program produced by `go build`
CMD ["./COMPRAVENTA_INSO2"]