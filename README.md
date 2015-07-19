# Caesar
-------------

Caesar是用go语言开发的消息队列。

> **主要特点:**
> - 支持消息队列类型划分，消息储存划分为内存、数据库及文件。

> - 多消息队列同时运行。

> - 支持后台注册用户，通过命令行客户端对消息队列进行控制：增加、删除、启动、停止。

> - 支持Restful API接口。

## Caesar安装

> 1、安装golang运行环境

> 2、安装mysql数据库，新建mysql数据库CAESAR，导入建表语句——Caesar/sql/caesar.sql

> 3、导入caesar及依赖包

  
  > - go get -u github.com/jin06/Caesar
  
  > - go get -u github.com/Sirupsen/logrus
  
  > - go get -u github.com/ant0ine/go-json-rest/rest
  
  > - go get -u github.com/jin06/Caesar/control
  
### Caesar运行环境 mysql及golang版本
![Alt text](/image/version.png)
  
## Caesar运行及配置

###1、Caesar服务器程启动：

  > cd $GOPATH/github.com/jin06/Caesar/src
  
  > go run CaesarServer.go
   
###2、Caesar服务器的配置文件

  配置文件地址 Caesar/config
  
  （1）db.yaml 设置mysql数据库
  
              >  username: beijing  //数据库用户名
              >  password: beijing  //数据库密码
              >  dbname: caesar     //数据库名
              >  address: 127.0.0.1:3306   //数据库地址及端口
                
  （2）msgserver.yaml 设置消息服务器
  
              >  listenport: "1991"   //地址默认本机，设置监听端口
                
  （3）server.yaml  设置控制服务端口
  
              >  rpcAddress: 127.0.0.1:1212   //该地址为Caesar客户端连接地址
                
###3、Caesar客户端程序启动：

    cd $GOPATH/github.com/jin06/Caesar/src
    go run CaesarClient.go
    
###4、Caesar客户端的配置文件 Caesar/client/config/cient.yaml

                 serveraddress: 127.0.0.1:1212    //连接服务端地址 对应2中的配置文件server.yaml的rpcAddress
                 localaddress: 127.0.0.1:2213     //客户端启动端口
                 
###5、CaesarServer启动参数说明

| 命令     | 参数 | 说明   |      使用例子              |
| :------- | ----: | :---: |:---:                    |
| address | [ip:port] |  服务器启动地址及监听端口  |  go run CaesarServer.go -address 127.0.0.1:1212   |
| version    |    |  服务端版本说明   |      go run CaesarServer.go -version            |
| runtime     | [int] |  服务器运行级别  |        未实现               |

###6、CaesarClient启动参数说明

| 命令     | 参数 | 说明   |      使用例子              |
| :------- | ----: | :---: |:---:                    |
| local | [ip:port] |  客户端启动地址及监听端口  |  go run CaesarClient.go -local 127.0.0.1:1212   |
| server    |  [ip:port]  |  服务端的地址   |      go run CaesarClient.go -server 127.0.0.1:9999             |
| version     |  |  客户端程序版本说明  |        go run CaesarClient.go -version               |

###7、CaesarServer启动图示
![Alt text](/image/ServerRun.png)
###8、CaesarClient启动图示
![Alt text](/image/clientrun.png)

## Caesar使用

成功运行Caesar服务器，启动Caesar客户端程序，启动服务器端程序。启动后在命令行输入login登录服务器，数据库默认有两个用户admin及jinlong。其中admin为管理员，可以增加、删除用户，普通用户jinlong只有对自己的消息队列的操作权限。

     用户密码
     
      > admin用户密码 123456
      > jinlong用户密码为 123456 
####命令列表  

| 命令     | 参数 | 说明   |      使用例子              | 
| :------- | ----: | :---: |:---:                    |
| login |  |  用户登录操作  |  login后，根据提示输入用户名密码   | 
| info    |  用户登录信息  |     |      info             |    
| exit     |  |  用户退出  |        exit               |
| myqueue     |  |  用户的所有队列  |        myqueue               |
| create user     |  |  创建用户  |       管理员才有执行权限               |
| delete user     | [userid] |  所有用户  |      管理员才有执行权限, delete user 123 //123为用户id   |
| create mq     |   |  创建message queue  |      create mq               |
| delete mq     | [mqid] |  删除message queue  |      delete mq 123 124 111 //删除队列号为123，124，111的队列    |
| start mq     |  [mqid] |  运行队列  |       start 123 111 //运行123，111队列               |
| stop mq     | [mqid] |  停止队列  |       stop 123 111 //停止123，111队列                |

