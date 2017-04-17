# NTFS


1、手动挂载

    hdiutil eject  /Volumes/disk_name
    sudo mount_ntfs -o rw,nobrowser /dev/disk1s1  /pwd/mount_point

2 、改动mount指令，自动mount分区

    sudo mv /sbin/mount_ntfs  /mount_ntfs.orig
    sudo touch /sbin/mount_ntfs
    sudo vim mount_ntfs
        #!/bin/sh
        /sbin/mount_ntfs.orig -o rw "$@"
    sudo chmod 755 /sbin/mount_ntfs

## cli

	sudo mount -t ntfs -o rw,auto,nobrowse /dev/disk3s1 ~/Desktop/external

参考文件

* <http://sourceforge.net/projects/ntfsfree/>
* <http://bbs.feng.com/read-htm-tid-1042827.html>
