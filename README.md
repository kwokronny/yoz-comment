
<center>
<img src="./docs/logo.png" width="100px"/>
<p style="font-size:24px;font-weight:bold;">YozComment</p>
</center>

## 介绍
开源的 golang 评论系统。因自己博客与他人共用服务器，暂时知名的几个系统都是需要在服务器安装一定的依赖或环境，不想增加服务器负担，顺带学习 golang 的态度，便自己造了这个轮子。


- #### 前端

	前端应用 Typescript 开发，开发架构及思路部分主要借鉴 [utterance](https://github.com/utterance/utterances) 。通过在网站内嵌 iframe 减少对代码间的冲突与安全问题。

	后台管理采用纯前端方式通过 CDN 引用 iviewui+axios 完成配置安装页面，管理页面及登陆界面

- #### 后端

	后端运用 gin+gorm 开发，开发架构主要是MVC架构，也应用 ORM 模式方便后面的开发中拓展数据库种类
	- 多级评论功能
	- 后台管理及配置安装
	- 敏感字识别
	- 邮件通知功能
	- JWT 鉴权

## 安装

```bash
	git clone https://github.com/kwokronny/YozComment.git
	cd ${workspace}
	npm install parcel-bundler -g
	npm install
	npm run build
```
![](./docs/install.jpg)