#####1、用户操作
用户操作必须为管理员才可以。使用admin用户登录，create user为创建用户。delete为删除用户：

  > - 使用login命令登录
  
  > - 创建用户test
  
  > - 删除用户test
  
  > - 使用users命令查看所有用户
  
![Alt text](/image/usercurd.png)

#####2、队列操作

  > - 使用普通用户jinlong登录
  
  > - 创建队列mq
  
  > - 启动队列mq
  
  > - 停止队列mq
  
  > - 删除队列mq
  
  > - 使用myqueue命令查看所有用户
  
  ![Alt text](/image/mqcrud.png)

## Caesar接口----Resutful API

| 访问地址     | 说明 | 其他   |     
| :------- | ----: | :---: |
| http://127.0.0.1/send_msg/32002 | post message |  其中32002为队列号，如果往其他队列post请更改该号码即可。此地址为不可靠消息队列，即队列储存在内存  | 
| http://127.0.0.1/receive_msg/32002 | get message |  其中32002为队列号，如果从其他队列get请更改该号码即可。此地址为不可靠消息队列，即队列储存在内存   |  
| http://127.0.0.1/send/32002 | post message |  其中32002为队列号，如果往其他队列post请更改该号码即可。此地址为可靠消息队列，即队列缓存到数据库  | 
| http://127.0.0.1/receive/32002 | get message |  其中32002为队列号，如果从其他队列get请更改该号码即可。此地址为可靠消息队列，即队列缓存到数据库 | 
| http://127.0.0.1/r/testmq/32002 | get message |  测试mq32002是否运行，其中32002为队列号，如果想测试其他队列请更改该号码即可。 |

> - 可靠消息队列Post方法 

> curl -i -H 'Content-Type: application/json' \ -d '{"Value":"test message"}' http://127.0.0.1:1991/send_msg/32002

> - 不可靠消息队列Post方法  

> curl -i -H 'Content-Type: application/json' \ -d '{"Value":"test message"}' http://127.0.0.1:1991/send/32002

> - 可靠消息队列Get方法

> curl -i http://127.0.0.1:1991/receive_msg/32002

> - 不可靠消息队列Get方法

> curl -i http://127.0.0.1:1991/receive/32002

## API返回的http报文及代码含义

| KEY     | VALUE | 说明   |     
| :------- | ----: | :---: |
| 1016 | post success |  报文正确送达  |  
| 1011 | server receive, but not save to db |  信息未保存到数据库，一般为可靠消息读列使用  | 
| 1010 | mq not running |  消息队列未启动  | 1011
| 1011 | no message in db | 队列中已经没有消息可以读了 |
| 1020 | mq is running | 测试接口，测试mq是否运行 |

## API测试的shell脚本地址 

其中 receive_get.sh及 send_post.sh为 不可靠消息队列的测试脚本，receive_msg_get.sh及 send_msg_post.sh为可靠消息队列的测试脚本。

> github.com/jin06/Caesar/test/receive_get.sh

> github.com/jin06/Caesar/test/receive_msg_get.sh

> github.com/jin06/Caesar/test/send_msg_post.sh

> github.com/jin06/Caesar/test/send_post.sh

## API使用全过程图示

>  - 运行Server
>  - 运行Client，登录，启动一个队列
>  - post一个消息
>  - get一个消息

 ![Alt text](/image/restfulapi.png)
 
 ![Alt text](/image/restfulapi2.png)
 
# 鸣谢
  鸣谢以下项目
  > - github.com/Sirupsen/logrus
  
      使用该项目进行了日志的美化
  
  > - github.com/ant0ine/go-json-rest/rest
  
      使用该项目简化了restful api接口的编写
      
  > - github.com/jin06/Caesar/control
  
      使用该项目对配置文件进行读取操作
