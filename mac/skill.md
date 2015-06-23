# mac skill


sudo scutil --set HostName yourname

sudo dseditgroup -o edit -a your_username -t user _www

mac os x 终端terminal打开速度很慢 (2013-05-20 14:11:33)转载▼
原因：大量日志累计造成的
解决：sudo rm -rf /private/var/log/asl/*.asl
