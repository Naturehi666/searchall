# Naturehi

欢迎关注我的公众号

嗨嗨安全 ---- 更多好文章等你来读

![图片](https://github.com/sujiawei00/naturahi/assets/58332933/82b33850-748d-4e03-8da0-d6552a4c91c5)

searchall3.5可以快速搜索服务器中的有关username，passsword,账号，口令的敏感信息还有浏览器的账户密码。

项目已开源，希望大家能够理解我的代码写的或许很乱，很糟糕的问题，谢谢！
欢迎批评与指正


如果你觉得这个项目对你有帮助，你可以请作者喝奶茶

![image](https://github.com/Naturehi666/searchall/assets/58332933/07759057-9072-457c-a378-5d3aab611cd1)






开发者日志

5.18日 更新3.1版本

1.解决已知bug
2.增加规则库，进行过滤。
3.将扫出同一文件下的账户，密码整合到一起
4.追加时间戳。
5.search.txt 文件不在是覆盖数据，而是追加数据。


6.6日 更新3.2版本

1.解决已知bug
2.增加规则库，例如账号，密码，口令进行过滤。


6.24日 更新3.4版本

1.解决已知bug
2.识别是否安装向日葵，并自动解析历史连接记录


![image](https://github.com/Naturehi666/naturehi/assets/58332933/2a038208-c428-4b72-b823-9f51e7e2d26a)


7.12日 更新3.5版本

1.解决已知bug
2.增加读取浏览器账户，密码功能

感谢大佬！参考了大佬的工具 https://github.com/moonD4rk/HackBrowserData 集成到了一起。
![image](https://github.com/Naturehi666/naturehi/assets/58332933/62b8fb88-c986-4be5-b043-921df5ed8de8)


7.13日 更新3.5.1版本


1.解决linux输入根目录报错问题，目前已禁止根目录扫描


![image](https://github.com/Naturehi666/searchall/assets/58332933/a14f513c-3b2c-4634-b184-4af9595b8f0b)



7.14日 更新3.5.2版本


1.彻底解决linux输入根目录问题
2.解决占用cpu,内存过高等问题


![image](https://github.com/Naturehi666/searchall/assets/58332933/f4260ef6-cf01-4ad1-ad6d-3360d9e0a6ba)



![image](https://github.com/Naturehi666/searchall/assets/58332933/c708169b-56b5-4e0c-acf4-cffdf5dc8733)



7.19日 更新3.5.3版本


1.解决因线程过多导致目录扫不全问题
2.增加/var/log/secure 扫描时取出登陆成功记录
3.增加accessKeyId,accessKeySecret,jdbc匹配规则
4.增加浏览器生成的result文件夹打包功能
5.增加扫描有效文件数实时响应功能
6.增加扫描服务器中是否安装docker功能


![image](https://github.com/Naturehi666/searchall/assets/58332933/440229d6-f0c3-473f-a612-c3ed7b289400)





8.12日 更新3.5.4版本


增加 CorpId，CorpSecret，qq.im.sdkappid，qq.im.privateKey，qq.im.identifier的读取规则



![image](https://github.com/Naturehi666/searchall/assets/58332933/02f063c9-5227-4534-b654-2e93cfa62560)



8.12日 更新3.5.5版本


用户可自定义字段进行检索查询

![image](https://github.com/Naturehi666/searchall/assets/58332933/85a6512a-f4ad-459b-8d76-94033d158896)



8.30日 更新3.5.6版本

此版本可在谷歌运行时读取Cookie

![image](https://github.com/Naturehi666/searchall/assets/58332933/94c82800-59d0-4992-9182-c786b27de08f)









下一步开发计划

加入 doc，ppt，xls读取功能，敬请期待

鸣谢

感谢网上开源的相关项目！

免责声明

本工具仅能在取得足够合法授权的企业安全建设中使用，在使用本工具过程中，您应确保自己所有行为符合当地的法律法规。 如您在使用本工具的过程中存在任何非法行为，您将自行承担所有后果，本工具所有开发者和所有贡献者不承担任何法律及连带责任。 除非您已充分阅读、完全理解并接受本协议所有条款，否则，请您不要安装并使用本工具。 您的使用行为或者您以其他任何明示或者默示方式表示接受本协议的，即视为您已阅读并同意本协议的约束。







