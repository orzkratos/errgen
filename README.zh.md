# errgen

errgenkratos 的兼容性包装器，提供经典的 protoc-gen-go-errors 命令，同时使用最新的实现。

## 英文文档

[ENGLISH README](README.md)

## 概述

此项目为现有的 protoc-gen-go-errors 用户提供向后兼容性，同时将所有功能委托给新的 [errgenkratos](https://github.com/orzkratos/errgenkratos) 项目。

**对于新项目，我们推荐直接使用 [errgenkratos](https://github.com/orzkratos/errgenkratos) 以获取最新功能和改进。**

## 安装

```bash
go install github.com/orzkratos/errgen/cmd/protoc-gen-go-errors@latest
```

## 使用方法

```bash
protoc --go-errors_out=paths=source_relative:./your_output_dir your_proto_files.proto
```

### 生成代码示例

```go
// 由 protoc-gen-go-errors 生成

// 用户未找到
func IsUserNotFound(err error) bool {
    return errgenkratos.IsError(err, ErrorReason_USER_NOT_FOUND, 404)
}

// 用户未找到  
func ErrorUserNotFound(format string, args ...interface{}) *errors.Error {
    return errgenkratos.NewError(404, ErrorReason_USER_NOT_FOUND, format, args...)
}
```

### 设置自定义元数据字段

在业务代码中配置元数据字段：

```go
import "github.com/orzkratos/errgen"

func init() {
    // 设置自定义元数据字段名
    errgen.SetMetadataFieldName("codeNum")
}
```

## 迁移到 errgenkratos

对于新功能和活跃开发，请考虑迁移到 [errgenkratos](https://github.com/orzkratos/errgenkratos)：

```bash
# 安装新插件
go install github.com/orzkratos/errgenkratos/cmd/protoc-gen-orzkratos-errors@latest

# 更新 protoc 命令
protoc --orzkratos-errors_out=paths=source_relative:./your_output_dir your_proto_files.proto

# 更新导入
import "github.com/orzkratos/errgenkratos"
```

## 环境要求

- Go 1.25.0+
- Protocol Buffers 编译器 (protoc)
- Kratos v2.8.0+

---

<!-- TEMPLATE (ZH) BEGIN: STANDARD PROJECT FOOTER -->

## 📄 许可证类型

MIT 许可证。详见 [LICENSE](LICENSE)。

---

## 🤝 项目贡献

非常欢迎贡献代码！报告 BUG、建议功能、贡献代码：

- 🐛 **发现问题？** 在 GitHub 上提交问题并附上重现步骤
- 💡 **功能建议？** 创建 issue 讨论您的想法
- 📖 **文档疑惑？** 报告问题，帮助我们改进文档
- 🚀 **需要功能？** 分享使用场景，帮助理解需求
- ⚡ **性能瓶颈？** 报告慢操作，帮助我们优化性能
- 🔧 **配置困扰？** 询问复杂设置的相关问题
- 📢 **关注进展？** 关注仓库以获取新版本和功能
- 🌟 **成功案例？** 分享这个包如何改善工作流程
- 💬 **意见反馈？** 欢迎所有建议和宝贵意见

---

## 🔧 代码贡献

新代码贡献，请遵循此流程：

1. **Fork**：在 GitHub 上 Fork 仓库（使用网页界面）
2. **克隆**：克隆 Fork 的项目（`git clone https://github.com/yourname/repo-name.git`）
3. **导航**：进入克隆的项目（`cd repo-name`）
4. **分支**：创建功能分支（`git checkout -b feature/xxx`）
5. **编码**：实现您的更改并编写全面的测试
6. **测试**：（Golang 项目）确保测试通过（`go test ./...`）并遵循 Go 代码风格约定
7. **文档**：为面向用户的更改更新文档，并使用有意义的提交消息
8. **暂存**：暂存更改（`git add .`）
9. **提交**：提交更改（`git commit -m "Add feature xxx"`）确保向后兼容的代码
10. **推送**：推送到分支（`git push origin feature/xxx`）
11. **PR**：在 GitHub 上打开 Pull Request（在 GitHub 网页上）并提供详细描述

请确保测试通过并包含相关的文档更新。

---

## 🌟 项目支持

非常欢迎通过提交 Pull Request 和报告问题来为此项目做出贡献。

**项目支持：**

- ⭐ **给予星标**如果项目对您有帮助
- 🤝 **分享项目**给团队成员和（golang）编程朋友
- 📝 **撰写博客**关于开发工具和工作流程 - 我们提供写作支持
- 🌟 **加入生态** - 致力于支持开源和（golang）开发场景

**使用这个包快乐编程！** 🎉

<!-- TEMPLATE (ZH) END: STANDARD PROJECT FOOTER -->

---

## GitHub 标星点赞

[![Stargazers](https://starchart.cc/orzkratos/errgen.svg?variant=adaptive)](https://starchart.cc/orzkratos/errgen)