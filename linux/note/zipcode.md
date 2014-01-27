# zip解压乱码

用 7z 在 LANG=C 的环境下解压 zip 文件：

	LANG=C 7z x zipfile.zip

测试文件名的编码转换，查看是否有乱码 (PS: 只需要关注转换后的文件名是否有乱码，不需要关注目录名)：

	convmv -f cp936 -t utf8 -r -- *

如果没有乱码，就进行实际的转换：

	convmv -f cp936 -t utf8 -r --notest -- *
	
如果还需要转换某个文件的编码：

	iconv -f cp936 -t utf8 -o output.txt input.txt