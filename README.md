
- <img src ="https://img.shields.io/badge/PostgreSQL-316192?style=for-the-badge&logo=postgresql&logoColor=white"/>
- <img src ="https://img.shields.io/badge/Go-00ADD8?style=for-the-badge&logo=go&logoColor=white"/>

### Requirements

- postgresql@14 [how to set up postgresql@17](./docs/postgresql.md)

### How to Start

1. Run Database Server (local)

    ```shell
    brew services start postgresql@14
    ```

2. Run Server (local)

    ```shell
    go run cmd/blog-server/main.go
    ```
