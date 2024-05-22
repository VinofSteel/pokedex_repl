# POKÉMON REPL
Essa é uma aplicação ***REPL*** (*Read-eval-print loop*) *barebones* que roda em uma CLI e simula o funcionamento de uma Pokédex, usando a fantástica [PokéAPI](https://pokeapi.co/).
Nesta API, a *standard library* de Go foi utilizada exclusivamente, com ênfase no aprendizado de:
- Criação de CLI (*Command line interface*) e manipulação de I/O
- Prática de execução de requisições HTTP em um contexto de API em Go
- Familiarização com *tooling* e desenvolvimento local
- *Caching* e sua utilização para melhora de performance (nada além da *standard library* foi usado, o que significa que o cache só dura conforme a sessão)
- Utilização de concorrência para otimização de APIs em Go (Goroutines, Mutexes e Channels), especialmente na utilização do *map* criado para *cacheing*
- Testes automatizados onde necessário

## RODANDO A CLI
Para rodar a aplicação é bem simples
- Instale a linguagem [Go](https://go.dev/dl/) no seu ambiente, no mínimo versão `1.22.3` já que ela disponibiliza a *runtime* da linguagem.
- Caso você possua a ferramenta de terminal *Make* (nativa na maioria das distros de Linux), rode `make` no terminal no diretório raiz do repositório, isso fará o *build* do executável da CLI e a rodará no seu terminal.
- Caso não possua o *Make* por algum motivo, é possível rodar a aplicação com o comando `go build && ./pokedex-repl`, que fará o mesmo descrito acima.
- Digite o comando `help` para verificar os comandos disponíveis e pressione `enter`.
- Pronto! É só usar a aplicação.
- Essa aplicação possui testes automatizados. Para rodá-los, você pode rodar o comando `make test` no diretório raiz do repositório (ou `go test ./...`).