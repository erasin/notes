# google服务框架

获得 ROOT权限，安装RE(root explorer)文件管理器。

## 文件位置

`/data/app` 为安装的app位置

`/system/app` 为系统app位置

`/etc/hosts` 为 hosts 文件位置

## gapps

到 [gapps](//goo.im/gapps) 下载文件包。 一般情况下，使用 `recovery` 刷机安装包裹即可。

解压gapps后查看`system/app`，直接使用RE替换需要用的app即可。

文件	 							| 功用
--------------------------------|----------------
ChromeBookmarksSyncAdapter.apk 	| 浏览器同步
ConfigUpdater.apk 				| 
GenieWidget.apk 				|
GmsCore.apk 					|
GoogleBackupTransport.apk 		|
GoogleCalendarSyncAdapter.apk 	| 日历同步
GoogleContactsSyncAdapter.apk 	| 通讯录同步
GoogleFeedback.apk 				|
GoogleLoginService.apk 			| 帐号同步登陆
GooglePartnerSetup.apk 			|
GoogleServicesFramework.apk 	| google服务框架
GoogleTTS.apk 					| TTS
LatinImeDictionaryPack.apk 		|
MediaUploader.apk 				|
NetworkLocation.apk 			| 利用网络进行定位
OneTimeInitializer.apk 			| 
Phonesky.apk 					|
QuickSearchBox.apk 				|
SetupWizard.apk 				|
Talk.apk 						|
Talkback.apk 					|
VoiceSearchStub.apk 			|


NetworkLocation.apk 可以使用 百度的 BaiduNetworkLocation.apk 替换，小米用的默认为百度。

利用百度APP下载**google play**，然后使用RE将`/data/app`中apk移动到`/system/app`中，否则会跳出。

## GPS

然后使用RE `/etc/gps.conf`,

	NTP_SERVER=europe.pool.ntp.org
	XTRA_SERVER_1=http://xtra1.gpsonextra.net/xtra.bin
	XTRA_SERVER_2=http://xtra2.gpsonextra.net/xtra.bin
	XTRA_SERVER_3=http://xtra3.gpsonextra.net/xtra.bin
	SUPL_HOST=supl.google.com
	SUPL_PORT=7276

将第一行`europe.pool.ntp.org`修改为`cn.pool.ntp.org`。
