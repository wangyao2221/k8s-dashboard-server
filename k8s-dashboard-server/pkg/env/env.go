package env

import (
	"flag"
	"fmt"
	"strings"
)

var (
	active Environment
	// 使用接口+私有struct实现了枚举变量的功能
	dev Environment = &environment{value: "dev"}
	fat Environment = &environment{value: "fat"}
	uat Environment = &environment{value: "uat"}
	pro Environment = &environment{value: "pro"}
)

// TODO 搞清楚这种声明的作用
var _ Environment = (*environment)(nil)

// Environment 环境配置
type Environment interface {
	Value() string
	IsDev() bool
	IsFat() bool
	IsUat() bool
	IsPro() bool
	t()
}

type environment struct {
	value string
}

func (e *environment) Value() string {
	return e.value
}

func (e *environment) IsDev() bool {
	return e.value == "dev"
}

func (e *environment) IsFat() bool {
	return e.value == "fat"
}

func (e *environment) IsUat() bool {
	return e.value == "uat"
}

func (e *environment) IsPro() bool {
	return e.value == "pro"
}

func (e *environment) t() {}

// 在这里去取env的变量，就不用把active暴露出去了，之提供接口实现让使用者获取变量
func init() {
	env := flag.String("env", "", "请输入运行环境:\n dev:开发环境\n fat:测试环境\n uat:预上线环境\n pro:正式环境\n")
	flag.Parse()

	switch strings.ToLower(strings.TrimSpace(*env)) {
	case "dev":
		active = dev
	case "fat":
		active = fat
	case "uat":
		active = uat
	case "pro":
		active = pro
	default:
		active = fat
		fmt.Println("Warning: '-env' cannot be found, or it is illegal. The default 'fat' will be used.")
	}
}

// Active 当前配置的env
func Active() Environment {
	return active
}
