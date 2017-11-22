# cloudgo-io
一个简单的web应用程序，功能包括：
- 支持静态文件服务
- 支持简单js访问
- 提交表单，并输出一个表格
- 对/unknow给出开发中的提示，返回码501

因上次使用了martini框架，本次作业也使用martini框架。

## Usage
```
./main
or 
./main -p port
```

## Example
**静态文件服务：**
```
./main
[martini] listening on :8080 (development)
[martini] Started GET / for [::1]:50484
[martini] [Static] Serving /index.html
[martini] Completed 200 OK in 53.0366ms
```

```
curl -v http://localhost:8080
详细信息: GET http://localhost:8080/ with 0-byte payload
详细信息: received 433-byte response of content type text/html; charset=utf-8


StatusCode        : 200
StatusDescription : OK
Content           : <html>
                    <head>
                      <link rel="stylesheet" href="css/main.css"/>
                      <script src="http://code.jquery.com/jquery-latest.js"></script>
                      <script src="js/hello.js"></script>
                    </head>
                    <body>
                      <img src="i...
RawContent        : HTTP/1.1 200 OK
                    Accept-Ranges: bytes
                    Content-Length: 433
                    Content-Type: text/html; charset=utf-8
                    Date: Tue, 21 Nov 2017 18:21:35 GMT
                    Last-Modified: Tue, 21 Nov 2017 16:57:15 GMT

                    <html>
                    <head>...
Forms             : {}
Headers           : {[Accept-Ranges, bytes], [Content-Length, 433], [Content-Type, text/html; charset=utf-8], [Date, Tu
                    e, 21 Nov 2017 18:21:35 GMT]...}
Images            : {@{innerHTML=; innerText=; outerHTML=<IMG src="images/cng.png" width=48 height=48>; outerText=; tag
                    Name=IMG; src=images/cng.png; width=48; height=48}}
InputFields       : {}
Links             : {}
ParsedHtml        : mshtml.HTMLDocumentClass
RawContentLength  : 433
```
                                            
![](https://github.com/Suenaa/golang/blob/master/cloudgo-io/README/static.PNG)


**简单JS访问：**
```
./main
[martini] listening on :8080 (development)
[martini] Started GET /json for [::1]:50582
[martini] Completed 200 OK in 514µs
```

```
curl -v http://localhost:8080/json
详细信息: GET http://localhost:8080/json with 0-byte payload
详细信息: received 52-byte response of content type application/json; charset=UTF-8


StatusCode        : 200
StatusDescription : OK
Content           : {
                      "id": "101010",
                      "content": "Hello from Go!"
                    }

RawContent        : HTTP/1.1 200 OK
                    Content-Length: 52
                    Content-Type: application/json; charset=UTF-8
                    Date: Tue, 21 Nov 2017 18:24:07 GMT

                    {
                      "id": "101010",
                      "content": "Hello from Go!"
                    }

Forms             : {}
Headers           : {[Content-Length, 52], [Content-Type, application/json; charset=UTF-8], [Date, Tue, 21 Nov 2017 18:
                    24:07 GMT]}
Images            : {}
InputFields       : {}
Links             : {}
ParsedHtml        : mshtml.HTMLDocumentClass
RawContentLength  : 52
```                                                       
![](https://github.com/Suenaa/golang/blob/master/cloudgo-io/README/json.PNG)                              


**提交表单，并输出一个表格：**                                                       

访问 http://localhost:8080/login:     

![](https://github.com/Suenaa/golang/blob/master/cloudgo-io/README/login.PNG)                           
输入信息并且点击登录：                                                                                 

![](https://github.com/Suenaa/golang/blob/master/cloudgo-io/README/loginsuc.PNG)                                       
                                          
                                                                             
**对/unknow给出开发中的提示，返回码501：**                                                
访问 http://localhost:8080/unknow                                
![](https://github.com/Suenaa/golang/blob/master/cloudgo-io/README/unknow.PNG)
