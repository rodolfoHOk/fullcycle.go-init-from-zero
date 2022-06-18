# Go Começando do Zero 🚀

> Projeto criado durante o esquenta da Imersão FullCycle 8

## 👨‍💻 Tecnologias e bibliotecas utilizadas 👩‍💻

- GO : Linguagem programação

### 📚 bibliotecas adicionais 🗃️

- go-sqlite3
- echo

## 📖 Guia 📃

- Instalar Go

        acessar: https://go.dev/doc/install e baixar go?.linux-amd64.tar.gz
        sudo tar -C /usr/local -xzf go1.17.1.linux-amd64.tar.gz
        export PATH=$PATH:/usr/local/go/bin
        go version

- Rodar código no desenvolvimento

        go run nome-do-arquivo.go

- Baixar biblioteca externa

        go mod init goInitFromZero <- nome do modulo
        go mod tidy

- Para compilar go-sqlite3

        sudo apt install --reinstall build-essential

- Criar tabela

        sudo apt install sqlite3
        sqlite3 test2.db
        create table products (name text, price number);
        ^D

- Baixar biblioteca echo após baixar go-sqlite3

        go mod tidy

## 🔥 Repositórios da Imersão FullCycle 8 ✨

- [Arquitetura Limpa com Typescript](https://github.com/rodolfoHOk/fullcycle.typescrit-clean-arch)

- [React Maps](https://github.com/rodolfoHOk/fullcycle.react-maps)

- [Go Iniciando do Zero](https://github.com/rodolfoHOk/fullcycle.go-init-from-zero)
