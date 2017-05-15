# docker on mac

使用 : docker， docker-machine, docker-compose, kitematic

```sh
brew install docker-machine 
brew cask install Caskroom/cask/kitematic
```

创建一台 docker machine

```sh
docker-machine create –driver virtualbox default
–driver virtualbox
本机使用的话 使用virtualbox驱动
driver 还可以是amazonec2, azure, digitalocean, exoscale, generic, google, openstack, rackspace, softlayer, virtualbox, vmwarefusion, vmwarevcloudair, vmwarevsphere 等 配置对应的秘钥和参数后 可以直接在云服务中创建设备
default
docker-machine 中的机器名 之后的管理命令均使用该名称
```

设置环境变量

```sh
eval “$(docker-machine env default)”
–driver virtualbox
default
之前创建的machine的名称
```

使用docker客户端管理docker, 设置环境变量，使用docker管理docker设备