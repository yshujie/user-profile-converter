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
├── go.mod
└── go.sum
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

#### 构建 Docker 镜像
```
docker-compose build
```

### 运行

#### 运行 Docker 容器
```
docker-compose up -d
```

## 使用指南
