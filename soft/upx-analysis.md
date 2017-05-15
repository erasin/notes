分析 upx 壳
====================

> <http://www.chinapyg.com/thread-76768-1-1.html>

UPX是公认的一款优秀的压缩壳。无论是兼容性还是压缩率都是非常让人满意的。当然，想用它来保护您的软件大概是不可能的了。今天闲来无事，就顺手分析了一下UPX的壳，就当做学习一下这种壳该如何弄。下面是一点小小的心得，就分享给大家了。

加壳软件为 UPX309w，目标文件为XP32下的记事本（多次躺枪的软件{:soso_e113:}），保护方式为默认保护，使用的逆向工具为PYG版的OD，以上，下面正式开始。

首先说一下我的想法，UPX的压缩原理是利用了重复字符的存储方法。比如说，要存储10个a，如aaaaaaaaaa，只需要保存一个a以及它连续出现了10次的状况就可以了。即使不是相邻的，如 aabbbaaccc,那么就要保存a，然后a重复一次，接着在偏移为3的地方重复aa一次。这样来保存文件，将有效的减少文件的大小，这就是UPX的压缩理论。当然实际上，它使用的方法要更为巧妙。

众所周知，UPX压缩后的文件有两个区段，UPX0,UPX1,其中UPX0是一个空的区段，是原程序用来保存代码的地方。而UPX1则是保存被压缩后的程序数据的地方，这是一个二进制的数据流。UPX的解码代码便是使用该数据来进行解码而得到原本的程序。好了，知道了这些便开始分析吧。

首先是UPX的入口

```
01014260 >  60              PUSHAD                                      ; 保存寄存器
01014261    BE 00000101     MOV     ESI, notepad_.01010000              ; 取UPX1段地址, ESI <- UPX0
01014266    8DBE 0010FFFF   LEA     EDI, DWORD PTR DS:[ESI+0xFFFF1000]  ; 取UPX0段地址, EDI <- UPX1
0101426C    57              PUSH    EDI                                 ; 保存UPX0段到堆栈中
0101426D    83CD FF         OR      EBP, 0xFFFFFFFF                     ; SET EBP = -1，原地址与目的地址的偏移
01014270    EB 10           JMP     SHORT notepad_.01014282             ; 跳转到解码代码
```

进行了一下代码的初始化以后，便直接跳向了解码代码段

