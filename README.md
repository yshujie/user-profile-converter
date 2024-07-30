# user-profile-converter

用户画像转化器

## 功能特性

## 软件架构
```
user-profile-converter/
│
├── cmd/
│   └── app/
│       └── main.go
├── internal/
│   ├── app/
│   │   └── business_logic.go
│   ├── model/
│   │   ├── user.go
│   │   ├── child.go
│   │   └── student.go
│   ├── repository/
│   │   ├── mysql/
│   │   │   ├── user.go
│   │   │   ├── child.go
│   │   │   └── student.go
│   │   └── mongodb/
│   │       └── mongo.go
│   ├── service/
│   │   └── data_service.go
├── pkg/
│   ├── mysql/
│   │   ├── logging_driver.go
│   │   └── mysql.go
│   └── mongodb/
│       └── mongodb.go
├── Dockerfile
├── Dockerfile.dev
├── docker-compose.yml
├── docker-compose.override.yml
├── docker-compose.prod.yml
├── .env
├── .env.prod
├── .air.toml
├── go.mod
├── go.sum
```

## 快速开始

### 依赖检查

### 构建

#### 拉取项目源码
```
mkdir ~/code/go/src/github.com/yshujie
cd ~/code/go/src/github.com/yshujie
git clone git@github.com:yshujie/user-profile-converter.git
```

### 构建&运行

#### 开发环境运行 Docker 容器
```
docker-compose up --build
```

#### 生产环境运行 Docker 容器
```
docker-compose -f docker-compose.yml -f docker-compose.prod.yml up --build
```

## 使用指南
