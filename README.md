# hulk
基于项目的配置下发中心. 可以准确，快速的下发配置数据。提供SDK，简化客户端操作。

# How to use Hulk?

Hulk提供SDK(当前提供golang版本), sdk在启动时会自动完成初始化和环境变量注入。 为了能让sdk正常工作，需要准备以下三个环境变量:

* HULK_ENDPOINT Hulk服务API地址.
* HULK_PROJECT_NAME 需要获取的项目名称
* HULK_PROJECT_VERSION 配置版本

```golang
import _ "github.com/andy-zhangtao/hulk/sdk/hulk_go"

func main(){
    hulk_go.Run() //此行可有可无. 如果没有，在使用dep管理依赖时，会认为hulk_go没有被真正使用，导致无法引入到Vendor中
}
```