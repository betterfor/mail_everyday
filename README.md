# mail_everyday

使用Github Action定时发送邮件，每天汇报天气和头图。

添加了新的功能：

- 使用Github Action定时任务触发，无需部署自己的服务器
- 获取ONE接口数据，利用API方式
- 使用邮件发送html

## 效果图

![](xxx)

## 如何获取邮件数据

### 1、获取天气预报

爬取墨迹天气的页面，使用[goquery](github.com/PuerkitoBio/goquery)获取天气数据，包括最近三天的天气和天气提示。

### 2、获取ONE头图

使用api访问ONE页面，获取token，然后使用token查询今日的图文信息

### 3、如何渲染出页面

使用template定义好[模板](https://github.com/betterfor/mail_everyday/blob/main/mail.tpl)，然后将数据注入

> 模板和思路整体参考 [NodeMail](https://github.com/Vincedream/NodeMail)

### 4、如何发送邮件

使用Github Action定时任务每天触发编译

```yaml
      - name: Send Mail
        uses: betterfor/action-send-mail@main
        with:
          # 必需，邮件服务器地址
          server_address: smtp.qq.com
          # 必需，邮件服务器端口，默认25 (如果端口为465，则会使用TLS连接)
          server_port: 465
          # 可选 (建议): 邮件服务器用户
          username: ${{secrets.MAIL_USERNAME}}
          # 可选 (建议): 邮件服务器密码
          password: ${{secrets.MAIL_PASSWORD}}
          # 必需，邮件主题
          subject: 一封暖暖的邮件
          # 必需，收件人地址
          to: ${{secrets.MAIL_TO}}
          # 必需，发送人全名 (地址可以省略)
          from: betterfor # <alice@example.com>
          # 可选，HTML内容，可从文件读取
          html_body: file://output/output.html
```         

其中的secrets在Settings->Secrets->Action中定义，Github Action会读取这些变量。

> MAIL_USERNAME和MAIL_PASSWORD是邮箱的账号密码，密码是邮箱的授权码，`server_address`需要根据具体的邮件服务提供商

## 最后

一封暖暖的邮件，能让github代码每天都有提交内容，实现自动化邮件。

**暖暖的，很贴心。**
