启动jaeger
```bash
docker run --rm --name jaeger \
-e COLLECTOR_ZIPKIN_HOST_PORT=:9411 \
-p 6831:6831/udp \
-p 6832:6832/udp \
-p 5778:5778 \
-p 16686:16686 \
-p 4317:4317 \
-p 4318:4318 \
-p 14250:14250 \
-p 14268:14268 \
-p 14269:14269 \
-p 9411:9411 \
jaegertracing/all-in-one:latest
```


* Port	Protocol	Component	Function
* 6831	UDP	agent	accept jaeger.thrift over Thrift-compact protocol (used by most SDKs)
* 6832	UDP	agent	accept jaeger.thrift over Thrift-binary protocol (used by Node.js SDK)
* 5775	UDP	agent	(deprecated) accept zipkin.thrift over compact Thrift protocol (used by legacy clients only)
* 5778	HTTP	agent	serve configs (sampling, etc.)
* 16686	HTTP	query	serve frontend
* 4317	HTTP	collector	accept OpenTelemetry Protocol (OTLP) over gRPC
* 4318	HTTP	collector	accept OpenTelemetry Protocol (OTLP) over HTTP
* 14268	HTTP	collector	accept jaeger.thrift directly from clients
* 14250	HTTP	collector	accept model.proto
* 9411	HTTP	collector	Zipkin compatible endpoint (optional)


**访问Jaeger UI http://localhost:16686** giti 