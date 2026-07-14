# 🚀 Workshop GDG Americana: Assistente de Código Inteligente com Go & Gemini SDK

Seja muito bem-vindo(a) ao workshop prático de **Go (Golang)** integrado ao ecossistema de inteligência artificial do Google! 
Este projeto foi desenvolvido especialmente para o **Google I/O Extended** no **GDG Americana**, com o objetivo de apresentar a linguagem Go para estudantes e iniciantes de forma extremamente prática, moderna e alinhada com as melhores práticas de engenharia de software do mercado.

Neste repositório, você encontrará a implementação de uma API HTTP em Go que atua como um **Tutor Inteligente de Programação**. 
O projeto analisa códigos enviados pelos usuários, explicando seu funcionamento e sugerindo melhorias de performance e segurança utilizando o novo **Gemini SDK oficial do Google**.

---

## 🗺️ Estrutura do Workshop (120 minutos)

O nosso tempo de 2 horas será rigorosamente dividido para garantir o melhor aproveitamento entre teoria e prática:

*   **`[00:00 - 00:15]` ── Introdução: Cloud Shell Editor + Go no Ecossistema Google**
    *   Por que o Google criou o Go em 2009 e como a linguagem se tornou a espinha dorsal da infraestrutura em nuvem moderna.
    *   Apresentação e ambientação no [Google Cloud Shell Editor](https://shell.cloud.google.com) (ambiente 100% online e gratuito, sem necessidade de instalar nada localmente).
*   **`[00:15 - 00:30]` ── Arquitetura Hexagonal & SOLID (Ports & Adapters de forma simples)**
    *   Desmistificando a Arquitetura Hexagonal: como isolar o núcleo da nossa aplicação (regras de negócio) de detalhes externos de infraestrutura (bancos de dados, APIs de terceiros, servidores HTTP).
    *   Inversão de Dependência (a letra **D** do SOLID): acoplamento fraco por meio de interfaces implícitas do Go.
*   **`[00:30 - 01:25]` ── LIVE CODING: Construindo o Assistente de Código com Gemini SDK**
    *   Mão na massa! Criação da estrutura de pastas do zero, definição do domínio e implementação do adapter de integração usando o novíssimo SDK unificado `google.golang.org/genai`.
*   **`[01:25 - 01:45]` ── Testes Unitários com Mocks e Observabilidade (`slog` + `Context`)**
    *   Escrevendo testes de unidade velozes e determinísticos utilizando Mocks manuais, provando o valor do design desacoplado.
    *   Introdução a logs estruturados de alta performance com a biblioteca nativa `log/slog` e propagação de contextos de execução com `context.Context`.
*   **`[01:45 - 02:00]` ── Q&A & Próximos Passos**
    *   Espaço aberto para dúvidas, discussões arquiteturais e recursos recomendados para continuar a jornada de aprendizado em Go.

---

## 🏗️ Visão Geral da Arquitetura (Hexagonal)

Nossa aplicação adota a **Arquitetura Hexagonal (Ports & Adapters)**. Isso garante que a nossa lógica de negócios não conheça detalhes do mundo exterior.

```text
               +---------------------------------------------------+
               |                    EXTERNAL                       |
               |                                                   |
  [Client] --------> [Adapter HTTP (Port de Entrada / Inbound)]   |
               |                        |                          |
               |                        v                          |
               |             [Caso de Uso (Core/Domain)]           |
               |                        |                          |
               |                        v                          |
               |    [Adapter Gemini (Port de Saída / Outbound)] ------> [Google Gemini API]
               |                                                   |
               +---------------------------------------------------+
```

### Estrutura do Projeto
- **`cmd/`**: Define a porta de entrada em nosso projeto.
- **`internal/core/domain/`**: Define as entidades de negócio e as *Ports* (interfaces que descrevem as necessidades do sistema).
- - **`internal/core/ports`**: Define D do SOLID, onde definimos nossas interfaces.
- **`internal/core/usecase/`**: Contém as regras de negócio puras (casos de uso) que orquestram o fluxo de execução.
- **`internal/adapter/gemini/`**: Implementação da *Port* de saída que consome a API do Gemini.
- **`internal/adapter/http/`**: Implementação da *Port* de entrada que recebe requisições HTTP e as encaminha para o caso de uso correspondente.

## 🛠️ Como Executar o Projeto no Google Cloud Shell

Siga este guia passo a passo para colocar a aplicação para rodar online sem instalar nenhuma dependência na sua máquina.

### 1. Acessar o Ambiente

Abra o navegador e acesse:

<https://shell.cloud.google.com>

> O Cloud Shell já possui o compilador Go mais recente pré-instalado em sua máquina virtual dedicada.

```bash
# Criar o diretório do workshop e acessá-lo
mkdir -p gdg-workshop && cd gdg-workshop

# Clonar o repositório ou inicializar a estrutura
go mod init gdg-workshop

# Adicionar o SDK oficial de IA Generativa do Google
go get google.golang.org/genai
```

### 3. Obter sua Chave de API do Gemini

Para que o SDK consiga se autenticar com os servidores do Google, você precisará de uma chave de API:

1. Acesse o **Google AI Studio**.
2. Faça login com sua conta Google e clique no botão **Get API Key**.
3. Crie uma nova chave e copie o código gerado.
4. No terminal do Cloud Shell, defina a variável de ambiente:

  ```bash
  export GEMINI_API_KEY="COLE_SUA_CHAVE_AQUI"
  ```
### 4. Executar os Testes Unitários
Como o nosso código é modular e possui dependências invertidas, você pode validar as regras de negócio sem gastar créditos de API e de forma extremamente rápida:

```bash
go test -v ./...
```

### 5. Iniciar o Servidor
Com o ambiente configurado, inicialize a sua API HTTP:

```bash
go run main.go
```

O console exibirá que o servidor está escutando na porta :8080 com os logs estruturados no formato JSON!

6. Testar a API
Abra um segundo terminal no Cloud Shell (clicando no ícone de + no menu superior do terminal) ou faça um teste usando curl no próprio terminal local:

```bash
curl -X POST http://localhost:8080/api/analyze \
  -H "Content-Type: application/json" \
  -d '{
    "language": "javascript",
    "code": "const soma = (a, b) => a + b; console.log(soma(5, 5));"
  }'
```

## 📈 Melhores Práticas & Destaques de Go Moderno Aplicados

Este projeto serve como um modelo moderno de escrita de código Go:

- **Roteamento Nativo Avançado (Go 1.22+)**: Utilizamos o `http.NewServeMux` atualizado para declarar métodos HTTP de forma direta e limpa, eliminando a dependência excessiva de frameworks de roteamento de terceiros.
- **Logs Estruturados (`slog`)**: Implementamos a biblioteca `log/slog` de alta performance, projetada nativamente no Go 1.21. Seus logs estruturados em JSON são perfeitamente compatíveis com o agregador do Google Cloud Logging.
- **Gerenciamento de Contexto (`context.Context`)**: Propagamos o contexto HTTP nativo por todas as camadas até o SDK do Gemini, permitindo limites estritos de timeout e garantindo observabilidade com cancelamentos limpos de rotina.
- **Segurança no Gemini SDK**: Consumo robusto do retorno utilizando o método de extração segura `resp.Text()`, evitando erros comuns de ponteiros vazios de string.

## 📚 Referências & Documentação Oficial

Para se aprofundar ainda mais na linguagem e continuar estudando:

- **Histórico de Releases do Go:** <https://go.dev/doc/devel/release>
- **Documentação Geral do Go:** <https://go.dev/doc/>
- **Ambiente de Testes Online:** <https://go.dev/play/>
- **SDK do Gemini em Go:** <https://github.com/google/generative-ai-go>
- **Boas Práticas de Teste em Go:** <https://pkg.go.dev/testing>
- **Guia do Cloud Shell Editor:** <https://cloud.google.com/shell/docs/editor-overview>

---

Desenvolvido com 💙 para a comunidade do **GDG Americana** e entusiastas do ecossistema Google. **Bom workshop!**