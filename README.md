# LLM Context Extractor

Ferramenta em Go para extrair código de um diretório para um único arquivo JSON, otimizado para uso como contexto em LLMs.

## Funcionalidades

- Extrai automaticamente arquivos de código fonte (.go, .js, .ts, .py, .java, etc.)
- Inclui arquivos de configuração (.yaml, .yml, .json, .toml, .ini)
- Inclui arquivos de infraestrutura (Dockerfile, Makefile)
- Remove todos os comentários do código
- Exclui automaticamente arquivos de documentação, testes e dependências
- Suporta padrões de exclusão customizados com wildcards
- Gera JSON com estrutura simples: `{"path": "conteúdo"}`

## Instalação

```bash
# Clonar o repositório
git clone https://github.com/crystian/llm-context-extractor.git
cd llm-context-extractor

# Compilar
go build -o llm-context-extractor

# Ou usar go run
go run main.go -input ./seu-projeto -output contexto.json
```

## Uso

```bash
# Uso básico
./llm-context-extractor -input ./seu-projeto -output contexto.json

# Com verbose para ver progresso detalhado
./llm-context-extractor -input ./seu-projeto -output contexto.json -verbose

# Excluir diretórios ou arquivos específicos
./llm-context-extractor -input ./seu-projeto -output contexto.json -exclude "node_modules,vendor"

# Excluir todo o conteúdo de um diretório (/**)
./llm-context-extractor -input ./seu-projeto -output contexto.json -exclude "docs/**,tests/**"

# Excluir subdiretórios imediatos (/*/)
./llm-context-extractor -input ./seu-projeto -output contexto.json -exclude "temp/*/"

# Excluir arquivos por padrão (wildcard)
./llm-context-extractor -input ./seu-projeto -output contexto.json -exclude "*.test.go,*.md"

# Múltiplos padrões de exclusão
./llm-context-extractor -input ./seu-projeto -output contexto.json -exclude "config,legacy.go,*.tmp"

# Usar go run diretamente
go run main.go -input ./seu-projeto -output contexto.json
```

## Flags

- `-input`: Diretório de entrada (obrigatório)
- `-output`: Arquivo JSON de saída (padrão: context.json)
- `-verbose`: Mostrar progresso detalhado
- `-exclude`: Padrões adicionais de exclusão (opcional, separados por vírgula)

### Padrões de Exclusão

O parâmetro `-exclude` suporta:

- **Diretórios**: `config`, `docs`, `temp`
- **Arquivos específicos**: `legacy.go`, `old_file.js`
- **Wildcards**: `*.go`, `*.md`, `test_*`
- **Conteúdo de diretório**: `docs/**`, `tests/**` (exclui todo o conteúdo)
- **Subdiretórios imediatos**: `temp/*/` (exclui subdiretórios do nível 1)
- **Caminhos**: `src/utils`, `config/*.yaml`
- **Múltiplos padrões**: `config,legacy.go,*.tmp`

## Arquivos Incluídos

### Código Fonte

- Go: `.go`
- JavaScript/TypeScript: `.js`, `.ts`
- Python: `.py`
- Java: `.java`
- C/C++: `.c`, `.cpp`, `.h`
- Rust: `.rs`
- Ruby: `.rb`
- PHP: `.php`

### Configuração

- YAML: `.yaml`, `.yml`
- JSON: `.json`
- TOML: `.toml`
- INI: `.ini`, `.cfg`

### Infraestrutura

- `Dockerfile`, `dockerfile`
- `Makefile`, `makefile`

## Arquivos Excluídos Automaticamente

### Arquivos

- `README*`, `CHANGELOG*`, `LICENSE*`, `CONTRIBUTING*`
- Arquivos `.md`, `.txt`, `.rst`, `.log`

### Diretórios

- **Controle de versão**: `.git`
- **Dependências**: `node_modules`, `vendor`
- **Ambientes virtuais**: `.venv`, `venv`, `env`, `.env`
- **Testes**: `test`, `spec`, `tests`, `__tests__`
- **Build**: `dist`, `build`, `target`, `bin`, `obj`

### Testes

- Arquivos `_test.go`
- Arquivos com `.test.` ou `.spec.` no nome

## Exemplo de Saída

```json
{
  "src/main.go": "package main\n\nimport \"fmt\"\n\nfunc main() {\n\tfmt.Println(\"Hello, World!\")\n}",
  "config/config.yaml": "server:\n  port: 8080\n  host: localhost",
  "Dockerfile": "FROM golang:1.21\nWORKDIR /app\nCOPY . .\nRUN go build"
}
```

## Estrutura do Projeto

```text
llm-context-extractor/
├── go.mod
├── go.sum
├── main.go
├── extractor/
│   ├── extractor.go       # Lógica principal de extração
│   ├── filters.go         # Filtros de inclusão/exclusão
│   ├── comment_remover.go # Remoção de comentários
│   └── extractor_test.go  # Testes unitários
└── README.md
```

## Desenvolvimento

### Executar Testes

```bash
go test ./extractor/... -v
```

### Formatar Código

```bash
go fmt ./...
```

## Requisitos

- Go 1.22 ou superior

## Licença

MIT


