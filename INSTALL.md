# Instruções de Instalação e Solução de Problemas

## Problema de Versão do Go

Se você encontrar o erro:
```
compile: version "go1.23.2" does not match go tool version "go1.22.3"
```

Isso indica uma incompatibilidade entre a versão do Go especificada no `go.mod` e a versão instalada no sistema.

### Soluções

#### Opção 1: Atualizar o Go para a versão especificada

```bash
# Baixar e instalar a versão mais recente do Go
wget https://go.dev/dl/go1.23.2.linux-amd64.tar.gz
sudo rm -rf /usr/local/go
sudo tar -C /usr/local -xzf go1.23.2.linux-amd64.tar.gz

# Verificar a versão
go version
```

#### Opção 2: Ajustar o go.mod para a versão instalada

```bash
# Verificar qual versão você tem
go version

# Editar go.mod para usar a versão instalada
# Se você tem go1.22.3, altere a linha para:
# go 1.22.3
```

#### Opção 3: Usar go mod tidy para limpar dependências

```bash
go clean -cache
go clean -modcache
go mod tidy
```

## Instalação em Ambientes Diferentes

### Linux

```bash
# Baixar Go
wget https://go.dev/dl/go1.23.2.linux-amd64.tar.gz

# Instalar
sudo tar -C /usr/local -xzf go1.23.2.linux-amd64.tar.gz

# Adicionar ao PATH
export PATH=$PATH:/usr/local/go/bin
echo 'export PATH=$PATH:/usr/local/go/bin' >> ~/.bashrc
source ~/.bashrc

# Verificar
go version
```

### macOS

```bash
# Usar Homebrew
brew install go

# Ou baixar manualmente
# Visite https://go.dev/dl/ e baixe a versão para macOS
```

### Windows

1. Baixe o instalador em https://go.dev/dl/
2. Execute o instalador
3. Reinicie o terminal

## Compilação

```bash
# Compilar o projeto
go build -o llm-context-extractor

# Executar
./llm-context-extractor -input ./seu-projeto -output contexto.json
```

## Testes

```bash
# Executar todos os testes
go test ./... -v

# Executar apenas testes do extractor
go test ./extractor/... -v
```

## Exemplo de Uso

```bash
# Criar um projeto de teste
mkdir test-project
cd test-project

# Criar alguns arquivos
echo 'package main
import "fmt"
func main() {
    fmt.Println("Hello")
}' > main.go

echo '# Config
port: 8080' > config.yaml

# Voltar e executar o extractor
cd ..
./llm-context-extractor -input ./test-project -output contexto.json

# Ver o resultado
cat contexto.json
```

## Solução de Problemas Comuns

### Erro: "command not found: go"

O Go não está instalado ou não está no PATH. Siga as instruções de instalação acima.

### Erro: "cannot find package"

Execute `go mod tidy` para baixar as dependências.

### Erro: "permission denied"

Verifique as permissões do diretório ou use `sudo` para instalar o Go.

### Erro de versão incompatível

Siga as soluções na seção "Problema de Versão do Go" acima.