# echo-rest-api

My golang training. Writing clean rest-api on echo

### Technology

<a href="https://echo.labstack.com"><img height="80" src="https://cdn.labstack.com/images/echo-logo.svg"></a>

[![Nginx][Nginx-badge]][Nginx-url]

### Installation and launch

1. Clone repo

    ```shell
    git clone https://github.com/HoneyBebra/echo-rest-api.git
    cd echo-rest-api
    ```

2. Create a .env-files as in deploy/.env.example and service/.env.example

3. Deploy

   ```shell
   cd deploy
   docker-compose up -d
   ```

# Architecture

![GDD][Current-architecture-url]

<!-- MARKDOWN LINKS & BADGES -->

[Nginx-url]: https://nginx.org
[Nginx-badge]: https://img.shields.io/badge/nginx-%23009639.svg?style=for-the-badge&logo=nginx&logoColor=white~~[Current-architecture-url]: ./architecture/current_architecture.png
