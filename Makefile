include motoEnv.env
export

build-frontend:
	cd web && npm run build

build: build-frontend
	go build -o motoex .

run-dev:
	@echo "Run 'cd web && npm run dev' in another terminal"
	go run .

run: build
	./motoex