```
01014276    90              NOP                                         ; EDI目的地址 EDX源地址 ECX 复制长度 EBX二进制数据流
01014277    90              NOP                                         ; UPX1段为二进制数据流，根据二进制来进行流程的控制与解码
01014278    8A06            MOV     AL, BYTE PTR DS:[ESI]               ; 从ESI中取1字节数据到EDI去
0101427A    46              INC     ESI
0101427B    8807            MOV     BYTE PTR DS:[EDI], AL
0101427D    47              INC     EDI
0101427E >  01DB            ADD     EBX, EBX                            ; 移动EBX的高位到CF标志中去
01014280    75 07           JNZ     SHORT notepad_.01014289
01014282    8B1E            MOV     EBX, DWORD PTR DS:[ESI]             ; EBX <- DWORD:UPX1[i]
01014284    83EE FC         SUB     ESI, -0x4                           ; i++，此时CF = 1，目的是令下面从组EBX时,保证最低位肯定为1，使得流程控制中使用掉EBX的所有位
01014287    11DB            ADC     EBX, EBX                            ; 重组EBX,CF标志的数值排到EBX的低位，同时EBX的高位排到CF高位
01014289  ^ 72 ED           JB      SHORT notepad_.01014278             ; 高位为1，跳转，从ESI取1byte到EDI去 （CF标志）
0101428B    B8 01000000     MOV     EAX, 0x1                            ; 高位不为1，不跳转，给EAX1
01014290    01DB            ADD     EBX, EBX                            ; EBX高位进CF寄存器
01014292    75 07           JNZ     SHORT notepad_.0101429B
01014294    8B1E            MOV     EBX, DWORD PTR DS:[ESI]             ; 取下一组数据
01014296    83EE FC         SUB     ESI, -0x4
01014299    11DB            ADC     EBX, EBX
0101429B    11C0            ADC     EAX, EAX                            ; CF标志位左移进EAX
0101429D    01DB            ADD     EBX, EBX                            ; EBX高位进CF标志位中
0101429F  ^ 73 EF           JNB     SHORT notepad_.01014290             ; CF为0，跳转，继续移动EBX高位到EAX中，CF为1，不跳转，进行下一步
010142A1    75 09           JNZ     SHORT notepad_.010142AC
010142A3    8B1E            MOV     EBX, DWORD PTR DS:[ESI]
010142A5    83EE FC         SUB     ESI, -0x4
010142A8    11DB            ADC     EBX, EBX
010142AA  ^ 73 E4           JNB     SHORT notepad_.01014290
010142AC    31C9            XOR     ECX, ECX                            ; 初始化ECX = 0 (计算得到CopyLen)
010142AE    83E8 03         SUB     EAX, 0x3
010142B1    72 0D           JB      SHORT notepad_.010142C0             ; 如果EAX >= 11,不跳转，（流程分派）
010142B3    C1E0 08         SHL     EAX, 0x8                            ; EAX <<= 8
010142B6    8A06            MOV     AL, BYTE PTR DS:[ESI]               ; 从ESI中取出byte
010142B8    46              INC     ESI
010142B9    83F0 FF         XOR     EAX, 0xFFFFFFFF                     ; XOR 求出源数据与目标数据间的差距（偏移）
010142BC    74 74           JE      SHORT <notepad_.End Decode>         ; 得到的结果为0时退出解码阶段
010142BE    89C5            MOV     EBP, EAX                            ;  写入偏移大小
010142C0    01DB            ADD     EBX, EBX                            ; 取得EBX高位到CF标志
010142C2    75 07           JNZ     SHORT notepad_.010142CB
010142C4    8B1E            MOV     EBX, DWORD PTR DS:[ESI]
010142C6    83EE FC         SUB     ESI, -0x4
010142C9    11DB            ADC     EBX, EBX
010142CB    11C9            ADC     ECX, ECX                            ; CF标志左移入ECX中(CopyLen)
010142CD    01DB            ADD     EBX, EBX                            ; 取得EBX高位到CF标志
010142CF    75 07           JNZ     SHORT notepad_.010142D8
010142D1    8B1E            MOV     EBX, DWORD PTR DS:[ESI]
010142D3    83EE FC         SUB     ESI, -0x4
010142D6    11DB            ADC     EBX, EBX
010142D8    11C9            ADC     ECX, ECX                            ; CF标志左移入ECX中(CopyLen)，到此总共移动了2位EBX到ECX中,后面进入流程判断
010142DA    75 20           JNZ     SHORT notepad_.010142FC
010142DC    41              INC     ECX                                 ; 若ECX为0，则进入该流程，需要额外计算得到ECX的大小并且令ECX = 1
010142DD    01DB            ADD     EBX, EBX                            ; 将EBX高位进CF标志
010142DF    75 07           JNZ     SHORT notepad_.010142E8
010142E1    8B1E            MOV     EBX, DWORD PTR DS:[ESI]
010142E3    83EE FC         SUB     ESI, -0x4
010142E6    11DB            ADC     EBX, EBX
010142E8    11C9            ADC     ECX, ECX                            ; CF标志左移入ECX低位
010142EA    01DB            ADD     EBX, EBX                            ; EBX高位入CF标志
010142EC  ^ 73 EF           JNB     SHORT notepad_.010142DD             ; CF为0,跳转,继续移动EBX高位到ECX中，CF为1,不跳转,进入下一个流程
010142EE    75 09           JNZ     SHORT notepad_.010142F9
010142F0    8B1E            MOV     EBX, DWORD PTR DS:[ESI]
010142F2    83EE FC         SUB     ESI, -0x4
010142F5    11DB            ADC     EBX, EBX
010142F7  ^ 73 E4           JNB     SHORT notepad_.010142DD
010142F9    83C1 02         ADD     ECX, 0x2                            ; ECX += 2
010142FC    81FD 00F3FFFF   CMP     EBP, -0xD00                         ;  比较源数据地址与目标数据地址的偏移
01014302    83D1 01         ADC     ECX, 0x1
01014305    8D142F          LEA     EDX, DWORD PTR DS:[EDI+EBP]         ; 取源数据地址
01014308    83FD FC         CMP     EBP, -0x4                           ;  如果偏移小于4，则需要一个一个字节的复制，否则可以4个字节4个字节的复制(效率问题)
0101430B    76 0F           JBE     SHORT notepad_.0101431C
0101430D    8A02            MOV     AL, BYTE PTR DS:[EDX]               ; memcpy((byte*)EDI, (byte*)EDX, (DWORD)ECX)
0101430F    42              INC     EDX                                 ; 从EDX中复制ECX长度数据到EDI去
01014310    8807            MOV     BYTE PTR DS:[EDI], AL
01014312    47              INC     EDI
01014313    49              DEC     ECX
01014314  ^ 75 F7           JNZ     SHORT notepad_.0101430D
01014316  ^ E9 63FFFFFF     JMP     <notepad_.Loop>
```

这里，通过对二进制代码流UPX1的处理，解码出原程序代码存放到UPX0中去，汇编代码可能会有点乱，下面给出伪代码：

