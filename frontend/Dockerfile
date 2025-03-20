FROM node:1.24-alpine

# Set the Current Working Directory inside the container
WORKDIR /app

# Copiar package.json y package-lock.json antes de copiar todo el código
COPY package*.json ./

# Instalar dependencias
RUN npm install

COPY . .

# Build the React app
RUN npm run build

# This container exposes port 8080 to the outside world
EXPOSE 3000