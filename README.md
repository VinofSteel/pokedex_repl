# POKÉMON REPL
Essa é uma aplicação ***REPL*** (*Read-eval-print loop*) que roda em uma CLI e simula o funcionamento de uma Pokédex, usando a fantástica [PokéAPI](https://pokeapi.co/).
Nesta API, a *standard library* de Go foi utilizada exclusivamente, com ênfase no aprendizado de:
- Criação de CLI (*Command line interface*)
- Prática de execução de requisições HTTP em um contexto de API
- Familiarização com *tooling* e desenvolvimento local
- *Caching* e sua utilização para melhora de performance, em conjunto com o próximo item
- Utilização de concorrência para otimização de APIs em Go (Goroutines, Mutexes e Channels)
- Testes automatizados (especialmente unitários)