```
byte* UPX0,*UPX1;                                //UPX1数据要解码到UPX0中
DWORD iCopyLen(ECX),Dst(EDI) = 0,Src(EDX) = 0;        //复制Src到Est中
       //复制数据后，Dst与Src自动加上iCopyLen的大小。
DWORD pPoint(ESI) = 0;
DWORD Data(EBX);                                //从UPX1中读取的二进制数据流
       // 每当Data取完8个二进制数据时（肯定可以，见①），便会Data = UPX1[pPoint],pPoing += 4;为了方便，后面不再写出
       // 每当Data & 1000 0000B 判断高位数据是否为1，便自动左移一位 即 Data <<= 1,后面也不写出
DWORD ptrNow = -1, ptrNew = 0;                        //源地址与目标地址的偏移
jmp @Loop;
while(true)
{
 @newChar:
    memcpy(&UPX0[Dst], &UPX1[pPoint], 1);
 @Loop:
    if(Data & 1000 0000B)
        jmp @newChar;
    ptrNew = 1;
 @ptrRun:
    ptrNew </ Data;                                //这里假定 </ 表示从Data高位左移到ptrNew的低位中去
    if(Data & 0000 0000B)
        jmp @ptrRun;                                //高位为0时，不断移位
    ptrNew -= 3;
    if(ptrNew >= 0)
    {
        ptrNew <<= 4;
        ptrNew += byte:UPX1[pPoint];
        ptrNew ^= -1;
        if(ptrNew == 0)
            break;
        ptrNow = ptrNew;
    }
    iCopyLen = 0;
    iCopyLen </ Data;
    iCopyLen </ Data;                                //从Data中左移2位到iCopyLen中,0~3
    if(iCopyLen == 0)                                //证明复制的长度大于3，需要从新计算长度
    {
 @CopyLenRun:
       iCopyLen = 1;
       iCopyLen </ Data;
       if(Data & 0000 0000B)                        //高位为0时，继续移动高位到iCopyLen中
           jmp @CopyLenRun;
       iCopyLen += 2;
    }
    if(ptrNow > 0xFFFFF300)                        //作用？
       iCopyLen += 1;
    else iCopyLen += 2;
    Src = Dst + ptrNow;
    memcpy(Dst, Src, iCopyLen);
}

PS:
①EBX的更新的汇编代码
ADD     EBX, EBX                            ; 移动EBX的高位到CF标志中去
JNZ     NEXT
MOV     EBX, DWORD PTR DS:[ESI]             ; EBX <- DWORD:UPX1[i]
SUB     ESI, -0x4                           ; i++，此时CF = 1，目的是令下面从组EBX时,保证最低位肯定为1，使得流程控制中使
用掉EBX的所有位
ADC     EBX, EBX                            ; 重组EBX,CF标志的数值排到EBX的低位(保证低位为1)，同时EBX的高位排到CF高位
NEXT:
```

好了，做完解码的工作以后，便是进行call与jmp的重定位，首先是解密程序内的call与jmp

```
01014332 >  5E              POP     ESI                                 ; 解密已结束，此时进行从定位
01014333    89F7            MOV     EDI, ESI                            ; 恢复UPX0的位置， ESI，EDI <- UPX0
01014335    B9 32010000     MOV     ECX, 0x132                          ;  共有132个函数被变形,修程序内的jmp与call
0101433A    8A07            MOV     AL, BYTE PTR DS:[EDI]               ; 从代码段中寻找E8,E9的位置
0101433C    47              INC     EDI                                 ; E8 <- CALL ???
0101433D    2C E8           SUB     AL, 0xE8                            ; E9 <- jmp
0101433F    3C 01           CMP     AL, 0x1
01014341  ^ 77 F7           JA      SHORT notepad_.0101433A
01014343    803F 01         CMP     BYTE PTR DS:[EDI], 0x1              ; 需要修复的CALL或jmp均被解压为 E9 01 或 E8 01,故需要判断后面是否为01
01014346  ^ 75 F2           JNZ     SHORT notepad_.0101433A
01014348    8B07            MOV     EAX, DWORD PTR DS:[EDI]             ; 取4个字节
0101434A    8A5F 04         MOV     BL, BYTE PTR DS:[EDI+0x4]           ;  取下一个判断位置
0101434D    66:C1E8 08      SHR     AX, 0x8                             ; AX >>= 8,丢弃低位01
01014351    C1C0 10         ROL     EAX, 0x10                           ; EAX <<= 10 (ROL),高位变低位
01014354    86C4            XCHG    AH, AL                              ; XCHG AH,AL -> 总的来说就是把01去掉，EAX整个翻转，求得目标位置距离UPX0起始位置的偏移
01014356    29F8            SUB     EAX, EDI                            ; 减去当前call的位置
01014358    80EB E8         SUB     BL, 0xE8                            ;   ->下一轮的判断数据
0101435B    01F0            ADD     EAX, ESI                            ; 加上UPX0起始位置，求得目标位置距离当前位置的偏移
0101435D    8907            MOV     DWORD PTR DS:[EDI], EAX             ; 写回
0101435F    83C7 05         ADD     EDI, 0x5
01014362    88D8            MOV     AL, BL
01014364  ^ E2 D9           LOOPD   SHORT notepad_.0101433F             ; 循环
```

