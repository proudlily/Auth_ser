你好！
很冒昧用这样的方式来和你沟通，如有打扰请忽略我的提交哈。我是光年实验室（gnlab.com）的HR，在招Golang开发工程师，我们是一个技术型团队，技术氛围非常好。全职和兼职都可以，不过最好是全职，工作地点杭州。
我们公司是做流量增长的，Golang负责开发SAAS平台的应用，我们做的很多应用是全新的，工作非常有挑战也很有意思，是国内很多大厂的顾问。
如果有兴趣的话加我微信：13515810775  ，也可以访问 https://gnlab.com/，联系客服转发给HR。
##asyoume Auth是什么？
* asyoume是极限的传承项目，定义了一系列互联网相关的组件
* auth是一个基于oauth2的用户授权模块,定义了纯粹api形式的接口


##需要实现的功能
* 用户远程授权登陆
* 用户操作日志
* 用户创建自己授权的应用
* 用户授权给应用
* 应用通过api获取用户数据
* 授权平台处理应用回馈的信息
* 用户任务模块
* 用户邮件模块

##使用方法
* 获取依赖  go get github.com/asyoume/postgresql  <br/> go get github.com/asyoume/inotify 
* 安装依赖  go install github.com/asyoume/postgresql/pgsql_map
* 获取源码  go get github.com/asyoume/Auth_ser
* 配置文件  编辑源码目录下conf/app.json文件配置依赖的数据库，依赖的缓存服务
* 运行系统  源码目录下执行sh hack/http.sh启动http协议的服务 || sh hack/run.sh启动thrift协议的服务
