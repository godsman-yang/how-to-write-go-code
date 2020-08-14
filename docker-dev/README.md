# golang docker development environment

## 사전준비
#### 폴더 구성
- docker-compose.yml 파일 위치는 소스의 docker-dev 폴더
  - docker-compose 기준 상위 폴더에 소스의 root가 있음.

#### linux 환경
- wsl2 추천(windows)
  - [Linux용 Windows 하위 시스템 설명서](https://docs.microsoft.com/ko-kr/windows/wsl/)

#### docker 설치
- Docker Desktop 추천(windows, macOS)
  - [Docker DeskTop](https://www.docker.com/products/docker-desktop)

#### visual studio code
- Remote Development 환경 구성
- [Developing inside a Container](https://code.visualstudio.com/docs/remote/containers)

## 코드 가져오기
```
$ git clone https://github.com/godsman-yang/how-to-write-go-code.git
$ cd how-to-write-go-code
```

## docker 개발환경 생성하기
```
$ docker-compose -f ./docker-dev/docker-compose.yml up -d
```

## 소스 확인하고 빌드하기
```
$ docker exec -it test-go-dev bash
$ go build
```

## docker 개발환경 정지/삭제하기
```
$ docker-compose -f ./docker-dev/docker-compose.yml stop
$ docker-compose -f ./docker-dev/docker-compose.yml down
```

## 도커 삭제 후 다시 실행하기
* 소스 위치를 변경하였거나
* 도커 파일을 수정하였거나
```
$ docker-compose -f ./docker-dev/docker-compose.yml down
$ docker-compose -f ./docker-dev/docker-compose.yml up -d
```
