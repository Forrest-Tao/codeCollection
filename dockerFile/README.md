```bash
#构建镜像
$  docker build -t myapp .

#确定镜像构建成功
$ docker images | grep myapp
myapp                                                  latest          034d09a4e57e   About a minute ago   9.38MB

#使用 CMD中的参数
$ docker run --rm myapp
in init function fileName: inDockerfile
in main function
inDockerfile
done

#覆盖CMD中的参数
$ docker run --rm myapp -fileName="zxczxca"
in init function fileName: zxczxca
in main function
zxczxca
done
```