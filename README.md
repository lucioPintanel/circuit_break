# Circuit Break
## Descrição
O projeto Circuit Break é uma implementação de um padrão de circuito (circuit breaker) em sistemas distribuídos para melhorar a resiliência e a estabilidade. Este projeto busca oferecer uma solução fácil de usar para gerenciar falhas em sistemas complexos, limitando os danos causados por falhas repetidas.

## Funcionalidades
* Detecção automática de falhas
* Retentativas configuráveis
* Timeouts personalizado
* Monitoramento e logging
* Modo half-open para recuperação gradual

## Tecnologias Utilizadas
* Linguagem de Programação: [Go]
* Frameworks: [github.com/sony/gobreaker]

## Instalação
Para clonar e rodar este projeto, você vai precisar do [Golang] instalado na sua máquina.
```
git clone https://github.com/lucioPintanel/circuit_break.git
cd circuit_break
go mod tidy
```

## Como Usar
Depois de instalar as dependências, você pode executar o teste:
```
go test ./...
```

## Contribuição
Contribuições são bem-vindas! Por favor, abra uma issue ou faça um fork do repositório e envie um pull request com suas melhorias.

## Contato
Se você tiver alguma dúvida ou sugestão, sinta-se à vontade para entrar em contato:
* Nome: Lucio Pintanel
* Email: lm.pintanel@gmail.com
* LinkedIn: www.linkedin.com/in/lucio-pintanel