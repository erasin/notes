# planuml 语法


需要java包。

安装 [Graphviz](http://192.168.0.78/share/soft/graphviz-2.36.msi)。

下载 [planuml](http://192.168.0.78/share/soft/planuml.jar)。

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



## Activity Diagram 新版流程图

Current syntax for activity diagram has several limitations and drawbacks (for example, it's difficult to maintain).
So a completely new syntax and implementation is proposed as beta version to users (starting with V7947), so that we could define a better format and syntax. 
Another advantage of this new implementation is that it's done without the need of having GraphViz installed (as for sequence diagrams).

The new syntax will replace the old one. However, for compatibility reason, the old syntax will still be recognized, to ensure ascending compatibility. 
Users are simply encouraged to migrate to the new syntax.

Some features are still missing : partition, skins, notes, error management, titles... But the product is already useable to have simple or very complex diagrams.

If you have ideas or find issues, please provide feedback so that we could enhance this new syntax.

### Simple Activity

Activities label starts with : and ends with ;
Text formatting can be done using creole wiki syntax.

They are implicitely linked in their definition order.

    @startuml

    :Hello world;

    :This is on defined on
    several **lines**;

    @enduml
        

### Start/Stop

You can use start and stop keywords to denote the beginning and the end of a diagram.

    @startuml

    start

    :Hello world;

    :This is on defined on
    several **lines**;

    stop

    @enduml
        

### Conditional

You can use if, then and else keywords to put tests if your diagram. Labels can be provided using parentheses.

    @startuml

    start

    if (graphviz installed?) then (yes)
      :process all\ndiagrams;
    else (no)
      :process only
      __sequence__ and __activity__ diagrams;
    endif

    stop

    @enduml
        

### Repeat loop

You can use repeat and repeatwhile keywords to have repeat loops.

    @startuml

    start

    repeat
      :read data;
      :generate diagrams;
    repeat while (more data?)

    stop

    @enduml
        

### While loop

You can use while and end while keywords to have repeat loops.

    @startuml

    start

    while (data available?)
      :read data;
      :generate diagrams;
    endwhile

    stop

    @enduml
        

It is possible to provide a label after the endwhile keyword, or using the is keyword.

    @startuml
    while (check filesize ?) is (not empty)
      :read file;
    endwhile (empty)
    :close file;
    @enduml
        

### Parallel processing

You can use fork, fork again and end fork keywords to denote parallel processing.
@startuml

start

if (multiprocessor?) then (yes)
  fork
    :Treatment 1;
  fork again
    :Treatment 2;
  end fork
else (monoproc)
  :Treatment 1;
  :Treatment 2;
endif


@enduml
        

### Notes

Text formatting can be done using creole wiki syntax.

    @startuml

    start
    :foo1;
    note left: This is a note
    :foo2;
    note right
      This note is on several
      //lines// and can
      contain <b>HTML</b>
      ====
      * Calling the method ""foo()"" is prohibited
    end note
    stop

    @enduml
        

### Color

You can use specify a color for some activities.

    @startuml

    start
    :starting progress;
    #red:reading configuration files
    These files must do be edited at this point!;
    #AAAAAA:ending of the process;

    @enduml
        

Complete example

    @startuml

    start
    :ClickServlet.handleRequest();
    :new page;
    if (Page.onSecurityCheck) then (true)
      :Page.onInit();
      if (isForward?) then (no)
        :Process controls;
        if (continue processing?) then (no)
          stop
        endif
        
        if (isPost?) then (yes)
          :Page.onPost();
        else (no)
          :Page.onGet();
        endif
        :Page.onRender();
      endif
    else (false)
    endif

    if (do redirect?) then (yes)
      :redirect process;
    else
      if (do forward?) then (yes)
        :Forward request;
      else (no)
        :Render page template;
      endif
    endif

    stop

    @enduml


## Sequence Diagram 时序图

The description language is highly derived from http://www.websequencediagrams.com .

Let's have few example :

Basic examples

The sequence "->" is used to draw a message between two participants. Participants do not have to be explicitly declared.
To have a dotted arrow, you use -->

It is also possible to use <- and <--. That does not change the drawing, but may improve readability.

    @startuml
    Alice -> Bob: Authentication Request
    Bob --> Alice: Authentication Response

    Alice -> Bob: Another authentication Request
    Alice <-- Bob: another authentication Response
    @enduml
        

### Comments

Everything that starts with simple quote ' is a comment.

You can also put comments on several lines using /' to start and '/ to end.

### Declaring participant

It is possible to change participant order using the participant keyword. .
It is also possible to use other keywords to declare a participant.

* actor
* boundary
* control
* entity
* database

    @startuml
    actor Foo1
    boundary Foo2
    control Foo3
    entity Foo4
    database Foo5
    Foo1 -> Foo2 : To boundary
    Foo1 -> Foo3 : To control
    Foo1 -> Foo4 : To entity
    Foo1 -> Foo5 : To database

    @enduml
        

You can rename a participant using the as keyword.

You can also change the background color of actor or participant.

    @startuml
    actor Bob #red
    ' The only difference between actor
    'and participant is the drawing
    participant Alice
    participant "I have a really\nlong name" as L #99FF99
    /' You can also declare:
       participant L as "I have a really\nlong name"  #99FF99
      '/

    Alice->Bob: Authentication Request
    Bob->Alice: Authentication Response
    Bob->L: Log transaction
    @enduml
        

### Use non-letters in participants

You can use quotes to define participants. And you can use the as keyword to give an alias to thoses participants.

    @startuml
    Alice -> "Bob()" : Hello
    "Bob()" -> "This is very\nlong" as Long
    ' You can also declare:
    ' "Bob()" -> Long as "This is very\nlong"
    Long --> "Bob()" : ok
    @enduml
        

### Message to Self

A participant can send a message to itself.
It is also possible to have multilines using \n.

    @startuml
    Alice->Alice: This is a signal to self.\nIt also demonstrates\nmultiline \ntext
    @enduml
        

### Change arrow style

You can change arrow style by several ways:
use \ or / instead of < or > to have only the bottom or top part of the arrow
repeat the arrow head (for example, >> or //) head to have a thin drawing
use -- instead of - to have a dotted arrow
add a final "o" at arrow head
use bidirectional arrow

    @startuml
    Bob -> Alice
    Bob ->> Alice
    Bob -\ Alice
    Bob \\- Alice
    Bob //-- Alice

    Bob ->o Alice
    Bob o\\-- Alice

    Bob <-> Alice
    Bob <->o Alice
    @enduml
            

### Change arrow color

You can change the color of individual arrows using the following notation:

    @startuml
    Bob -[#red]> Alice : hello
    Alice -[#0000FF]->Bob : ok
    @enduml
        

### Message sequence numbering

The keyword autonumber is used to automatically add number to messages.

    @startuml
    autonumber
    Bob -> Alice : Authentication Request
    Bob <- Alice : Authentication Response
    @enduml

You can specify a startnumber with autonumber 'start' , and also an increment with autonumber 'start' 'increment'

    @startuml
    autonumber
    Bob -> Alice : Authentication Request
    Bob <- Alice : Authentication Response

    autonumber 15
    Bob -> Alice : Another authentication Request
    Bob <- Alice : Another authentication Response

    autonumber 40 10
    Bob -> Alice : Yet another authentication Request
    Bob <- Alice : Yet another authentication Response

    @enduml

You can specify a format for your number by using between double-quote.
The formatting is done with the Java class DecimalFormat ('0' means digit, '#' means digit and zero if absent).

You can use some html tag in the format.

    @startuml
    autonumber "<b>[000]"
    Bob -> Alice : Authentication Request
    Bob <- Alice : Authentication Response

    autonumber 15 "<b>(<u>##</u>)"
    Bob -> Alice : Another authentication Request
    Bob <- Alice : Another authentication Response

    autonumber 40 10 "<font color=red><b>Message 0  "
    Bob -> Alice : Yet another authentication Request
    Bob <- Alice : Yet another authentication Response

    @enduml

### Title

The title keywords is used to put a title.

    @startuml

    title Simple communication example

    Alice -> Bob: Authentication Request
    Bob --> Alice: Authentication Response

    @enduml

### Legend the diagram

The legend and end legend are keywords is used to put a legend.
You can optionnaly specify to have left, right or center alignment for the legend.

    @startuml

    Alice -> Bob : Hello
    legend right
      Short
      legend
    endlegend

    @enduml

### Splitting diagrams

The newpage keyword is used to split a diagram into several images.
You can put a title for the new page just after the newpage keyword.

This is very handy with Word to print long diagram on several pages.

    @startuml

    Alice -> Bob : message 1
    Alice -> Bob : message 2

    newpage

    Alice -> Bob : message 3
    Alice -> Bob : message 4

    newpage A title for the\nlast page

    Alice -> Bob : message 5
    Alice -> Bob : message 6
    @enduml
        



### Grouping message

It is possible to group messages together using the following keywords:

* alt/else
* opt
* loop
* par
* break
* critical
* group, followed by a text to be displayed

It is possible a add a text that will be displayed into the header (except for group).

The end keyword is used to close the group.

Note that it is possible to nest groups.

    @startuml
    Alice -> Bob: Authentication Request

    alt successful case

        Bob -> Alice: Authentication Accepted
        
    else some kind of failure

        Bob -> Alice: Authentication Failure
        group My own label
            Alice -> Log : Log attack start
            loop 1000 times
                Alice -> Bob: DNS Attack
            end
            Alice -> Log : Log attack end
        end
        
    else Another type of failure

       Bob -> Alice: Please repeat
       
    end
    @enduml
        

### Notes on messages

It is possible to put notes on message using the note left or note right keywords just after the message .
You can have multilines note using the end note keywords.

    @startuml
    Alice->Bob : hello
    note left: this is a first note

    Bob->Alice : ok
    note right: this is another note

    Bob->Bob : I am thinking
    note left
        a note
        can also be defined
        on several lines
    end note
    @enduml
        

### Some other notes

It is also possible to place notes relative to participant with note left of , note right of or note over keywords.
It is possible to highlight a note by changing its background color.

You can also have multilines note using the end note keywords.

    @startuml
    participant Alice
    participant Bob
    note left of Alice #aqua
        This is displayed 
        left of Alice. 
    end note
     
    note right of Alice: This is displayed right of Alice.

    note over Alice: This is displayed over Alice.

    note over Alice, Bob #FFAAAA: This is displayed\n over Bob and Alice.

    note over Bob, Alice
        This is yet another
        example of
        a long note.
    end note
    @enduml
        

### Formatting using HTML

It is also possible to use few html-like tags like :

* `<b>` for bold text
* `<u>` or `<u:#AAAAAA>` or `<u:colorName>` for underline
* `<i>` for italic
* `<s>` or `<s:#AAAAAA>` or `<s:colorName>` for strike text
* `<w> or `<w:#AAAAAA>` or `<w:colorName>` for wave underline text
* `<color:#AAAAAA>` or `<color:colorName>`
* `<back:#AAAAAA>` or `<back:colorName>` for background color
* `<size:nn> to change font size
* `<img src="file">` or `<img:file>` : the file must be accessible by the filesystem
* `<img src="http://url">` or `<img:http://url>` : the URL must be available from the Internet

example

    @startuml
    participant Alice
    participant "The <b>Famous</b> Bob" as Bob

    Alice -> Bob : A <i>well formated</i> message
    note right of Alice 
     This is <back:cadetblue><size:18>displayed</size></back> 
     <u>left of</u> Alice. 
    end note
    note left of Bob 
     <u:red>This</u> is <color #118888>displayed</color> 
     <b><color purple>left of</color> <s:red>Alice</strike> Bob</b>. 
    end note
    note over Alice, Bob
     <w:#FF33FF>This is hosted</w> by <img sourceforge.jpg>
    end note
     
    @enduml
        

### Divider

If you want, you can split a diagram using == separator to divide your diagram into logical steps.

    @startuml

    == Initialisation ==

    Alice -> Bob: Authentication Request
    Bob --> Alice: Authentication Response

    == Repetition ==

    Alice -> Bob: Another authentication Request
    Alice <-- Bob: another authentication Response

    @enduml
        

### Reference

You can use reference in a diagram, using the keyword ref over.

    @startuml
    participant Alice
    actor Bob

    ref over Alice, Bob : init

    Alice -> Bob : hello

    ref over Bob
      This can be on
      several lines
    end ref
    @enduml
        

### Delay

You can use ... to indicate a delay in the diagram. And it is also possible to put a message with this delay.

    @startuml

    Alice -> Bob: Authentication Request
    ...
    Bob --> Alice: Authentication Response
    ...5 minutes latter...
    Bob --> Alice: Bye !

    @enduml
        

### Space

You can use ||| to indicate some spacing in the diagram. 
It is also possible to specify a number of pixel to be used.

    @startuml

    Alice -> Bob: message 1
    Bob --> Alice: ok
    |||
    Alice -> Bob: message 2
    Bob --> Alice: ok
    ||45||
    Alice -> Bob: message 3
    Bob --> Alice: ok

    @enduml
        

### Lifeline Activation and Destruction

The activate and deactivate are used to denote participant activation.
Once a participant is activated, its lifeline appears.

The activate and deactivate apply on the previous message.

The destroy denote the end of the lifeline of a participant.

    @startuml
    participant User

    User -> A: DoWork
    activate A

    A -> B: << createRequest >>
    activate B

    B -> C: DoWork
    activate C
    C --> B: WorkDone
    destroy C

    B --> A: RequestCreated
    deactivate B

    A -> User: Done
    deactivate A

    @enduml
        

Nested lifeline can be used, and it is possible to add a color on the lifeline.

    @startuml
    participant User

    User -> A: DoWork
    activate A #FFBBBB

    A -> A: Internal call
    activate A #DarkSalmon

    A -> B: << createRequest >>
    activate B

    B --> A: RequestCreated
    deactivate B
    deactivate A
    A -> User: Done
    deactivate A

    @enduml
        

### Participant creation

You can use the create keyword just before the first reception of a message to emphasize the fact that this message is actually creating this new object.

    @startuml
    Bob -> Alice : hello

    create Other
    Alice -> Other : new

    create control String
    Alice -> String
    note right : You can also put notes!

    Alice --> Bob : ok

    @enduml
        

### Incoming and outgoing messages

You can use incoming or outgoing arrows if you want to focus on a part of the diagram.
Use square brackets to denotate the left "[" or the right "]" side of the diagram.

    @startuml
    [-> A: DoWork

    activate A

    A -> A: Internal call
    activate A

    A ->] : << createRequest >>

    A<--] : RequestCreated
    deactivate A
    [<- A: Done
    deactivate A
    @enduml
        

### Stereotypes and Spots

It is possible to add stereotypes to participants using << and >>.
In the stereotype, you can add a spotted character in a colored circle using the syntax (X,color).

    @startuml

    participant "Famous Bob" as Bob << Generated >>
    participant Alice << (C,#ADD1B2) Testable >>

    Bob->Alice: First message

    @enduml
        
或者

    @startuml

    participant Bob << (C,#ADD1B2) >>
    participant Alice << (C,#ADD1B2) >>

    Bob->Alice: First message

    @enduml
        

### More information on titles

You can use some HTML tags in the title.

    @startuml

    title <u>Simple</u> communication example

    Alice -> Bob: Authentication Request
    Bob -> Alice: Authentication Response

    @enduml

You can add newline using \n in the title description.

    @startuml

    title <u>Simple</u> communication example\non several lines

    Alice -> Bob: Authentication Request
    Bob -> Alice: Authentication Response

    @enduml

You can also define title on several lines using title and end title keywords.

    @startuml

    title
     <u>Simple</u> communication example
     on <i>several</i> lines and using <font color=red>html</font>
     This is hosted by <img:sourceforge.jpg>
    end title

    Alice -> Bob: Authentication Request
    Bob -> Alice: Authentication Response

    @enduml

### Participants englober

It is possible to draw a box arround some participants, using box and end box commands.
You can add an optional title or a optional background color, after the box keyword.

    @startuml

    box "Internal Service" #LightBlue
        participant Bob
        participant Alice
    end box
    participant Other

    Bob -> Alice : hello
    Alice -> Other : hello

    @enduml

### Removing Footer

You can use the hide footbox keywords to remove the footer of the diagram.

    @startuml

    hide footbox
    title Footer removed

    Alice -> Bob: Authentication Request
    Bob --> Alice: Authentication Response

@enduml

Skinparam

You can use the skinparam command to change colors and fonts for the drawing.
You can use this command :

In the diagram definition, like any other commands,
In an included file,
In a configuration file, provided in the command line or the ANT task.

    @startuml
    skinparam backgroundColor #EEEBDC

    skinparam sequence {
        ArrowColor DeepSkyBlue
        ActorBorderColor DeepSkyBlue
        LifeLineBorderColor blue
        LifeLineBackgroundColor #A9DCDF
        
        ParticipantBorderColor DeepSkyBlue
        ParticipantBackgroundColor DodgerBlue
        ParticipantFontName Impact
        ParticipantFontSize 17
        ParticipantFontColor #A9DCDF
        
        ActorBackgroundColor aqua
        ActorFontColor DeepSkyBlue
        ActorFontSize 17
        ActorFontName Aapex
    }

    actor User
    participant "First Class" as A
    participant "Second Class" as B
    participant "Last Class" as C

    User -> A: DoWork
    activate A

    A -> B: Create Request
    activate B

    B -> C: DoWork
    activate C
    C --> B: WorkDone
    destroy C

    B --> A: Request Created
    deactivate B

    A --> User: Done
    deactivate A

    @enduml