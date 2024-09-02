package main

import (
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"math/rand"
	"net/http"
	"time"
)

// Counter 是一个计数器，只能递增。
var (
	httpRequestsTotal = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "http_requests_total",
			Help: "Total number of HTTP requests made.",
		},
		[]string{"method", "endpoint"},
	)
)

// Gauge 是一个可增可减的指标，通常用于表示当前状态的数值。
var (
	inProgressRequests = prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "in_progress_requests",
		Help: "Number of requests currently being handled.",
	})
)

// Histogram 是一个直方图，用于记录数据的分布，例如请求的时长。
var (
	requestDurationHistogram = prometheus.NewHistogram(prometheus.HistogramOpts{
		Name:    "http_request_duration_seconds",
		Help:    "Histogram of latencies for HTTP requests.",
		Buckets: prometheus.DefBuckets, // 使用默认的桶配置
	})
)

// Summary 是一个摘要，用于记录请求的统计数据，类似于 Histogram，但它提供了百分位数的功能。
var (
	requestDurationSummary = prometheus.NewSummary(prometheus.SummaryOpts{
		Name:       "http_request_duration_summary",
		Help:       "Summary of latencies for HTTP requests.",
		Objectives: map[float64]float64{0.5: 0.05, 0.9: 0.01, 0.99: 0.001},
	})
)

func init() {
	// 将所有的指标注册到全局默认注册表中
	prometheus.MustRegister(httpRequestsTotal)
	prometheus.MustRegister(inProgressRequests)
	prometheus.MustRegister(requestDurationHistogram)
	prometheus.MustRegister(requestDurationSummary)
}

// 模拟处理请求的函数
func handleRequest(c *gin.Context) {
	startTime := time.Now()

	// 增加进行中的请求计数
	inProgressRequests.Inc()
	defer inProgressRequests.Dec()

	// 模拟处理时间
	time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)

	// 记录请求的总数，标签包含方法和端点信息
	httpRequestsTotal.WithLabelValues(c.Request.Method, c.Request.URL.Path).Inc()

	// 计算请求处理时长
	duration := time.Since(startTime).Seconds()

	// 将时长记录到 Histogram 和 Summary 中
	requestDurationHistogram.Observe(duration)
	requestDurationSummary.Observe(duration)

	// 响应客户端
	c.String(http.StatusOK, "Hello, Prometheus!")
}

func main() {
	// 设置时区为上海
	loc, err := time.LoadLocation("Asia/Shanghai")
	if err != nil {
		panic(err)
	}
	time.Local = loc

	r := gin.Default()

	// 创建 HTTP 路由，并设置处理函数
	r.GET("/api", handleRequest)

	// 将 Prometheus 的指标处理器挂载到 /metrics 路径
	r.GET("/metrics", gin.WrapH(promhttp.Handler()))

	// 启动 HTTP 服务器
	r.Run(":8080")
}
