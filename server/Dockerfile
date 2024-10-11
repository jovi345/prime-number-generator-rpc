# menggunakan image go:1.23 sebagai base image
FROM golang:1.23

# membuat working directory
WORKDIR /app

# meng-copy source code ke dalam image
COPY . .

# mem-build server
RUN go build -o server main.go

# perintah untuk menjalankan server
CMD ["./server"]