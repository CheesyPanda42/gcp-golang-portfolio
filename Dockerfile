FROM scratch
WORKDIR /app
COPY portfolio /app/portfolio
CMD ["/app/portfolio"]

# Build app and copy
# Use scratch image