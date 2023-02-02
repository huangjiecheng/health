package plugin

import (
	"fmt"
	"health/share"
)

const (
	validatePath = "/usr/local/water/plugins/"
)

type validateArgs struct {
	ComponentName  share.ComponentName       `json:"component_name"`
	StrategyName   share.StrategyName        `json:"strategy_name"`
	StrategyConfig model.StrategyConfigModel `json:"strategy_config"`
}

func ValidateConfigContent(component share.ComponentName, strategy share.StrategyName, config model.StrategyConfigModel) (err error) {
	// 从创建的可执行文件中创建一个新插件。通过 TCP 连接到它
	p := pingo.NewPlugin("tcp", validatePath+"validate_content_plugin")
	// 启动插件
	p.Start()
	// 使用完插件后停止它
	defer p.Stop()

	var (
		args = validateArgs{
			ComponentName:  common.ComponentName(component),
			StrategyName:   common.StrategyName(strategy),
			StrategyConfig: config,
		}
		resp string
	)

	// 从先前创建的对象调用函数
	if err = p.Call("Plugin.ValidateContent", args, &resp); err != nil {
		return fmt.Errorf("配置内容无法通过校验 err:%v", err)
	}

	return
}
