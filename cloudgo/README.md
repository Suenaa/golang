# cloudgo

## curl测试
运行cloudgo：                              
```./main -p 9090
[martini] listening on :9090 (development)
[martini] Started GET / for [::1]:55107
[martini] Completed 200 OK in 970.5µs```

```
curl -v http://localhost:9090
详细信息: GET http://localhost:9090/ with 0-byte payload
详细信息: received 13-byte response of content type text/plain; charset=utf-8


StatusCode        : 200
StatusDescription : OK
Content           : Hello Golang!
RawContent        : HTTP/1.1 200 OK
                    Content-Length: 13
                    Content-Type: text/plain; charset=utf-8
                    Date: Wed, 15 Nov 2017 09:21:18 GMT

                    Hello Golang!
Forms             : {}
Headers           : {[Content-Length, 13], [Content-Type, text/plain; charset=utf-8], [Date, Wed, 15 Nov 2017 09:21:18
                    GMT]}
Images            : {}
InputFields       : {}
Links             : {}
ParsedHtml        : mshtml.HTMLDocumentClass
RawContentLength  : 13

```

## ab测试
```
.\ab -n 1000 -c 100 http://localhost:9090/
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
Server Port:            9090

Document Path:          /
Document Length:        13 bytes

Concurrency Level:      100
Time taken for tests:   0.708 seconds
Complete requests:      1000
Failed requests:        0
Total transferred:      130000 bytes
HTML transferred:       13000 bytes
Requests per second:    1412.98 [#/sec] (mean)
Time per request:       70.772 [ms] (mean)
Time per request:       0.708 [ms] (mean, across all concurrent requests)
Transfer rate:          179.38 [Kbytes/sec] received

Connection Times (ms)
              min  mean[+/-sd] median   max
Connect:        0    0   0.4      0       1
Processing:     2   66  11.9     69     111
Waiting:        1   66  12.4     69     111
Total:          2   66  11.9     70     111

Percentage of the requests served within a certain time (ms)
  50%     70
  66%     71
  75%     72
  80%     72
  90%     74
  95%     74
  98%     75
  99%     76
 100%    111 (longest request)
 ```