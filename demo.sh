#!/bin/bash

# Script de demonstração do LLM Context Extractor

echo "=== LLM Context Extractor - Demonstração ==="
echo ""

# Criar diretório de teste
echo "1. Criando projeto de teste..."
mkdir -p demo-project/src demo-project/config demo-project/docs

# Criar arquivo Go com comentários
cat > demo-project/src/main.go << 'EOF'
package main

import "fmt"

// This is a main function
func main() {
	// Print hello world
	fmt.Println("Hello, World!")

	/*
	 * Multi-line comment
	 * that should be removed
	 */

	x := 42 // inline comment
	y := x + 1

	fmt.Printf("x=%d, y=%d\n", x, y)
}
EOF

# Criar arquivo YAML com comentários
cat > demo-project/config/config.yaml << 'EOF'
# This is a config file
# with comments

server:
  port: 8080  # server port
  host: localhost

database:
  # Database configuration
  driver: postgres
  url: "postgres://user:pass@localhost/db"
EOF

# Criar arquivo README (deve ser excluído)
cat > demo-project/README.md << 'EOF'
# README

This is a test project for the LLM context extractor.

## Installation

Run the extractor with:
```bash
go run main.go -input ./demo-project -output context.json
```
EOF

# Criar arquivo de teste (deve ser excluído)
cat > demo-project/src/main_test.go << 'EOF'
package main

import "testing"

func TestMain(t *testing.T) {
	// This is a test function
	if true {
		t.Log("test")
	}
}
EOF

# Criar Dockerfile
cat > demo-project/Dockerfile << 'EOF'
# Dockerfile for test project
# Multi-stage build

FROM golang:1.21 AS builder
WORKDIR /app
COPY . .
RUN go build -o app

FROM alpine:latest
COPY --from=builder /app/app /usr/local/bin/app
CMD ["app"]
EOF

# Criar Makefile
cat > demo-project/Makefile << 'EOF'
# Makefile for test project

.PHONY: build test clean

build:
	go build -o app

test:
	go test ./...

clean:
	rm -f app
EOF

# Criar package.json
cat > demo-project/package.json << 'EOF'
{
  "description": "Test project for LLM context extractor",
  "main": "src/main.go",
  "name": "demo-project",
  "scripts": {
    "build": "go build -o app",
    "test": "go test ./..."
  },
  "version": "1.0.0"
}
EOF

echo "✓ Projeto de teste criado"
echo ""

# Mostrar estrutura do projeto
echo "2. Estrutura do projeto criado:"
find demo-project -type f | sort
echo ""

# Mostrar o que deve ser incluído
echo "3. Arquivos que DEVEM ser incluídos:"
echo "   - src/main.go (código Go)"
echo "   - config/config.yaml (configuração YAML)"
echo "   - Dockerfile (infraestrutura)"
echo "   - Makefile (infraestrutura)"
echo "   - package.json (configuração JSON)"
echo ""

# Mostrar o que deve ser excluído
echo "4. Arquivos que DEVEM ser excluídos:"
echo "   - README.md (documentação)"
echo "   - src/main_test.go (teste)"
echo ""

# Mostrar exemplo de saída esperada
echo "5. Exemplo de saída JSON esperada:"
cat << 'EOF'
{
  "config/config.yaml": "server:\n  port: 8080\n  host: localhost\ndatabase:\n  driver: postgres\n  url: \"postgres://user:pass@localhost/db\"",
  "Dockerfile": "FROM golang:1.21 AS builder\nWORKDIR /app\nCOPY . .\nRUN go build -o app\n\nFROM alpine:latest\nCOPY --from=builder /app/app /usr/local/bin/app\nCMD [\"app\"]",
  "Makefile": ".PHONY: build test clean\n\nbuild:\n\tgo build -o app\n\ntest:\n\tgo test ./...\n\nclean:\n\trm -f app",
  "package.json": "{\n  \"description\": \"Test project for LLM context extractor\",\n  \"main\": \"src/main.go\",\n  \"name\": \"demo-project\",\n  \"scripts\": {\n    \"build\": \"go build -o app\",\n    \"test\": \"go test ./...\"\n  },\n  \"version\": \"1.0.0\"\n}",
  "src/main.go": "package main\n\nimport \"fmt\"\n\nfunc main() {\n\tfmt.Println(\"Hello, World!\")\n\n\tx := 42\n\ty := x + 1\n\n\tfmt.Printf(\"x=%d, y=%d\\n\", x, y)\n}"
}
EOF

echo ""
echo "6. Para executar o extractor:"
echo "   ./llm-context-extractor -input ./demo-project -output contexto.json"
echo ""
echo "   Ou com go run:"
echo "   go run main.go -input ./demo-project -output contexto.json"
echo ""

echo "=== Fim da demonstração ==="
echo ""
echo "O projeto de teste está em: ./demo-project"
echo "Você pode usá-lo para testar o extractor."