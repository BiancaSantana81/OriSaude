# OríSaude

**OríSaude** é a plataforma final do curso **Aprofunda PretaLab**, criada para conectar pessoas negras a profissionais negros da área da saúde, promovendo acesso facilitado a serviços de saúde com representatividade.


**Próximos passos em desenvolvimento:**

- Testes unitários e de integração
- Pipeline de CI/CD no GitHub Actions
- Autenticação e autorização para profissionais e usuários

---

## Sobre o Projeto

O projeto foi desenvolvido como parte do curso **Aprofunda PretaLab**, com o objetivo de construir uma API REST para gerenciamento de profissionais da saúde, aplicando boas práticas como:

- Arquitetura limpa (Clean Architecture)
- Testes unitários e de integração [EM CONSTRUÇÃO]
- Injeção de dependências

A aplicação permite **criar, consultar, atualizar e deletar profissionais**, com integração ao **Firebase Realtime Database** como banco de dados principal.

## 🌐 Disponível em

A API está atualmente **rodando no Google Cloud Run**:

[link para aplicação](https://ori-saude-api-69026747543.us-central1.run.app/professionals/)

> Você pode acessar diretamente ou usar esta URL como **base URL** para testes no Postman.

---

## 🧩 Integração com Banco de Dados

- A aplicação utiliza **Firebase Realtime Database** para armazenar os dados dos profissionais.
- A arquitetura baseada em repositórios permite fácil manutenção e substituição do banco, se necessário.
- Em testes, a aplicação pode se conectar a instâncias de banco separadas, garantindo validação real dos fluxos.

---

## Modelagem de Dados

### 👤 Schema Professional

```json
{
  "id": "string",
  "name": "string",
  "category": "string",
  "city": "string",
  "contact": "string"
}
```

## 🌐 Endpoints da API

| 🛠 Método | 📍 Rota | 📝 Descrição | 📦 Request Body | 📤 Response | 🔑 Autenticação |
|-----------|--------|-------------|----------------|------------|----------------|
| POST      | `/professionals`       | Cria um novo profissional | `{ "name": "string", "category": "string", "city": "string", "contact": "string" }` | `201 Created` `{ "id": "string", "name": "...", "category": "...", "city": "...", "contact": "..." }` | ❌ Não |
| GET       | `/professionals`       | Lista todos os profissionais | N/A | `200 OK` `[ { "id": "...", "name": "...", "category": "...", "city": "...", "contact": "..." } ]` | ❌ Não |
| GET       | `/professionals/:id`   | Retorna profissional pelo ID | N/A | `200 OK` `{ "id": "...", "name": "...", "category": "...", "city": "...", "contact": "..." }` ou `404 Not Found` | ❌ Não |
| PUT       | `/professionals/:id`   | Atualiza profissional pelo ID | `{ "name": "string", "category": "string", "city": "string", "contact": "string" }` | `200 OK` `{ "id": "...", "name": "...", "category": "...", "city": "...", "contact": "..." }` | ❌ Não |
| DELETE    | `/professionals/:id`   | Remove profissional pelo ID | N/A | `204 No Content` | ❌ Não |


---

## Tecnologias Utilizadas

- Go (Golang)
- Gin (framework web)
- Firebase Realtime Database
- Arquitetura Limpa (Clean Architecture)
- Testes unitários e integração [EM CONSTRUÇÃO]
- GitHub Actions (CI/CD) [EM CONSTRUÇÃO]
- Google Cloud para deploy [EM CONSTRUÇÃO]

---

## Uso

Pré-requisitos

```
Go 1.20+

Firebase project configurado

Variáveis de ambiente:

FIREBASE_CREDENTIALS → Caminho para o arquivo JSON de credenciais

FIREBASE_DB_URL → URL do Realtime Database
```

### Executando a aplicação
```
# Inicializa Firebase e roda o servidor
go run ./src/main.go

```

## 🌐 Teste a API no Postman

Para facilitar, você pode importar todos os endpoints da API diretamente no Postman usando nossa coleção:

[Importar coleção Postman](https://joint-operations-operator-4654730-7615344.postman.co/workspace/Bianca-Santana's-Workspace~44b97a57-f7b6-477e-8895-4fdef975a126/collection/47108375-2d68b8a5-80c2-40b8-9c84-001ac19c2df5?action=share&source=copy-link&creator=47108375)

> 🔧 Observação: Se você estiver rodando a API localmente (`http://localhost:8080`), altere o **base URL** no Postman para apontar para o seu servidor local.
