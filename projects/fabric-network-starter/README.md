# Smart LMS Fabric Network Starter

Project này là cấu hình khởi tạo Hyperledger Fabric network dựa trên Hyperledger Bevel.

Mục tiêu mặc định:

- Fabric `2.5.4`
- Consensus `raft`
- 1 orderer organization: `supplychain`
- 1 peer organization: `manufacturer`
- 1 channel: `smartlmschannel`
- 1 chaincode placeholder: `smartlms`

## Cấu trúc

```text
projects/fabric-network-starter/
  network.yaml              # Cấu hình Bevel chính
  env.example               # Các giá trị cần thay trước khi deploy
  chaincode/                # Smart contract Go mẫu
  scripts/
    build-image.ps1         # Build Bevel build image local
    validate.ps1            # Validate network.yaml bằng Bevel build image
    deploy.ps1              # Chạy Bevel deploy từ Docker
```

## Việc cần chuẩn bị

Bevel không tạo sẵn hạ tầng cloud cho bạn. Trước khi deploy, cần có:

- Kubernetes cluster đã truy cập được bằng `kubectl`
- Hashicorp Vault đã init/unseal và có root token
- Git repository dùng cho GitOps, user/token có quyền push
- DNS/ingress phù hợp với `external_url_suffix`
- Docker chạy được trên máy local

Nếu dùng Kubernetes local như `minikube`, hãy đổi `cloud_provider`, `k8s.context`, `k8s.config_file`, DNS và ingress theo môi trường thật của bạn. File hiện tại để dạng starter, chưa nên chạy production khi chưa thay secret.

## Cách cấu hình

1. Copy kubeconfig dùng để deploy vào:

```powershell
Copy-Item $env:USERPROFILE\.kube\config projects\fabric-network-starter\config
```

2. Mở `network.yaml` và thay toàn bộ placeholder:

- `CHANGE_ME_DOMAIN`
- `CHANGE_ME_K8S_CONTEXT`
- `CHANGE_ME_K8S_CONFIG`
- `CHANGE_ME_VAULT_ADDR`
- `CHANGE_ME_VAULT_ROOT_TOKEN`
- `CHANGE_ME_GIT_USERNAME`
- `CHANGE_ME_GIT_TOKEN`
- `CHANGE_ME_GIT_EMAIL`

3. Build Bevel build image nếu máy chưa có:

```powershell
.\projects\fabric-network-starter\scripts\build-image.ps1
```

4. Validate file:

```powershell
.\projects\fabric-network-starter\scripts\validate.ps1
```

5. Deploy network:

```powershell
.\projects\fabric-network-starter\scripts\deploy.ps1
```

Script deploy mount repo hiện tại vào `/home/bevel` và mount project này vào `/home/bevel/build`, đúng với `run.sh` của Bevel.

## Chaincode Go

Chaincode mẫu nằm ở `chaincode/`. Bevel sẽ pull source từ Git theo cấu hình trong `network.yaml`:

```yaml
repository:
  url: "github.com/CHANGE_ME_GIT_USERNAME/bevel.git"
  branch: main
  path: "projects/fabric-network-starter/chaincode"
```

Vì vậy sau khi chỉnh chaincode, hãy push project này lên branch GitOps mà bạn khai báo trước khi chạy deploy/chaincode ops.

## Lệnh Ansible tương đương

Nếu bạn chạy trực tiếp trong Bevel build container hoặc controller machine:

```bash
ansible-playbook -vv /home/bevel/platforms/shared/configuration/site.yaml \
  --inventory-file=/home/bevel/platforms/shared/inventory/ \
  -e "@/home/bevel/build/network.yaml" \
  -e "ansible_python_interpreter=/usr/bin/python3"
```

## Ghi chú bảo mật

Không commit token thật vào Git. `network.yaml` trong starter này chứa placeholder. Khi đi production, nên đưa secret vào secret manager hoặc file private không commit.
