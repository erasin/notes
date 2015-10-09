# rootless

关闭 rootless

    sudo nvram boot-args="kext-dev-mode=1 rootless=0";sudo reboot

打开 rootless

    sudo nvram -d boot-args && sudo reboot


开机按option 选择recovery mode, 然后在恢复模式的终端里运行csrutil disable命令关闭rootless，也可以用csrutil enable命令恢复rootless


http://www.zhihu.com/question/31116473
