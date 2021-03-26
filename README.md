
<center>
<img src="./docs/logo.png" width="100px"/>
<p style="font-size:24px;font-weight:bold;">YozComment</p>
<p>作者：<a href="https://kwokronny.top">KwokRonny</a></p>
</center>

## 介绍
开源的 golang 评论系统。因自己博客与他人共用服务器，暂时知名的几个系统都是需要在服务器安装一定的依赖或环境，不想增加服务器负担，顺带学习 golang 的态度，便自己造了这个轮子。

![](./docs/preview.jpg)

- ### 优势
	* 支持MySQL、SQLite、PostgreSQL
	* 多次评论
	* 部署简单，可视化配置
	* 支持 明/暗 主题
	* 支持响应式
	* 接入Gravatar头像显示
	* 敏感词识别
	

- #### 前端

	前端应用 Typescript 开发，开发架构及思路部分主要借鉴 [utterance](https://github.com/utterance/utterances) 。通过在网站内嵌 iframe 减少对代码间的冲突与安全问题。

	后台管理采用纯前端方式通过 CDN 引用 iviewui+axios 完成配置页面，管理页面及登陆界面

- #### 后端

	后端运用 gin+gorm 开发，开发架构主要是MVC架构
	- 多级评论功能
	- 后台管理及配置安装
	- 敏感字识别
	- 邮件通知功能
	- JWT 鉴权

## 安装

- ### 部署至服务器
	```bash
	
	```

- ### 配置安装
	运行脚本当检测到相对目录下 `config/config.yaml` 不存在，访问 http://localhost:9975 会进入安装配置页面
	
	![](./docs/install.jpg)

	根据配置页面操作完成后会在相应位置生成 `config.yaml` 配置文件，并重新运行脚本启动服务

- ### 在页面中放入引用代码完成评论系统部署

	```html
	<script 
		id="YozComment" 
		src="http://localhost:9975/client.js" 
		token="页面唯一TOKEN" 
		theme="{light|dark}" 
		crossorigin="anonymous" 
		async></script>
	```