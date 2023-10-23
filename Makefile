run_apigw:
	docker-compose --env-file ./containers/api_gateway/.env.example -f docker-compose.api_gateway.yml up

run_auth:
	docker-compose --env-file ./containers/auth/.env.example -f docker-compose.auth.yml up

run_product:
	docker-compose --env-file ./containers/product/.env.example -f docker-compose.product.yml up