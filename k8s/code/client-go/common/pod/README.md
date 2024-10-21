```bash
kubeconfig->rest.Config->clientset->具体的client(比如Corev1)
->具体的资源对象->restclient->http.Client->HTTP请求->response
```