## run/api: run the cmd/api application

.PHONY: run/api run/api/win
run/api:
	@set -a && source .envrc && set +a && go run ./cmd/api
run/api/win:
	@powershell -Command "Get-Content .envrc | ForEach-Object { if ($$_ -match '^([^=]+)=(.*)$$') { [System.Environment]::SetEnvironmentVariable($$matches[1], $$matches[2], 'Process') } }; go run ./cmd/api"