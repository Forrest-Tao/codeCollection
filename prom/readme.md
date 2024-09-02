**prom**
```bash
docker volume create prometheus-data
docker run \
    -d \
    --name=prometheus \
    -p 9090:9090 \
    -v ./prometheus.yml:/etc/prometheus/prometheus.yml \ #指定自己的prometheus.yaml文件
    -v prometheus-data:/prometheus \
    prom/prometheus
```
- 可以通过 http://127.0.0.1:9090 查看。


**grafana**
```bash
docker run -d --name=grafana -p 3000:3000 grafana/grafana-oss
```
- 启动成功后，使用浏览器打开 http://localhost:3000 
- 默认的登录账号是 “admin” / “admin”。