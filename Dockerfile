# Step 1: Use the official Go image with Alpine Linux for a smaller size
FROM golang:1.24-alpine


# 🔹 Why Do We Use WORKDIR /app?
# It helps keep things organized inside the container.
# Any files copied into the container will go into /app.
# Any commands you run (RUN, CMD, etc.) will execute inside /app.

# Step 2: Set the working directory inside the container
WORKDIR /app

# Step 3: Copy the Go module files (to cache dependencies)
COPY go.mod go.sum ./

# Step 4: Download Go dependencies
RUN go mod tidy

# Step 5: Copy the rest of the application files
COPY . .

# Step 6: Build the Go application
RUN go build -o main .

# Step 7: Expose the application port
EXPOSE 8080

# Step 8: Run the application
CMD ["./main"]


# docker rm -f my-mongo  # Force remove the existing container
# docker run -d --name my-mongo -p 27018:27017 mongo  # Start a new container on port 27018
