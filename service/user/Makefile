all:  genapi gensql genrpc
gensql:
	goctl model mysql ddl -src ./model/user.sql -dir ./model -c

genapi:
	goctl api format -dir api
	goctl api go -api ./api/user.api -dir ./api
#	goctl api plugin -plugin goctl-swagger="swagger -filename qyhxplatform.json -host 127.0.0.1:8825 -basepath /" -api api/main.api -dir ./internal/handler/swagger
genrpc:
	goctl rpc protoc ./rpc/user.proto --go_out=./rpc/types --go-grpc_out=./rpc/types --zrpc_out=./rpc
