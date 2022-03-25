image:
	docker build -t sample-keda-metrics-api .
push: image
	docker tag sample-keda-metrics-api adoborroto/sample-keda-metrics-api:latest
	docker push adoborroto/sample-keda-metrics-api:latest
run: image
	docker run -p 8090:8090 -it sample-keda-metrics-api
build: 
	go build -o bin/server .