因为UPX对call与jmp也是存在一定的变形的，故需要进行还原
然后进行API的重定位。

```
01014366    8DBE 00200100   LEA     EDI, DWORD PTR DS:[ESI+0x12000]     ; 从定位API函数
0101436C    8B07            MOV     EAX, DWORD PTR DS:[EDI]
0101436E    09C0            OR      EAX, EAX
01014370    74 3C           JE      SHORT notepad_.010143AE
01014372    8B5F 04         MOV     EBX, DWORD PTR DS:[EDI+0x4]         ; 取源程序导入表的偏移
01014375    8D8430 14AE0100 LEA     EAX, DWORD PTR DS:[EAX+ESI+0x1AE14] ; 定位到从定位表中,取dll名
0101437C    01F3            ADD     EBX, ESI                            ;  + 源程序基址得到导入表地址
0101437E    50              PUSH    EAX
0101437F    83C7 08         ADD     EDI, 0x8
01014382    FF96 DCAE0100   CALL    NEAR DWORD PTR DS:[ESI+0x1AEDC]     ; LoadLibraryA,加载dll
01014388    95              XCHG    EAX, EBP                            ; 保存dll位置
01014389    8A07            MOV     AL, BYTE PTR DS:[EDI]
0101438B    47              INC     EDI
0101438C    08C0            OR      AL, AL
0101438E  ^ 74 DC           JE      SHORT notepad_.0101436C
01014390    89F9            MOV     ECX, EDI                            ; 取函数名
01014392    57              PUSH    EDI
01014393    48              DEC     EAX
01014394    F2:AE           REPNE   SCAS BYTE PTR ES:[EDI]
01014396    55              PUSH    EBP
01014397    FF96 E0AE0100   CALL    NEAR DWORD PTR DS:[ESI+0x1AEE0]     ; GetProcAddress，从定位API函数
0101439D    09C0            OR      EAX, EAX
0101439F    74 07           JE      SHORT notepad_.010143A8
010143A1    8903            MOV     DWORD PTR DS:[EBX], EAX             ; 写入原程序的导入表中
010143A3    83C3 04         ADD     EBX, 0x4
010143A6  ^ EB E1           JMP     SHORT notepad_.01014389
```

非常简单得调用LoadLibraryA与GetProcAddress来进行API重定位，这点相信大家可以很好的处理。  
最后便是返回源程序OEP

```
010143A8    FF96 F0AE0100   CALL    NEAR DWORD PTR DS:[ESI+0x1AEF0]     ; ExitProcess 发生错误退出程序
010143AE    8BAE E4AE0100   MOV     EBP, DWORD PTR DS:[ESI+0x1AEE4]
010143B4    8DBE 00F0FFFF   LEA     EDI, DWORD PTR DS:[ESI-0x1000]
010143BA    BB 00100000     MOV     EBX, 0x1000
010143BF    50              PUSH    EAX
010143C0    54              PUSH    ESP
010143C1    6A 04           PUSH    0x4
010143C3    53              PUSH    EBX
010143C4    57              PUSH    EDI
010143C5    FFD5            CALL    NEAR EBP                            ; VirtualProtect 新建内存
010143C7    8D87 FF010000   LEA     EAX, DWORD PTR DS:[EDI+0x1FF]
010143CD    8020 7F         AND     BYTE PTR DS:[EAX], 0x7F
010143D0    8060 28 7F      AND     BYTE PTR DS:[EAX+0x28], 0x7F
010143D4    58              POP     EAX
010143D5    50              PUSH    EAX
010143D6    54              PUSH    ESP
010143D7    50              PUSH    EAX
010143D8    53              PUSH    EBX
010143D9    57              PUSH    EDI
010143DA    FFD5            CALL    NEAR EBP                            ; VirtualProtect
010143DC    58              POP     EAX
010143DD    61              POPAD
010143DE    8D4424 80       LEA     EAX, DWORD PTR SS:[ESP-0x80]
010143E2    6A 00           PUSH    0x0
010143E4    39C4            CMP     ESP, EAX
010143E6  ^ 75 FA           JNZ     SHORT notepad_.010143E2
010143E8    83EC 80         SUB     ESP, -0x80                          ; 清空堆栈中残留的数据
010143EB  - E9 AD2FFFFF     JMP     notepad_.0100739D                   ; 跳向源程序OEP
```

UPX虽然简单，但却包含了加壳最基本的事情，包括代码变形，代码压缩，API重定位等思想

