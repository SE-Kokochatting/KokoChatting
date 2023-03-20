软件工程课程大作业 [Kokochatting](https://github.com/SE-Kokochatting/KokoChatting)，是一款仿 web 端 Telegram 的在线聊天室，可以[在线使用](http://kokochatting.danmoits.com/home)（注册两个用户并使用不同浏览器访问）。除了在线聊天，还支持 ① 登陆注册 ② 渲染好友列表和对话列表 ③ 添加好友、在线收到好友申请通知 ④ 删除或拉黑好友（拉黑后不再收到申请）⑤ 聊天时显示已读状态 ⑥ 查看用户信息 ⑦ 修改自己的头像。

- 使用 react 进行组件化开发，mobx 状态管理，ts 类型约束，对 fetch 进行简单的封装来发送 http 请求。
- 使用包管理器 pnpm，完成 vite.config、eslint、pretty、tsconfig、lint-staged 等配置。

- 先创建新的分支上完成开发，最后提交 pr 合并。commit message 使用 git-commit-plugin 进行规范。

- 使用 scss 变量，实现亮色模式和暗色模式的切换。

- 使用 svg-sprite 技术来渲染图标。

- 使用 websocket 实现实时相关的功能。客户端主动向服务端发送心跳包，如果长期未得到回复则认为已经断线，自动重连。

- 使用 flex 布局，具有一定的响应式，因为主打 pc 端，没有使用媒体查询。

- 使用 eventemitter，订阅事件并在特定的时机触发。比如自己或对方发送消息时调用滚动到页面的底部的函数，抑或是 websocket 断开后调用重连函数。

- 聊天时已读状态使用 IntersectionObserver Api，当某条消息渲染在页面上，回调函数将被触发，此时发送请求并附上该消息的 id，借由 websocket，对方可以立刻得知消息已阅。
