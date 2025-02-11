# Gunakan base image Golang versi terbaru
FROM golang:1.21-alpine

# Set environment variable
ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

# Buat direktori kerja
WORKDIR /app

# Copy semua file ke dalam container
COPY . .

# Download semua dependency
RUN go mod tidy

# Build aplikasi
RUN go build -o main .

# Expose port yang digunakan aplikasi
EXPOSE 8009

# Jalankan aplikasi
CMD ["./main"]