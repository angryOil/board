# 필독 브랜치 규칙

main - prod, stage - stage , dev - 개발서버

pr 순서 feat => dev => (hotfix/bug)stage => main

1. 최신 dev브랜치에서 feature 만들기
2. dev에 push전 dev pull 받기
3. bug/hotfix 를 제외한 브랜치(ex:feat)로 main/stage에 직접pr금지

## board 게시판 입니다.

게시판으로서 제목과 내용을 작성할수있습니다.

대부분의 기능은 todoList와 비슷합니다.

## entity 구조

```text
    board{
        id
        title
        content
        cafe_id
        board_type
        writer(member_id)
        created_at
        last_updated_at
    }
```

### makefile

```shell
# local postgres run (docker-compose)
make local-db
# local postgres migrate init
make local-init
# local postgres apply migrate
make local-migrate
```

# swagger 설정 [출처](https://www.soberkoder.com/swagger-go-api-swaggo/)

## dev 설정

```shell
go get -u github.com/swaggo/swag/cmd/swag
go get -u github.com/swaggo/http-swagger
go get -u github.com/alecthomas/template
```

## main에

```code
   import (_ "[project명]/docs")
```

```shell
# swagger json 생성   swag init -g [project main path].go
swag init -g cmd/app/main.go
```

## [스웨거 링크](http://localhost:8082/swagger/index.html)