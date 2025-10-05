generate-swagger:
	swagger generate server -f ./routes/swagger/userSwagger.yaml -t ./controllers/target

run: generate-swagger
	go run main.go
