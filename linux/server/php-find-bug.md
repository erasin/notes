# php hack check

find ./ -name “*.php” | xargs egrep ‘phpspy|c99sh|milw0rm|eval\(base64_decode|spider_bc|eval\(\$_POST\[' |awk -F: '{print $1}'|sort|uniq >> ${hacking_log}/hacking.log


#find /data/webtest/ -name "*.*" -type f -print0|xargs -0 egrep 'phpspy|c99sh|milw0rm|eval\(base64_decode|spider_bc|eval\(\$_POST\[|system \(\$_REQUEST\[|system\(\$_REQUEST\[|base64|base64_decode|eval \(\$_POST\[|readdir\(|copy\(\$_FILES|copy \(\$_FILES|move_uploaded_file\(\$_FILES|move_uploaded_file \(\$_FILES|cha88.cn|tools88.com|eval\(gzinflate\('|awk -F: '{print $1}'|sort|uniq >> /data/logs/hack123.log




find ./ -name "*.php" |xargs egrep "phpspy|c99sh|milw0rm|eval\(gunerpress|eval\(base64_decode|spider_bc"> /tmp/php.txt

grep -r –include=*.php  '[^a-z]eval($_POST' . > /tmp/eval.txt

grep -r –include=*.php  'file_put_contents(.*$_POST\[.*\]);' . > /tmp/file_put_contents.txt

find ./ -name "*.php" -type f -print0 | xargs -0 egrep "(phpspy|c99sh|milw0rm|eval\(gzuncompress\(base64_decode|eval\(base64_decode|spider_bc|gzinflate)" | awk -F: '{print $1}' | sort | uniq
查找最近一天被修改的PHP文件

find -mtime -1 -type f -name \*.php
修改网站php文件权限，只读

find -type f -name \*.php -exec chmod 444 {} \;
修改网站目录权限

find ./ -type d -exec chmod 555{} \;