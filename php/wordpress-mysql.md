#wordpress 数据库详解#
**title:**wordpress 数据库详解
**tag:**wordpress,mysql    
**info:**wordpress的数据库内容    

>wp_commentmeta：存储评论的元数据    
wp_comments：存储评论    
wp_links：存储友情链接（Blogroll）    
wp_options：存储WordPress系统选项和插件、主题配置    
wp_postmeta：存储文章（包括页面、上传文件、修订）的元数据    
wp_posts：存储文章（包括页面、上传文件、修订）    
wp_terms：存储每个目录、标签    
wp_term_relationships：存储每个文章、链接和对应分类的关系    
wp_term_taxonomy：存储每个目录、标签所对应的分类    
wp_usermeta：存储用户的元数据    
wp_users：存储用户         
- - - - - - - 
###wp_commentmeta###
>meta_id：自增唯一ID     
comment_id：对应评论ID     
meta_key：键名     
meta_value：键值     

###wp_comments###
>comment_ID：自增唯一ID     
comment_post_ID：对应文章ID     
comment_author：评论者     
comment_author_email：评论者邮箱     
comment_author_url：评论者网址     
comment_author_IP：评论者IP     
comment_date：评论时间     
comment_date_gmt：评论时间（GMT+0时间）     
comment_content：评论正文     
comment_karma：未知     
comment_approved：评论是否被批准     
comment_agent：评论者的USER AGENT     
comment_type：评论类型(pingback/普通)     
comment_parent：父评论ID     
user_id：评论者用户ID（不一定存在）     


###wp_links###
>link_id：自增唯一ID     
link_url：链接URL     
link_name：链接标题     
link_image：链接图片     
link_target：链接打开方式     
link_description：链接描述     
link_visible：是否可见（Y/N）     
link_owner：添加者用户ID     
link_rating：评分等级     
link_updated：未知     
link_rel：XFN关系     
link_notes：XFN注释     
link_rss：链接RSS地址     


###wp_options###
>option_id：自增唯一ID     
blog_id：博客ID，用于多用户博客，默认0     
option_name：键名     
option_value：键值     
autoload：在WordPress载入时自动载入（yes/no）     


###wp_postmeta###
>meta_id：自增唯一ID     
post_id：对应文章ID     
meta_key：键名     
meta_value：键值     


###wp_posts###
>ID：自增唯一ID     
post_author：对应作者ID     
post_date：发布时间     
post_date_gmt：发布时间（GMT+0时间）     
post_content：正文     
post_title：标题     
post_excerpt：摘录     
post_status：文章状态（publish/auto-draft/inherit等）     
comment_status：评论状态（open/closed）     
ping_status：PING状态（open/closed）     
post_password：文章密码     
post_name：文章缩略名     
to_ping：未知     
pinged：已经PING过的链接     
post_modified：修改时间     
post_modified_gmt：修改时间（GMT+0时间）     
post_content_filtered：未知     
post_parent：父文章，主要用于PAGE     
guid：未知     
menu_order：排序ID     
post_type：文章类型（post/page等）     
post_mime_type：MIME类型     
comment_count：评论总数     


###wp_terms###
term_id：分类ID     
name：分类名     
slug：缩略名     
term_group：未知     


###wp_term_relationships###
object_id：对应文章ID/链接ID     
term_taxonomy_id：对应分类方法ID    
term_order：排序     


###wp_term_taxonomy###
>term_taxonomy_id：分类方法ID     
term_id：taxonomy：分类方法(category/post_tag)     
description：未知     
parent：所属父分类方法ID     
count：文章数统计     

###wp_usermeta###
>umeta_id：自增唯一ID     
user_id：对应用户ID     
meta_key：键名     
meta_value：键值     

###wp_users###
>ID：自增唯一ID	     
user_login：登录名     
user_pass：密码     
user_nicename：昵称     
user_email：Email     
user_url：网址     
user_registered：注册时间     
user_activation_key：激活码     
user_status：用户状态     
display_name：显示名称     
