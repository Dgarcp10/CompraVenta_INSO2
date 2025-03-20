FROM node:1.24-alpine

# Set the Current Working Directory inside the container
WORKDIR /app

COPY . .

# Build the React app
RUN npm run build

# This container exposes port 8080 to the outside world
#EXPOSE 5050

# Run the binary program produced by `go build`
#CMD ["./COMPRAVENTA_INSO2"]