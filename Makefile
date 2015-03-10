all:
	go build && ./ecog
ecog:
	go build
run:
	go run app.go schema.go handlers.go utils.go db.go