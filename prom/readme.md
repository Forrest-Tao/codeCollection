**prom**
```bash
docker run -d --name=grafana -p 3000:3000 grafana/grafana-oss
```
- 可以通过http://127.0.0.1:9090/metrics查看。


**grafana**
```bash
# Create persistent volume for your data
docker volume create prometheus-data
# Start Prometheus container
docker run \
    -d \
    --name=prometheus \
    -p 9090:9090 \
    -v /path/to/prometheus.yml:/etc/prometheus/prometheus.yml \
    -v prometheus-data:/prometheus \
    prom/prometheus
```

- 启动成功后，使用浏览器打开 http://localhost:3000 
- 默认的登录账号是 “admin” / “admin”。