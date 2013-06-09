# openssl 

## 建立 CA

建立 CA 目录结构

按照 OpenSSL 的默认配置建立 CA ，需要在文件系统中建立相应的目录结构。相关的配置内容一般位于 /usr/ssl/openssl.cnf 内，详情可参见 config (1) 。在终端中使用如下命令建立目录结构：

    $ mkdir -p ./demoCA/{private,newcerts}
    $ touch ./demoCA/index.txt
    $ echo 01 > ./demoCA/serial

产生的目录结构如下：

    .
    `-- demoCA/
        |-- index.txt
        |-- newcerts/
        |-- private/
        `-- serial

生成 CA 证书的 RSA 密钥对

首先，我们要为 CA 建立 RSA 密钥对。打开终端，使用如下命令生成 RSA 密钥对：

    $ openssl genrsa -des3 -out ./demoCA/private/cakey.pem 2048

参数解释

    genrsa

        用于生成 RSA 密钥对的 OpenSSL 命令。

    -des3

        使用 3-DES 对称加密算法加密密钥对，该参数需要用户在密钥生成过程中输入一个口令用于加密。今后使用该密钥对时，需要输入相应的口令。如果不加该选项，则不对密钥进行加密。

    -out ./demoCA/private/cakey.pem

        令生成的密钥对保存到文件 ./demoCA/private/cakey.pem 。

    2048

        RSA 模数位数，在一定程度上表征了密钥强度。

该命令输出如下，用户应输入自己的密钥口令并确认：

    Generating RSA private key, 2048 bit long modulus
    ................................................+++
    .........................+++
    e is 65537 (0x10001)
    Enter pass phrase for ./demoCA/private/cakey.pem:<enter your pass-phrase>
    Verifying - Enter pass phrase for ./demoCA/private/cakey.pem:<re-enter your pass-phrase>

## 生成 CA 证书请求

为了获取一个 CA 根证书，我们需要先制作一份证书请求。先前生成的 CA 密钥对被用于对证书请求签名。

    $ openssl req -new -days 365 -key ./demoCA/private/cakey.pem -out careq.pem

参数解释

    req

        用于生成证书请求的 OpenSSL 命令。

    -new

        生成一个新的证书请求。该参数将令 OpenSSL 在证书请求生成过程中要求用户填写一些相应的字段。

    -days 365

        从生成之时算起，证书时效为 365 天。

    -key ./demoCA/private/cakey.pem

        指定 ./demoCA/private/cakey.pem 为证书所使用的密钥对文件。

    -out careq.pem

        令生成的证书请求保存到文件 careq.pem 。

该命令将提示用户输入密钥口令并填写证书相关信息字段，输出如下：

    Enter pass phrase for ./demoCA/private/cakey.pem:<enter you pass-phrase>
    You are about to be asked to enter information that will be incorporated
    into your certificate request.
    What you are about to enter is what is called a Distinguished Name or a DN.
    There are quite a few fields but you can leave some blank
    For some fields there will be a default value,
    If you enter '.', the field will be left blank.
    -----
    Country Name (2 letter code) [AU]:CN
    State or Province Name (full name) [Some-State]:ZJ
    Locality Name (eg, city) []:HZ
    Organization Name (eg, company) [Internet Widgits Pty Ltd]:Some Ltd. Corp.
    Organizational Unit Name (eg, section) []:Some Unit
    Common Name (eg, YOUR name) []:Someone
    Email Address []:some@email.com

    Please enter the following 'extra' attributes
    to be sent with your certificate request
    A challenge password []:
    An optional company name []:

## 对 CA 证书请求进行签名

在实际应用中，用户可以通过向知名 CA 递交证书请求来申请证书。但是在这里，我们需要建立的是一个根 CA ，只能由我们自己来对证书请求进行签名。所以我们让 OpenSSL 使用证书请求中附带的密钥对对该请求进行签名，也就是所谓的“ self sign ”：

    $ openssl ca -selfsign -in careq.pem -out cacert.pem

参数解释

    ca

        用于执行 CA 相关操作的 OpenSSL 命令。

    -selfsign

        使用对证书请求进行签名的密钥对来签发证书。

    -in careq.pem

        指定 careq.pem 为证书请求文件。

    -out ./demoCA/cacert.pem

        指定 ./demoCA/cacert.pem 为输出的证书。

