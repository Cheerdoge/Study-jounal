# 概念
1. 镜像：运行应用所需的所有内容，代码、运行时、库文件、环境变量和配置文件
2. 容器：运行镜像的示例
3. 仓库：存储镜像
4. dockerfile：一系列用于自动构建 Docker 镜像的指令

# 命令

## run
创建并启动容器
`docker run [OPTIONS] IMAGE [COMMAND] [ARG...]`
- **`-d`**: 后台运行容器并返回容器 ID。
- **`-it`**: 交互式运行容器，分配一个伪终端。
- **`--name`**: 给容器指定一个名称。
- **`-p`**: 端口映射，格式为 `host_port:container_port`。
- **`-v`**: 挂载卷，格式为 `host_dir:container_dir`。
- **`--rm`**: 容器停止后自动删除容器。
- **`--env` 或 `-e`**: 设置环境变量。
- **`--network`**: 指定容器的网络模式。
- **`--restart`**: 容器的重启策略（如 `no`、`on-failure`、`always`、`unless-stopped`）。
- **`-u`**: 指定用户。

## pull
拉取镜像
`docker pull [OPTIONS] NAME[:TAG|@DIGEST]`
- **`--all-tags, -a`**: 下载指定镜像的所有标签。
- **`--disable-content-trust`**: 跳过镜像签名验证。

## exec
表示在容器内执行
`docker exec [OPTIONS] CONTAINER COMMAND [ARG...]`
- **`-d, --detach`**: 在后台运行命令。
- **`--detach-keys`**: 覆盖分离容器的键序列。
- **`-e, --env`**: 设置环境变量。
- **`--env-file`**: 从文件中读取环境变量。
- **`-i, --interactive`**: 保持标准输入打开。
- **`--privileged`**: 给这个命令额外的权限。
- **`--user, -u`**: 以指定用户的身份运行命令。
- **`--workdir, -w`**: 指定命令的工作目录。
- **`-t, --tty`**: 分配一个伪终端。
/bin/bash 表示打开bash shell

# image
* ls：列出有tag的镜像，加-a把没tag的也列出来

# load
从tar文件中获取镜像
`docker load [OPTIONS]`
- **`-i, --input`**: 指定输入文件的路径。
- **`-q, --quiet`**: 安静模式，减少输出信息。

## build
从 Dockerfile 构建 Docker 镜像
`docker build [OPTIONS] PATH | URL | -`
其中-表示从标准输入读取dockerfile
- **`-t, --tag`**: 为构建的镜像指定名称和标签。
- **`-f, --file`**: 指定 Dockerfile 的路径（默认是 `PATH` 下的 `Dockerfile`）。
- **`--build-arg`**: 设置构建参数。
- **`--no-cache`**: 不使用缓存层构建镜像。
- **`--rm`**: 构建成功后删除中间容器（默认开启）。
- **`--force-rm`**: 无论构建成功与否，一律删除中间容器。
- **`--pull`**: 始终尝试从注册表拉取最新的基础镜像。

## save
将一个或多个 Docker 镜像保存到一个 tar 归档文件
`docker save [OPTIONS] IMAGE.tar [IMAGE...]`
- **`-o, --output`**: 指定输出文件的路径。

## compose
默认先查找docker-compose.yml或.yaml文件然后执行

### up
按照文件中的设置创建并运行容器
`docker compose up [OPTIONS] [SERVICE...]`
- `-d, --detach`：以后台模式运行容器，类似于在 `docker run` 中使用 `-d` 选项。
- `--build`：在启动之前强制重新构建镜像，即使镜像已存在。
- `--no-build`：阻止在启动时构建镜像，即使镜像不存在也不构建。
- `--force-recreate`：强制重新创建容器，即使它们已经存在且内容未发生变化。
- `--no-recreate`：如果容器已经存在，则不重新创建它们（默认行为是如果配置文件变化则重新创建）。
- `--remove-orphans`：移除不再在 Compose 文件中定义的孤立容器。
- `-V, --renew-anon-volumes`：重新创建匿名卷（删除旧的卷并创建新的）。

### build
构建docker-compose文件
`docker compose up [OPTIONS] [SERVICE...]`