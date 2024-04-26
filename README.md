Instruções para execução do projeto local com Docker (executar na raiz do projeto):

1. Executar o comando _make docker-compose_. O comando subirá o redis e o redis insight que pode ser acessado por meio da porta 8001.
2. Um vez que o container foi criado, basta configurar o limiter utilizando das variáveis de ambiente no arquivo .env. Sendo eles:
    MAX_REQUESTS_PER_SECOND -> Máximo de requisições por segundo
    BLOCKING_TIME_SECONDS -> Tempo de bloqueio em segundos do IP após atingir o número máximo de requisições
3. O arquivo .env também contém 3 tokens que podem ser enviando à API que contém diferentes valores para o número máximo de requisições por segundo. São enviados como no exemplo a seguir:

http://localhost:8080
API_KEY: bGltaXQ1

4. Testes de persistência de dados podem ser executados no arquivo internal/infra/database/redis/limiter_info_repository_test.go

5. Testes de bloqueio e benchmark de dados podem ser executados no arquivo internal/rate-limiter/rate-limiter_test.go

Chamadas de exemplo podem ser encontradas no arquivo api/api.http