该命令要求用户输入密钥口令并输出相关证书信息，请求用户确认：

    Using configuration from /usr/lib/ssl/openssl.cnf
    Enter pass phrase for ./demoCA/private/cakey.pem:<enter your pass-phrase>
    Check that the request matches the signature
    Signature ok
    Certificate Details:
            Serial Number: 2 (0x2)
            Validity
                Not Before: Jan 16 13:05:09 2008 GMT
                Not After : Jan 15 13:05:09 2009 GMT
            Subject:
                countryName = CN
                stateOrProvinceName = ZJ
                organizationName = Some Ltd. Corp.
                organizationalUnitName = Some Unit
                commonName = Someone
                emailAddress = some@email.com
            X509v3 extensions:
                X509v3 Basic Constraints:
                    CA:FALSE
                Netscape Comment:
                    OpenSSL Generated Certificate
                X509v3 Subject Key Identifier:
                    75:F5:3C:CC:C1:5E:6D:C3:8B:46:A8:08:E6:EA:29:E8:22:7E:70:03
                X509v3 Authority Key Identifier:
                    keyid:75:F5:3C:CC:C1:5E:6D:C3:8B:46:A8:08:E6:EA:29:E8:22:7E:70:03

    Certificate is to be certified until Jan 15 13:05:09 2009 GMT (365 days)
    Sign the certificate? [y/n]:y


    1 out of 1 certificate requests certified, commit? [y/n]y
    Write out database with 1 new entries
    Data Base Updated

一步完成 CA 证书请求生成及签名

以上两个步骤可以合二为一。利用 ca 命令的 -x509 参数，通过以下命令同时完成证书请求生成和签名从而生成 CA 根证书：

    $ openssl req -new -x509 -days 365 -key ./demoCA/private/cakey.pem -out ./demoCA/cacert.pem

参数解释

    req

        用于生成证书请求的 OpenSSL 命令。

    -new

        生成一个新的证书请求。该参数将令 OpenSSL 在证书请求生成过程中要求用户填写一些相应的字段。

    -x509

        生成一份 X.509 证书。

    -days 365

        从生成之时算起，证书时效为 365 天。

    -key ./demoCA/private/cakey.pem

        指定 cakey.pem 为证书所使用的密钥对文件。

    -out ./demoCA/cacert.pem

        令生成的证书保存到文件 ./demoCA/cacert.pem 。

该命令输出如下，用户应输入相应的字段：

    Enter pass phrase for ./demoCA/private/cakey.pem:
    You are about to be asked to enter information that will be incorporated
    into your certificate request.
    What you are about to enter is what is called a Distinguished Name or a DN.
    There are quite a few fields but you can leave some blank
    For some fields there will be a default value,
    If you enter '.', the field will be left blank.
    -----
    Country Name (2 letter code) [AU]:CN
    State or Province Name (full name) [Some-State]:ZJ
    Locality Name (eg, city) []:HZ
    Organization Name (eg, company) [Internet Widgits Pty Ltd]:Some Ltd. Corp.
    Organizational Unit Name (eg, section) []:Some Unit
    Common Name (eg, YOUR name) []:Someone
    Email Address []:some@email.com

至此，我们便已成功建立了一个私有根 CA 。在这个过程中，我们获得了一份 CA 密钥对文件 ./demoCA/private/cakey.pem 以及一份由此密钥对签名的 CA 根证书文件 ./demoCA/cacert.pem ，得到的 CA 目录结构如下：

	.
	|-- careq.pem
	`-- demoCA/
		|-- cacert.pem
		|-- index.txt
		|-- index.txt.attr
		|-- index.txt.old
		|-- newcerts/
		|   `-- 01.pem
		|-- private/
		|   `-- cakey.pem
		|-- serial
		`-- serial.old

注：如果在 CA 建立过程中跳过证书请求生成的步骤，则不会产生 careq.pem 文件。
签发证书

下面我们就可以利用建立起来的 CA 进行证书签发了。
生成用户证书 RSA 密钥对

参照 CA 的 RSA 密钥对生成过程，使用如下命令生成新的密钥对：

    $ openssl genrsa -des3 -out userkey.pem
    Generating RSA private key, 512 bit long modulus
    ....++++++++++++
    ...++++++++++++
    e is 65537 (0x10001)
    Enter pass phrase for userkey.pem:<enter your pass-phrase>
    Verifying - Enter pass phrase for userkey.pem:<re-enter your pass-phrase>

## 生成用户证书请求

