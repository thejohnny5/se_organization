# Stage 1: Build the Vue.js frontend
FROM node:latest as build-stage
WORKDIR /app
COPY package*.json ./
RUN npm install
COPY vite.config.js ./
COPY ui/ ui/
COPY index.html .
RUN npm run build

# Stage 2: Build the Go backend
FROM golang:latest as go-build-stage
WORKDIR /go/src/app
COPY go.mod .
COPY go.sum .
COPY pkg/ pkg/
COPY main.go ./
RUN go mod download

RUN CGO_ENABLED=0 GOOS=linux go build -o /go/bin/app

# Stage 3: Create the final image
FROM alpine:latest
WORKDIR /app

# Copy the built frontend from the build-stage
COPY --from=build-stage /app/dist /app/frontend

# Copy the compiled backend binary from the go-build-stage
COPY --from=go-build-stage /go/bin/app /app/backend

# Add any additional runtime dependencies needed
RUN apk add --no-cache ca-certificates

# Expose the port your app runs on
EXPOSE 8080

# Set the binary as the entrypoint of the container
ENTRYPOINT ["/app/backend"]
