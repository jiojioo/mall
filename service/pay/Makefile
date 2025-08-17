all:  genapi gensql genrpc
gensql:
	goctl model mysql ddl -src ./model/pay.sql -dir ./model -c

genapi:
	goctl api format -dir api
	goctl api go -api ./api/pay.api -dir ./api
	goctl api plugin -plugin goctl-swagger="swagger -filename oay.json -host 127.0.0.1:8103 -basepath /" -api ./api/pay.api -dir .api//internal/handler/swagger
genrpc:
	goctl rpc protoc ./rpc/pay.proto --go_out=./rpc/types --go-grpc_out=./rpc/types --zrpc_out=./rpc
