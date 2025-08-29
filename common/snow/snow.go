package snow

import (
	"errors"
	"fmt"
	"github.com/bwmarrin/snowflake"
)

// 全局雪花节点（单例，避免重复创建）
var snowflakeNode *snowflake.Node

func InitSnowflake(nodeID int64) error {
	if nodeID < 0 || nodeID > 1023 {
		return errors.New("nodeID must be between 0 and 1023")
	}
	node, err := snowflake.NewNode(nodeID)
	if err != nil {
		return fmt.Errorf("init snowflake node failed: %w", err)
	}
	snowflakeNode = node
	return nil
}

// 生成雪花ID（字符串形式）
// 注意：必须先调用InitSnowflake初始化节点，否则返回错误
func NewSnowflakeID() (int64, error) {
	if snowflakeNode == nil {
		return 0, errors.New("snowflake node not initialized")
	}
	id := snowflakeNode.Generate()
	return id.Int64(), nil
}
