#inux操作系统早已支持局域网唤醒(WOL)

Linux操作系统早已支持局域网唤醒(WOL)，而下一版本的Linux 3.1 Kernel系统内核将会引入__无线局域网唤醒__，简称__WWOL__。

和有线版本类似，无线局域网唤醒可以让系统在进入挂起状态(ACPI S3)或者另一种低功耗状态的同时继续保持无线连接处于开启状态，而一旦接到用户的特定操作，或者从AP断开连接，就会将系统唤醒，并恢复到原始状态。

该技术也可以用于在无线网卡找到合适的Wi-Fi连接时自动唤醒系统。

Linux 3.1对无线局域网唤醒的支持来自于五月份就发布的一个补丁，并最终在merge window阶段成功进入了主干树(mainline tree)。Intel也在他们的iwlagn驱动中加入了对该技术的支持。

Linux 3.1内核正式版将于下个月发布，新功能除了无线局域网唤醒之外还有Nouveau驱动在GeForce 400/500系列显卡上无需手动载入微代码、Sandy Bridge性能优化、大量DRM改进、支持开源处理器架构OpenSIRC、Wii Remote HID驱动、Intel Atom平台Poulsbo驱动改进和Cedar Trail初步支持

原文:[http://www.oschina.net/]
