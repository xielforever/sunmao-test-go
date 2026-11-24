# sunmao-test-go

用于 SunMao CI/CD 发布测试的 Go 后端测试项目。

## 接口

- `GET /healthz`：健康检查，返回 `200 OK` JSON
- `GET /` / `GET /api/info`：返回服务名、版本号（`APP_VERSION`）、主机名与时间

## 本地运行与测试

```bash
go test ./...
go run .
```

## 发布测试产物

- `Dockerfile`：多阶段构建轻量 Alpine 运行镜像（默认端口 `8080`）
- `deploy/k8s.yaml`：Kubernetes `Namespace` / `Deployment` / `Service` 清单（`sunmao-samples/sunmao-test-go`）
- `sunmao.yaml`：SunMao 流水线配置（支持 `server` / `container` / `k8s` / `all` 部署目标）
