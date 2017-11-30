# golang构建数据服务

使用了xorm实现了课程博客内容。                                                              
xorm确实是一个很省力的东西，用xorm即可实现课程博客上的dao内容。


## 实验结果

**添加数据**
post命令
```
curl -d "username=name&departname=d" http://localhost:8080/service/userinfo
{
  "UID": 0,
  "UserName": "name",
  "DepartName": "d",
  "CreateAt": "2017-11-30T13:40:11.4366116+08:00"
}
```

***查询数据***
```
$ curl http://localhost:8080/service/userinfo?userid=
[
  {
    "UID": 0,
    "UserName": "name",
    "DepartName": "d",
    "CreateAt": "2017-11-30T21:40:11+08:00"
  }
]
```


## ab测试
**使用xorm的ab测试**
```
.\ab -n 1000 -c 100 http://localhost:8080/
This is ApacheBench, Version 2.3 <$Revision: 1807734 $>
Copyright 1996 Adam Twiss, Zeus Technology Ltd, http://www.zeustech.net/
Licensed to The Apache Software Foundation, http://www.apache.org/

Benchmarking localhost (be patient)
Completed 100 requests
Completed 200 requests
Completed 300 requests
Completed 400 requests
Completed 500 requests
Completed 600 requests
Completed 700 requests
Completed 800 requests
Completed 900 requests
Completed 1000 requests
Finished 1000 requests


Server Software:
Server Hostname:        localhost
Server Port:            8080

Document Path:          /
Document Length:        19 bytes

Concurrency Level:      100
Time taken for tests:   0.779 seconds
Complete requests:      1000
Failed requests:        0
Non-2xx responses:      1000
Total transferred:      176000 bytes
HTML transferred:       19000 bytes
Requests per second:    1284.35 [#/sec] (mean)
Time per request:       77.860 [ms] (mean)
Time per request:       0.779 [ms] (mean, across all concurrent requests)
Transfer rate:          220.75 [Kbytes/sec] received

Connection Times (ms)
              min  mean[+/-sd] median   max
Connect:        0    0   0.3      0       1
Processing:    33   73  13.8     74      99
Waiting:       21   71  12.7     73      95
Total:         34   73  13.8     74      99

Percentage of the requests served within a certain time (ms)
  50%     74
  66%     78
  75%     79
  80%     80
  90%     96
  95%     97
  98%     98
  99%     99
 100%     99 (longest request)
```


**使用database/sql的ab测试**
```
 .\ab -n 1000 -c 100 http://localhost:8080/
This is ApacheBench, Version 2.3 <$Revision: 1807734 $>
Copyright 1996 Adam Twiss, Zeus Technology Ltd, http://www.zeustech.net/
Licensed to The Apache Software Foundation, http://www.apache.org/

Benchmarking localhost (be patient)
Completed 100 requests
Completed 200 requests
Completed 300 requests
Completed 400 requests
Completed 500 requests
Completed 600 requests
Completed 700 requests
Completed 800 requests
Completed 900 requests
Completed 1000 requests
Finished 1000 requests


Server Software:
Server Hostname:        localhost
Server Port:            8080

Document Path:          /
Document Length:        19 bytes

Concurrency Level:      100
Time taken for tests:   0.679 seconds
Complete requests:      1000
Failed requests:        0
Non-2xx responses:      1000
Total transferred:      176000 bytes
HTML transferred:       19000 bytes
Requests per second:    1472.72 [#/sec] (mean)
Time per request:       67.902 [ms] (mean)
Time per request:       0.679 [ms] (mean, across all concurrent requests)
Transfer rate:          253.12 [Kbytes/sec] received

Connection Times (ms)
              min  mean[+/-sd] median   max
Connect:        0    0   0.3      0       2
Processing:     6   65  17.1     71     118
Waiting:        6   63  17.8     69     117
Total:          6   65  17.2     71     118

Percentage of the requests served within a certain time (ms)
  50%     71
  66%     73
  75%     76
  80%     78
  90%     87
  95%     93
  98%     95
  99%     96
 100%    118 (longest request)
```

　　通过比较可以得出，两者之间性能是有差异的，使用database/sql要比使用xorm要快一些，这在实验的toys中是不明显的，但如果放在企业级的大程序中，影响却会很大。
