# fct7-devops-ci-desafio
Code Education | Full Cycle Turma 7 | DevOps | Iniciando com Docker - Desafio de CI

Para executar os testes unitários

```bash
go test ./github.com/kaissi/fct7-devops-ci-desafio -v
```

Para executar a aplicação

```bash
go run ./github.com/kaissi/fct7-devops-ci-desafio
```

Para gerar a imagem docker

```bash
 docker build . -t kaissi/fct7-devops-ci-desafio:latest
```

Para testar a imagem docker gerada

```bash
docker run --rm kaissi/fct7-devops-ci-desafio:latest
```
