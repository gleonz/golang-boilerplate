make run-gen:
swagger generate server -t internal/delivery/http -f ./swagger/swagger.yml --exclude-main -A server