参照 CA 的证书请求生成过程，使用如下命令生成新的证书请求：

    $ openssl req -new -days 365 -key userkey.pem -out userreq.pem
    Enter pass phrase for userkey.pem:<enter your pass-phrase>
    You are about to be asked to enter information that will be incorporated
    into your certificate request.
    What you are about to enter is what is called a Distinguished Name or a DN.
    There are quite a few fields but you can leave some blank
    For some fields there will be a default value,
    If you enter '.', the field will be left blank.
    -----
    Country Name (2 letter code) [AU]:CN
    State or Province Name (full name) [Some-State]:ZJ
    Locality Name (eg, city) []:HZ
    Organization Name (eg, company) [Internet Widgits Pty Ltd]:Some Ltd. Corp.
    Organizational Unit Name (eg, section) []:Some Other Unit
    Common Name (eg, YOUR name) []:Another
    Email Address []:another@email.com

    Please enter the following 'extra' attributes
    to be sent with your certificate request
    A challenge password []:
    An optional company name []:

## 签发用户证书

现在，我们可以用先前建立的 CA 来对用户的证书请求进行签名来为用户签发证书了。使用如下命令：

$ openssl ca -in userreq.pem -out usercert.pem
参数解释

ca

    用于执行 CA 相关操作的 OpenSSL 命令。

-in userreq.pem

    指定用户证书请求文件为 userreq.pem 。

-out usercert.pem

    指定输出的用户证书文件为 usercert.pem 。

该命令要求用户输入密钥口令并输出相关证书信息，请求用户确认：

	Using configuration from /usr/lib/ssl/openssl.cnf
	Enter pass phrase for ./demoCA/private/cakey.pem:<enter your pass-phrase>
	Check that the request matches the signature
	Signature ok
	Certificate Details:
			Serial Number: 2 (0x2)
			Validity
				Not Before: Jan 16 14:50:22 2008 GMT
				Not After : Jan 15 14:50:22 2009 GMT
			Subject:
				countryName               = CN
				stateOrProvinceName       = ZJ
				organizationName          = Some Ltd. Corp.
				organizationalUnitName    = Some Other Unit
				commonName                = Another
				emailAddress              = another@email.com
			X509v3 extensions:
				X509v3 Basic Constraints:
					CA:FALSE
				Netscape Comment:
					OpenSSL Generated Certificate
				X509v3 Subject Key Identifier:
					97:E7:8E:84:B1:45:27:83:94:A0:DC:24:79:7B:83:97:99:0B:36:A9
				X509v3 Authority Key Identifier:
					keyid:D9:87:12:94:B2:20:C7:22:AB:D4:D5:DF:33:DB:84:F3:B0:4A:EC:A2

	Certificate is to be certified until Jan 15 14:50:22 2009 GMT (365 days)
	Sign the certificate? [y/n]:y


	1 out of 1 certificate requests certified, commit? [y/n]y
	Write out database with 1 new entries
	Data Base Updated

## 步骤

至此，我们便完成了 CA 的建立及用户证书签发的全部工作。不妨把所有 shell 命令放到一起纵览一下：

    # 建立 CA 目录结构
    mkdir -p ./demoCA/{private,newcerts}
    touch ./demoCA/index.txt
    echo 01 > ./demoCA/serial

    # 生成 CA 的 RSA 密钥对
    openssl genrsa -des3 -out ./demoCA/private/cakey.pem 2048

    # 生成 CA 证书请求
    openssl req -new -days 365 -key ./demoCA/private/cakey.pem -out careq.pem

    # 自签发 CA 证书
    openssl ca -selfsign -in careq.pem -out ./demoCA/cacert.pem

    # 以上两步可以合二为一
    openssl req -new -x509 -days 365 -key ./demoCA/private/cakey.pem -out ./demoCA/cacert.pem

    # 生成用户的 RSA 密钥对
    openssl genrsa -des3 -out userkey.pem

    # 生成用户证书请求
    openssl req -new -days 365 -key userkey.pem -out userreq.pem

    # 使用 CA 签发用户证书
    openssl ca -in userreq.pem -out usercert.pem

了解了这些基础步骤之后，就可以通过脚本甚至 makefile 的方式来将这些工作自动化。 CA.pl 和 CA.sh 便是对 OpenSSL 的 CA 相关功能的简单封装，在 Debian 系统中，安装了 OpenSSL 后，可以在 /usr/lib/ssl/misc/ 目录下找到这两个文件。而 makefile 的解决方案则可以参考这里。

## 参考
[openssl](http://rhythm-zju.blog.163.com/blog/static/310042008015115718637/)  
[生成](http://rhythm-zju.blog.163.com/blog/static/310042008015115718637/)
