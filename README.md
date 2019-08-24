# 使用说明

### 功能描述：

#### 1.文件上传：

接口：user/upload 

例： 用Postman 上传一个文件 

成功返回 json：{“"stauts":  "OK","message": "上传成功"}

失败返回 json: {   "status" :"Fail",   "message" : "上传失败"}

文件过大返回 {   "status" : "Fail",   "message" : "文件过大"}

上传过 返回 {   "stauts":  "Fail",   "message": "已上传"}

#### 2.文件下载：

接口： user/download/:name  name (想要下载文件的全名称)

例： user/download/1.txt

如果有返回 文件下载

没有 就返回 json {"status":  "Fail","message": "文件不存在"}

#### 3.浏览文件

接口： scan

例：scan

返回json ：[{"filename":"1.txt","size":0},{"filename":"asd","size":0}]

#### 4.分享文件

接口 share

参数 name

例：name 文件下载接口 整个 uri 

然后生成一个二维码扫描 emmm  调用的别人的api

#### 5.注册账号

接口 register

参数 name passwd

成功 返回 user

#### 6.加密分享下载

接口 get/:name 

参数 secret 

例： get/1.txt?secret=1234

secret 正确 开始下载

不正确 就下不了~!

#### 7.一次性上传

接口 transmission

**参照1.文件上传**

返回一个 url 点击即可一次性下载 也可以分享出去

#### 8.一次性下载

接口 down

那就直接下载了 **参照7一次性上传**

#### 9.登录账号

接口 login 

参数 name passwd

例 login?name=admin&passwd=admin

成功 返回 {"message": "登录成功"}

失败 {"message": "请输入正确的账号或密码"}

未输入参数 {"message": "请输入账号或密码"}

#### 10.加密分享

接口 user/share

参数 filename secret（可选）不输入secret 会自动生成 一个

你要加密的文件 可以 私有的 文件 也可以是 公有的文件

例：user/share?filename=1.txt&secret=2134sdh

返回 json {"status":"OK","secret":secret,"Download": c.Request.Host+"/"+"get/"+filename}

在Download 点击 链接 后 看 **6.加密分享下载**





**注：好像成功设置了文件下载限速，大概在60kb/s**

关于断点续传：几个请求头 Range 如何 做到第一次请求后，暂停拿到range 这里卡住了