package flag

import "flag"

func InitFlag(port, runMode, configPath *string, isVersion *bool, serviceName, agentHostPort *string) error {
	flag.StringVar(port, "port", "", " 启动端口")
	flag.StringVar(runMode, "mode", "", " 启动模式")
	flag.StringVar(configPath, "path", "configs/", "配置文件路径")
	flag.BoolVar(isVersion, "version", false, "编译信息")
	flag.StringVar(serviceName, "serviceName", "blog-service", "Jaeger追踪服务名称")
	flag.StringVar(agentHostPort, "agentHostPort", "localhost:6831", "Jaeger客户端端口")
	flag.Parse()
	return nil
}
