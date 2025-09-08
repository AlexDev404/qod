## run/api: run the cmd/api application

.PHONY: run/api run/api/win
run/api:
	@set -a && source .envrc && set +a && go run ./cmd/api
run/api/win:
	@powershell -Command "Get-Content .envrc | ForEach-Object { if ($$_ -match '^([^=]+)=(.*)$$') { $$value = $$matches[2] -replace '^\"(.*)\"$$', '$$1'; [System.Environment]::SetEnvironmentVariable($$matches[1], $$value, 'Process') } }; go run ./cmd/api"