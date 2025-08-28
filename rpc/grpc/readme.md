# 环境安装

```shell
# 安装protobuf编译器
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
# 安装gRPC代码生成工具
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest

# 生成go代码

cd rpc

# 1. 核心工具：protoc
# protoc 是 Protobuf 的官方编译器，用于解析 .proto 文件并生成对应编程语言的代码。本身不直接生成 Go 代码，需要通过「插件」（如 protoc-gen-go、protoc-gen-go-grpc）实现。
# 2. --go_out=./pb
# 作用：指定 Protobuf 核心消息（数据结构）的 Go 代码生成路径。
# 解析：
# --go_out：是 protoc-gen-go 插件的参数前缀，用于配置该插件的输出。
# ./pb：生成的代码会放在当前目录下的（若不存在会自动创建）。
# 生成内容：基于 .proto 中定义的 message（消息类型），生成对应的 Go 结构体、编解码方法等。
# 3. --go_opt=paths=source_relative
# 作用：配置 protoc-gen-go 插件的代码生成路径规则。
# 解析：
# --go_opt：是 protoc-gen-go 插件的额外选项参数。
# paths=source_relative：表示「生成的文件路径与 .proto 源文件保持相对路径一致」。
# 举例：若 calculator.proto 在当前目录，生成的文件会直接放在 ./grpc/calculator.pb.go。
# 替代方案：若使用 paths=import（默认），会根据 .proto 中的 package 声明生成嵌套目录（可能导致路径过深，不推荐）。
# 4. --go-grpc_out=./grpc
# 作用：指定 gRPC 服务代码的生成路径。
# 解析：
# --go-grpc_out：是 protoc-gen-go-grpc 插件的参数前缀，用于配置 gRPC 相关代码的输出。
# ./calculator：与 --go_out 路径保持一致，确保 Protobuf 核心代码和 gRPC 代码在同一目录，便于引用。
# 生成内容：基于 .proto 中定义的 service（服务接口），生成 gRPC 客户端（Stub）和服务端（Server）的框架代码。
# 5. --go-grpc_opt=paths=source_relative
# 作用：配置 protoc-gen-go-grpc 插件的路径规则，与 --go_opt=paths=source_relative 含义一致，确保 gRPC 代码的路径与源文件相对路径一致。
# 6. calculator.proto
# 最后一个参数是待解析的 .proto 源文件路径（当前目录下的 calculator.proto）。

protoc \
    --go_out=./pb  --go_opt=paths=source_relative \
    --go-grpc_out=./pb  --go-grpc_opt=paths=source_relative \
    calculator.proto
```
