
all: backend 
release: backend

.PHONY: backend
backend:
	go build -o ./output/backend ./main.go

