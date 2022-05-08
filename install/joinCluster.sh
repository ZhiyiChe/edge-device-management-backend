#!/bin/bash
# 测试网络连通性
ping -c 3 master.zhiyiche.cn > /dev/null
if [ $? -ne 0 ]; then
    echo "cannot reach master.zhiyiche.cn"
    exit 1
fi
# 下载相关安装文件
curl http://master.zhiyiche.cn:8080/v1/downloadK3s -o k3s
curl http://master.zhiyiche.cn:8080/v1/downloadK3sImages -o k3s-airgap-images-amd64.tar
curl http://master.zhiyiche.cn:8080/v1/downloadInstallScript -o install.sh
# 将tar文件放在images目录下
mkdir -p /var/lib/rancher/k3s/agent/images/
mv ./k3s-airgap-images-amd64.tar /var/lib/rancher/k3s/agent/images/ -f
# 将k3s二进制文件放在/usr/local/bin/k3s路径下
chmod +x ./k3s
mv ./k3s /usr/local/bin/k3s -f
# 获取Token
token=$(curl http://master.zhiyiche.cn:8080/v1/queryToken -s)
echo "http://master.zhiyiche.cn:8080/v1/queryToken resp: $token"
# 执行K3s官方提供的安装脚本, 携带token加入集群
chmod +x ./install.sh
K3S_URL=https://master.zhiyiche.cn:6443 K3S_TOKEN=$token INSTALL_K3S_SKIP_DOWNLOAD=true ./install.sh
# 删除多余文件
rm ./install.sh