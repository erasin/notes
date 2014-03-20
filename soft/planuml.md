# planuml 语法

需要java包。

安装 [Graphviz](http://192.168.0.78/share/soft/graphviz-2.36.msi)。

下载 [planuml](/)。

在`planuml.jar`同级别文件夹建立`planuml.bat`.

添加内容

    java -jar plantuml.jar -gui 

执行后打开 `.pu` 为后缀的文件夹。

window 下使用ansi编码保存pu文件(默认的txt文档，修改pu为后缀)。


## 活动图


开始与结束

    @startuml
    title "标题"

    --> "文档内容"

    @enduml

进程节点
    
    (*)-->"开始"
    -->  "节点名称"
    -->[箭头的注释]  "节点名称"
    -up->  "节点名称"
    -down->  "节点名称"
    -left->  "节点名称"
    -right->  "节点名称"
    -->(*) 

判断

    if  "介绍" then
        -->[true] "真"
    else
        -->[false] "真"
    endif

单行注释

    note right:  右侧的内容
    note left:  右侧的内容

多行注释

    note right:
    多行内容
    endnote

分区

    partition "分区名称" {
        --> "基本节点"
    }

并发 （不怎么好用）

    --> ===并发名称1===

    --> 节点1

    --> ===并发名称2===


    --> ===并发名称1===

    --> 节点1

    --> ===并发名称2===


例子： fuxi.pu

    @startuml
    title 付息和结款

    (*) --> "准本信息"
    note right
    当前时间和，进行中结束项目的付息结款时间,
    最近一次付息时间,以及状态。
    end note

    --> ===INV_START=== 
    -left-> "付息"
    if "是否为付息日 && 付息锁" then
        -down-> "获取项目ID"
        partition "项目付息计算 in action_model" {
            -->[true] "开始付息模块"
            if "判定用户付息锁"
                --> "已经为用户结束计算"
            else
                --> "继续计算"
            endif

            --> "计算日利息"
            if "是否为首次付息" then
                -->[true] "计算首期天数"
            else
                --> "计算天数"
            endif
            --> "当期计算利息"
            --> "更新用户投资利息已付和未付"
            --> "更新系统付息提醒"
            --> "更新金钱日志记录"
            --> "发送用户消息"
        }
    else
      -->[false] "不做任何处理"
    endif
    -down-> ===INV_END===

    "准本信息" --> ===INV_START=== 
        -right-> "结款"
        if "是否为结款日" then
            -->[true] "开始结款计算"
            --> "处理" 
            partition "项目结款计算 in action_model" {
                --> "格式化时间"
                --> "计算日利息》"
                if "结束时间大于付息时间"
                    --> "获得最后付息天数"
                else
                    --> "获得最后付息天数"
                endif
                --> "当期计算利息》 "
                --> "更新用户投资利息已付和未付》"
                --> "更新系统付息提醒》"
                --> "更新金钱日志记录》"
                --> "发送用户付息消息》"
                --> "为用户结款》"
                --> "发送结款消息》"
                --> "标记项目结束"
            }
        else
            -->[false] "结束计算"
        endif
    -down-> ===INV_END===
    --> (*)

    @enduml