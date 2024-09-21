# Stage 1: Build the frontend
FROM node:20 AS frontend-builder

# Set the working directory to /app/frontend
WORKDIR /app/frontend

# Copy the frontend source code to the container
COPY ./frontend/package.json ./
COPY ./frontend/index.ts ./
COPY ./frontend/tsconfig.json ./
COPY ./frontend/index.html ./

# Install dependencies and build the frontend
RUN npm install
RUN npm run build

# Stage 2: Build the backend
FROM golang:1.23.1 AS backend-builder

# Set the working directory to /app/backend
WORKDIR /app/backend

# Copy the backend source code to the container
COPY ./backend ./

# Build the Go application
RUN go build -o server server.go openai_client.go

# Stage 3: Final stage
FROM golang:1.23.1

# Set the working directory to /app
WORKDIR /app

# Copy the built Go application from the previous stage
COPY --from=backend-builder /app/backend/server ./backend/server
COPY --from=frontend-builder /app/frontend ./frontend

# Expose the port the app runs on
EXPOSE 3000

# Set the working directory to /app/backend
WORKDIR /app/backend

# Command to run the Go application
CMD ["./server"]
