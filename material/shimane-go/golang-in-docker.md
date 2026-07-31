# Docker で Golang コードを動かす

## Ubuntu にインストールする

```
$ sudo apt -y install apt-transport-https ca-certificates curl software-properties-common # apt-transport-https のみでOK
$ curl -fsSL https://download.docker.com/linux/ubuntu/gpg | sudo apt-key add -
$ sudo add-apt-repository "deb [arch=amd64] https://download.docker.com/linux/ubuntu $(lsb_release -cs) stable"
$ sudo add-apt-repository "deb [arch=amd64] https://download.docker.com/linux/ubuntu disco stable"
$ sudo apt update
$ sudo apt install docker-ce docker-ce-cli containerd.io
```

```
ca-certificates はすでに最新バージョン (20190110) です。
ca-certificates は手動でインストールしたと設定されました。
curl はすでに最新バージョン (7.65.3-1ubuntu3) です。
software-properties-common はすでに最新バージョン (0.98.5) です。
software-properties-common は手動でインストールしたと設定されました。
以下のパッケージが新たにインストールされます:
  apt-transport-https
```

```
$ sudo docker run hello-world
Unable to find image 'hello-world:latest' locally
latest: Pulling from library/hello-world
1b930d010525: Pull complete 
Digest: sha256:4fe721ccc2e8dc7362278a29dc660d833570ec2682f4e4194f4ee23e415e1064
Status: Downloaded newer image for hello-world:latest

Hello from Docker!
This message shows that your installation appears to be working correctly.

To generate this message, Docker took the following steps:
 1. The Docker client contacted the Docker daemon.
 2. The Docker daemon pulled the "hello-world" image from the Docker Hub.
    (amd64)
 3. The Docker daemon created a new container from that image which runs the
    executable that produces the output you are currently reading.
 4. The Docker daemon streamed that output to the Docker client, which sent it
    to your terminal.

To try something more ambitious, you can run an Ubuntu container with:
 $ docker run -it ubuntu bash

Share images, automate workflows, and more with a free Docker ID:
 https://hub.docker.com/

For more examples and ideas, visit:
 https://docs.docker.com/get-started/
```

## docker グループの追加

```
$ sudo adduser $USER -G docker
$ sudo usermod -aG docker $USER
```




- [Get Docker Engine - Community for Ubuntu | Docker Documentation](https://docs.docker.com/install/linux/docker-ce/ubuntu/)
- [Dockerインストールメモ - Qiita]( https://qiita.com/n-yamanaka/items/ddb18943f5e43ca5ac2e)
- [Ubuntu 19.10にDockerをインストールする - Qiita](https://qiita.com/rarudonet/items/8c5e99f12adc85c73729)
- [GitHub - gitkado/docker-run-go-sample](https://github.com/gitkado/docker-run-go-sample)
- https://tracpath.com/works/devops/how_to_install_the_docker/
