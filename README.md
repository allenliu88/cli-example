# cli-example

Quick start for golang cli with cobra/viper/logrus.

## 初始化工程

```shell
mkdir -p /u01/repo/cli-example
cd /u01/repo/cli-example
go mod init github.com/allenliu88/cli-example

## 获取及安装cobra
go get -u github.com/spf13/cobra@latest
go install github.com/spf13/cobra-cli@latest

## 初始化工程目录
cobra-cli init --author "allen liu" --license apache --viper

## 添加info子命令
cobra-cli add info

$ tree
.
├── LICENSE
├── cmd
│   ├── info.go
│   └── root.go
├── go.mod
├── go.sum
└── main.go

1 directory, 6 files

## 安装日志依赖
go get -u github.com/sirupsen/logrus@latest
go install github.com/sirupsen/logrus@latest
```

## 特殊注意

这里面有几点需要注意：

- `infoCmd.Flags()`中并不包含当前命令的`PersistentFlags()`集合，虽然该函数的注释中写的是包含【**Flags returns the complete FlagSet that applies to this command (local and persistent declared here and by all parents).**】这里比较坑。因此，要想通过`viper.BindPFlags`完整绑定`Local Flags`及`Persistent Flags`，则需要同时执行`viper.BindPFlags(infoCmd.Flags())`及`viper.BindPFlags(infoCmd.PersistentFlags())`即可。
- 可通过如下`bindViper()`函数将viper变量反向绑定到Flag Set中，同样如上，如果需要`Local Flags`及`Persistent Flags`同时生效，则需要同时执行`bindViper(infoCmd.Flags())`及`bindViper(infoCmd.PersistentFlags())`。
- `viper.AutomaticEnv()`必须要在`viper.SetEnvPrefix(envPrefix)`之后调用，否则，环境变量固定前缀不会生效。
- 由`debugEnabled`Flags可知，当通过`viper.BindPFlags()`自动绑定到viper后，其环境变量格式为`CLI_DEBUGENABLED`，其中，`CLI`为固定环境变量前缀，而变量名部分就是`debugEnabled`直接全大写了，与常规的小驼峰不同单词之间通过下划线`_`的逻辑不符（Spring Boot中是通过下划线`_`分隔处理的）。
