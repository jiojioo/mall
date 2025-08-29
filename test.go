package main

import (
	"fmt"
	"github.com/dtm-labs/dtmgrpc"
)

func main() {
	dtmServer := "etcd://127.0.0.1:2379/dtmservice"

	// 尝试生成 GID
	gid := dtmgrpc.MustGenGid(dtmServer)
	
	fmt.Println("成功生成 GID:", gid)
}
