all:  genapi gensql genrpc
gensql:
	goctl model mysql ddl -src ./model/order.sql -dir ./model -c

genapi:
	goctl api format -dir api
	goctl api go -api ./api/order.api -dir ./api
genrpc:
	goctl rpc protoc ./rpc/order.proto --go_out=./rpc/types --go-grpc_out=./rpc/types --zrpc_out=./rpc
