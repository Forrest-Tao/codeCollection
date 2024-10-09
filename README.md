```bash
.
├──  KeepAlive
│   └── main.go
├── LRU
│   ├── noneList
│   │   └── lru.go
│   └── withList
│       └── lru.go
├── README.md
├── SSE
│   ├── go.mod
│   ├── go.sum
│   └── main.go
├── atomic
│   └── main.go
├── binarySearch
│   └── main.go
├── bloomFilter
│   ├── bloomFilterService.go
│   ├── encrypt.go
│   ├── go.mod
│   ├── go.sum
│   ├── lua
│   │   ├── BatchGetBits.lua
│   │   └── SetBits.lua
│   └── redisClient.go
├── bufpool
│   └── bufpool.go
├── casbin
│   ├── go.mod
│   ├── go.sum
│   └── main.go
├── channel
│   ├── another.go
│   ├── go.mod
│   ├── main.go
│   ├── makefile
│   ├── mp.go
│   ├── muti_choice.go
│   ├── print1to9.go
│   ├── print1to9_test.go
│   ├── run_1k_task_with_100_goroutines.go
│   └── run_1k_task_with_100_goroutines_test.go
├── cobra
│   ├── go.mod
│   ├── go.sum
│   └── myapp
│       ├── LICENSE
│       ├── cmd
│       │   ├── hello.go
│       │   └── root.go
│       └── main.go
├── context
│   ├── go.mod
│   ├── go.sum
│   ├── main.go
│   └── share.go
├── dockerFile
│   ├── Dockerfile
│   ├── README.md
│   ├── go.mod
│   └── main.go
├── domain
│   └── main.go
├── errgroup
│   ├── go.mod
│   ├── go.sum
│   └── main.go
├── etcd
│   ├── go.mod
│   ├── go.sum
│   └── main.go
├── exception
│   ├── go.mod
│   └── main.go
├── geo
│   └── main.go
├── gorm
│   ├── go.mod
│   ├── go.sum
│   ├── main.go
│   └── repo.go
├── heap
│   └── main.go
├── httpServer
│   └── main.go
├── json
│   ├── CMakeLists.txt
│   ├── cmake-build-debug
│   │   ├── CMakeCache.txt
│   │   ├── CMakeFiles
│   │   │   ├── 3.29.3
│   │   │   │   ├── CMakeCCompiler.cmake
│   │   │   │   ├── CMakeDetermineCompilerABI_C.bin
│   │   │   │   ├── CMakeSystem.cmake
│   │   │   │   └── CompilerIdC
│   │   │   │       └── CMakeCCompilerId.c
│   │   │   ├── CMakeConfigureLog.yaml
│   │   │   ├── TargetDirectories.txt
│   │   │   ├── clion-Debug-log.txt
│   │   │   ├── clion-environment.txt
│   │   │   ├── cmake.check_cache
│   │   │   ├── json.dir
│   │   │   │   └── main.c.o
│   │   │   └── rules.ninja
│   │   ├── Testing
│   │   │   └── Temporary
│   │   │       └── LastTest.log
│   │   ├── build.ninja
│   │   ├── cmake_install.cmake
│   │   └── json
│   ├── json
│   ├── main.c
│   └── marshal
│       ├── go.mod
│       └── main.go
├── k8s
│   ├── code
│   │   ├── client-go
│   │   │   └── main.go
│   │   ├── json
│   │   │   └── json.go
│   │   ├── kubevirt
│   │   │   └── win
│   │   │       └── win10.yaml
│   │   ├── rand
│   │   │   ├── go.mod
│   │   │   ├── go.sum
│   │   │   ├── rand.go
│   │   │   └── rand_test.go
│   │   ├── scheduling
│   │   │   ├── limit-range.yaml
│   │   │   ├── nginx-with-resource.yaml
│   │   │   ├── nginx-with-tolerations.yaml
│   │   │   └── nginx-without-resource.yaml
│   │   ├── service
│   │   │   ├── Dockerfile
│   │   │   ├── README.md
│   │   │   ├── go.mod
│   │   │   ├── kind-config.yaml
│   │   │   ├── main.go
│   │   │   └── yaml
│   │   │       ├── deloymenyt.yaml
│   │   │       ├── ns.yaml
│   │   │       └── service.yaml
│   │   ├── storage
│   │   │   └── storage.yaml
│   │   └── workqueue
│   │       └── queue.go
│   └── yaml
│       ├── cronjob
│       │   └── cronjob.yaml
│       ├── deamonset
│       ├── deployment
│       │   └── my-app-deployment.yaml
│       ├── endpoints
│       │   └── endpoints.yaml
│       ├── kubeVirt
│       │   ├── demo
│       │   │   ├── pvc.yaml
│       │   │   └── vm2.yaml
│       │   └── win
│       │       ├── hd_pvc.yaml
│       │       ├── iso_hd.yaml
│       │       ├── iso_update.md
│       │       ├── sc.yaml
│       │       └── win_vm.yaml
│       ├── network-policy
│       │   └── NetworkPolicy.yaml
│       └── wordpress
│           └── wordpress-all.yaml
├── kafka
│   ├── go.mod
│   ├── go.sum
│   └── main.go
├── kubebuilder-demo
│   ├── go.sum
│   └── project
│       ├── Dockerfile
│       ├── Makefile
│       ├── PROJECT
│       ├── README.md
│       ├── api
│       │   └── v1
│       │       ├── cronjob_types.go
│       │       ├── cronjob_webhook.go
│       │       ├── cronjob_webhook_test.go
│       │       ├── groupversion_info.go
│       │       ├── webhook_suite_test.go
│       │       └── zz_generated.deepcopy.go
│       ├── cmd
│       │   └── main.go
│       ├── config
│       │   ├── certmanager
│       │   │   ├── certificate.yaml
│       │   │   ├── kustomization.yaml
│       │   │   └── kustomizeconfig.yaml
│       │   ├── crd
│       │   │   ├── bases
│       │   │   │   └── batch.tutorial.kubebuilder.io_cronjobs.yaml
│       │   │   ├── kustomization.yaml
│       │   │   ├── kustomizeconfig.yaml
│       │   │   └── patches
│       │   │       ├── cainjection_in_cronjobs.yaml
│       │   │       └── webhook_in_cronjobs.yaml
│       │   ├── default
│       │   │   ├── kustomization.yaml
│       │   │   ├── manager_metrics_patch.yaml
│       │   │   ├── manager_webhook_patch.yaml
│       │   │   ├── metrics_service.yaml
│       │   │   └── webhookcainjection_patch.yaml
│       │   ├── manager
│       │   │   ├── kustomization.yaml
│       │   │   └── manager.yaml
│       │   ├── network-policy
│       │   │   ├── allow-metrics-traffic.yaml
│       │   │   ├── allow-webhook-traffic.yaml
│       │   │   └── kustomization.yaml
│       │   ├── prometheus
│       │   │   ├── kustomization.yaml
│       │   │   └── monitor.yaml
│       │   ├── rbac
│       │   │   ├── cronjob_editor_role.yaml
│       │   │   ├── cronjob_viewer_role.yaml
│       │   │   ├── kustomization.yaml
│       │   │   ├── leader_election_role.yaml
│       │   │   ├── leader_election_role_binding.yaml
│       │   │   ├── metrics_auth_role.yaml
│       │   │   ├── metrics_auth_role_binding.yaml
│       │   │   ├── metrics_reader_role.yaml
│       │   │   ├── role.yaml
│       │   │   ├── role_binding.yaml
│       │   │   └── service_account.yaml
│       │   ├── samples
│       │   │   ├── batch_v1_cronjob.yaml
│       │   │   └── kustomization.yaml
│       │   └── webhook
│       │       ├── kustomization.yaml
│       │       ├── kustomizeconfig.yaml
│       │       ├── manifests.yaml
│       │       └── service.yaml
│       ├── go.mod
│       ├── go.sum
│       ├── hack
│       │   └── boilerplate.go.txt
│       ├── internal
│       │   └── controller
│       │       ├── cronjob_controller.go
│       │       ├── cronjob_controller_test.go
│       │       └── suite_test.go
│       └── test
│           ├── e2e
│           │   ├── e2e_suite_test.go
│           │   └── e2e_test.go
│           └── utils
│               └── utils.go
├── log
│   └── main.go
├── logs.txt
├── mongo
│   ├── README.md
│   ├── docker-compose.yaml
│   ├── go.mod
│   ├── go.sum
│   ├── main.go
│   └── mongo_test.go
├── pattern
│   ├── builderPattern
│   │   ├── builder.go
│   │   ├── builder_test.go
│   │   └── go.mod
│   ├── decoratorPattern
│   │   ├── README.md
│   │   └── decorator.go
│   ├── options
│   │   ├── go.mod
│   │   ├── options.go
│   │   └── options_test.go
│   └── singleton
│       ├── go.mod
│       ├── singleton.go
│       └── singleton_test.go
├── pool
│   ├── README.md
│   ├── buff_test.go
│   ├── go.mod
│   └── struct_test.go
├── pprof
│   ├── go.mod
│   ├── go.sum
│   └── main.go
├── prom
│   ├── client
│   │   ├── go.mod
│   │   └── main.go
│   ├── examplars
│   │   ├── go.mod
│   │   ├── go.sum
│   │   └── main.go
│   ├── prometheus.yml
│   └── readme.md
├── quickSelect
│   └── main.go
├── quickSort
│   └── main.go
├── rateLimit
│   ├── leaking_bucket
│   │   ├── app.log
│   │   ├── go.mod
│   │   ├── go.sum
│   │   └── main.go
│   ├── sliding_window
│   │   ├── go.mod
│   │   └── main.go
│   └── token_bucket
│       ├── go.mod
│       ├── go.sum
│       └── main.go
├── redisPubSub
│   ├── pubServer
│   │   ├── go.mod
│   │   ├── go.sum
│   │   └── main.go
│   └── subClient
│       ├── go.mod
│       ├── go.sum
│       └── main.go
├── safeMap
│   └── main.go
├── singleflight
│   ├── go.mod
│   ├── go.sum
│   ├── singleflight.go
│   └── singleflight_test.go
├── skipList
│   └── skipList.go
├── spinLock
│   └── spinLock.go
├── trace
│   ├── gin-demo
│   │   ├── go.mod
│   │   ├── go.sum
│   │   └── main.go
│   ├── jaeger
│   │   ├── go.mod
│   │   ├── go.sum
│   │   └── main.go
│   ├── readme.md
│   ├── redis
│   │   ├── go.mod
│   │   ├── go.sum
│   │   └── main.go
│   └── zap-trace
│       ├── go.mod
│       ├── go.sum
│       └── main.go
├── ws
│   ├── go.mod
│   ├── go.sum
│   ├── main.go
│   └── ws.html
├── xtimer
│   ├── main.go
│   └── ticker_test.go
└── zeroTS
    └── main.go

115 directories, 251 files


```


- bloomFilter https://hur.st/bloomfilter/?n=200000000&p=&m=838860800&k=3
- k8s cluster with kind  https://github.com/Forrest-Tao/codeCollection/commit/f247ef519f587e3cba9ef76753ad04cbc2773440