swagger:
	swagger generate spec -o ./swagger.yaml --scan-models

serve-swagger:
	swagger serve -F=swagger swagger.yaml