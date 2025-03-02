### how to set up postgresql@14

### homebrew 를 사용한 설치

아래 homebrew 명령어 실행시 기본적인 database cluster 가 생성이 된다.

```shell
brew install postgresql
```


**설치 확인**

- home brew 를 사용한 설치임이 표시되는듯 하다.
- 버전 명시 없이 다운로드시 현재 기준 14 버전이 다운로드 되는듯 하다. 최신 버전은 17버전으로 보인다.

```shell
psql --version
# psql (PostgreSQL) 14.17 (Homebrew)
```

homebrew 를 사용하여 다운로드 받은 경우 bin path 등은 다음 명령어를 통해 확인할 수 있다.

```shell
brew info postgresql
```

### postgresql server 실행

- background 실행

    ```shell
    brew services start postgresql@14
    ```

    현재 백그라운드에서 데이터베이스 서버가 실행중인지는 다음 명령어를 사용해 확인할 수 있다.

    상태가 `started` 로 표기되어 있음을 확인한다.

    ```shell
    brew services list | grep postgresql start postgresql@14
    ```



- foreground 실행


    ```
    /usr/local/opt/postgresql@14/bin/postgres -D /usr/local/var/postgresql@14
    ```


## FAQ

### 

### database <user-name> does not exist

기본적으로는 "현재 로그인된 사용자명" 을 재료로 기본 데이터베이스가 생성되어야하지만, 이것이 생성되지 않은 상태로 다음 명령어를 사용한 후 재 시도해보자.

```
createdb
```