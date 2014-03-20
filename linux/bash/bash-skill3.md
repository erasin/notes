# 技巧


方法一，指定换行符读取：

    #! /bin/bash  
      
    for LINE in `cat /etc/passwd`  
    do   
            echo $LINE 
    done
 
方法二，文件重定向给read处理：

    #! /bin/bash  
      
    cat /etc/passwd | while read LINE  
    do
            echo $LINE 
    done
 
 
方法三，用read读取文件重定向：

    #! /bin/bash  
      
    while read LINE
    do
            echo $LINE 
    done < /etc/passwd
 
访问二和三比较相似，推荐用方法三