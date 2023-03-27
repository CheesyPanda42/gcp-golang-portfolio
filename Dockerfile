FROM scratch
WORKDIR /app
COPY . /app
EXPOSE 8080
ENV GIN_MODE release
CMD ["/app/portfolio"]

# Build app and copy
# Use scratch image