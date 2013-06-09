# nginx php mp4 上传 

## php.ini

	; 开启上传
	file_uploads = On

	; 上传暂存目录
	;upload_tmp_dir =

	; 上传文件最大值
	upload_max_filesize = 120M

	; 一次性上传文件数量
	max_file_uploads = 20

## nginx 

	# 上传文件最大值
	client_max_body_size 120M;  
	# 设置临时目录
	# client_body_temp_path /home/www/nginx_temp;

## ffmpeg

~~~{.php}
	$from_name = ; // 文件地址
	$timg_name = ; // 缩略图
	$to_name   = ; // 转化mp4路径
    $w = 600 ;
    $h = 350 ;

	exec("ffmpeg -i $from_name -y -f image2 -ss 5 -s $w*$h -vframes 1 $timg_name");
	exec("ffmpeg -i $from_name -ar 44100 -strict -2 -y $to_name");
~~~

