# proxy GoAgentX


<https://console.developers.google.com/>

建立google app engine

建立[app密码](https://security.google.com/settings/security/apppasswords)



#error

最近Mac goagent会遇到遇到这样的问题：

	GoAgent Starting...WARNING - [Mar 22 20:05:54] Load Crypto.Cipher.ARC4 Failed, Use Pure Python Instead.

解决方法简单：在终端中执行安装PyCrypto命令后可解决

	easy_install pycrypto

