# lag-api
Simulates slow API responses


### Docker image
`docker pull raittes/lag-api`


### Static Mode
Define rules in YML file for routes, responses and delays.

**Example:** `lag-api -static-rules static-example.yml`

`docker run -v $PWD/static-example.yml:/static.yml -p 8888:8888 -d lag-api -static-rules static.yml`

Test routes: 
```
curl -w "\n%{time_total}\n" localhost:8888/hello
curl -w "\n%{time_total}\n" localhost:8888/test
curl -w "\n%{time_total}\n" localhost:8888/slow
```


### Proxy Mode
Forward the requests to another endpoint and include a delay in response.

**Example:** `lag-api -proxy http://httpbin.org -lag 500ms`

`docker run -p 8888:8888 -d lag-api -proxy http://httpbin.org -lag 500ms`

Test:
```
curl -w "\n%{time_total}\n" localhost:8888/ip
curl -w "\n%{time_total}\n" localhost:8888/user-agent
curl -w "\n%{time_total}\n" localhost:8888/headers
```
