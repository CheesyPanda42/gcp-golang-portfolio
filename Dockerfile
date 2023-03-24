FROM golang
WORKDIR /app
COPY . /app
RUN go build -o portfolio
CMD ["./portfolio"]