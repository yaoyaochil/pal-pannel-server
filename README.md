# 幻兽帕鲁 服务器管理系统

[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)
[![Golang](https://img.shields.io/badge/Golang-1.21-blue)](https://reactjs.org/)


欢迎来到 pal-pannel-web 服务器管理系统的 GitHub 仓库！该系统旨在简化服务器管理任务，为您提供一个全面的解决方案，以高效管理您的服务器。

## 特性

- **用户友好的界面**：直观的界面设计使用户能够轻松地导航和执行任务。
- **实时监控玩家**：您可以实时监控服务器上的玩家，以便及时采取行动。
- **更安全更自由的保管存档**：通过简单的步骤，您可以轻松地备份和还原服务器数据。

## 截图

以下是 pal-pannel 服务器管理系统目前的一些截图：

<img width="1919" alt="image" src="https://github.com/yaoyaochil/pal-pannel-web/assets/49603204/af7a8bff-3730-46c1-9ee6-38e52d03aed8">
<img width="1919" alt="image" src="https://github.com/yaoyaochil/pal-pannel-web/assets/49603204/60dc5665-53ba-4498-9d2d-3b132e60b7f4">
<img width="1918" alt="image" src="https://github.com/yaoyaochil/pal-pannel-web/assets/49603204/f20598ac-5a1c-43d8-ac3d-1acc9b5c3b6c">




## 快速开始

要开始使用 pal-pannel-web 服务器管理系统，请按照以下步骤操作：

1. **克隆仓库**：使用以下命令将此仓库克隆到您的本地计算机：
```bash
git clone https://github.com/yaoyaochil/pal-pannel-server.git
```

2. **安装依赖**：进入项目目录并安装依赖：
```bash
go mod tidy
```

3. **配置环境变量**：修改config.yaml中的palu->server-path为你的linux服务端根目录：

```bash

# 服务器配置
palu:
  server-path: "/home/yaoyaochil/palu-server"
```
4. **运行应用程序**：通过运行以下命令启动应用程序：

```bash
go run main.go
```

4. **访问应用程序**：在浏览器中访问接口`http://localhost:8300`。

5. **访问web端**：在浏览器中访问接口`http://localhost:5173`。

##### 注意：您需要配合 pal-pannel-web 前端项目使用，您可以在这里找到前端项目：[pal-pannel-web](https://github.com/yaoyaochil/pal-pannel-web.git)


## 贡献

我们欢迎社区对 pal-pannel 服务器管理系统进行贡献。如果您想贡献，请遵循以下准则：

- 派生仓库并为您的功能或错误修复创建一个新分支。
- 进行您的更改，并确保代码通过所有测试。
- 提交一个带有您更改清晰描述的拉取请求。

## 联系我

如果您有任何问题或建议，请添加我的wechat，我们把酒言谈。

![101709224238_ pic](https://github.com/yaoyaochil/pal-pannel-server/assets/49603204/aed520f2-5f82-4cde-a0b2-89e64437c999)
## 许可证
MIT