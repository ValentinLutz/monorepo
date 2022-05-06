app/config/config.yaml: ./config/config.none-dev.yaml ## Copy none-dev config to app directory
	install -D config/config.none-dev.yaml app/config/config.yaml

app/config/cert.crt: ./config/none-dev.crt ## Copy none-dev tls crt to app directory
	install -D config/none-dev.crt app/config/cert.crt

app/config/cert.key: ./config/none-dev.key ## Copy none-dev tls key to app directory
	install -D config/none-dev.key app/config/cert.key

app/serve/openapi/orders.yaml: api-definition/orders.yaml ## Copy orders open api definition to app
	install -D api-definition/orders.yaml app/serve/openapi/orders.yaml

app.run:: app/config/config.yaml app/config/cert.crt app/config/cert.key app/serve/openapi/orders.yaml ## Run the app
	cd app && \
		go run main.go

app.build:: ## Build the app into an executable
	cd app && \
		go build
