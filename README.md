# Prova Detran CLI

Simulado de provas do Detran.

## Como instalar

```
go install github.com/jefersonf/prova-detran
```

Em seguida inicie um novo simulado.

```
prova-detran simulado
```

Para executar a partir do código fonte basta clonar este repositório.

```
cd prova-detran
go build .
```
e em seguida inicie um novo simulado.

```
./prova-detran simulado
```

## Opções (`help`)

```
Estude para o exame teórico do Detran por meio do Prova Detran,
        uma ferramenta de linha de comando que gera simulados da prova 
        a partir de questões reais e para todos os conteúdos abordados.
        Documentação completa está disponível em github.com/jefersonf/prova-detran

Usage:
  provadetran [command]

Available Commands:
  completion  Generate the autocompletion script for the specified shell
  help        Help about any command
  resultados  Mostra o resultado do simulado
  simulado    Inicia um novo simulado
  version     Imprime a versão do Prova Detran

Flags:
  -h, --help   help for provadetran
```

## Exemplo

![exemplo de simulado](/sample-output.png)

## Mini-FAQ

>As questões seguem o formato do exame oficial?

>**Sim**
---
>As questões abordam todos os temas cobrados?

>**Sim**
---
>É possível visualizar as placas de sinalização que os enunciados fazem referência?

>**Depende! Nesses casos uma representação gráfica aproximada é utilizada no próprio terminal, devido as limitações comuns a uma CLI**

## Como contruir

É possível contruir de duas formas:

1. Adicionando questões ao banco de questões da CLI. (Checar o formato do [banco de questões](#banco-de-questoes))
2. Melhorando e/ou evoluindo o código! Clone e crie uma issue, PRs são bem-vindos (_There's always room for improvements_).
3. Divulgue! Se é brasileiro sabe bem como é "peso" tirar a primeira habilitação, a tão cara CNH. Criei essa ferramenta para estudar a legislação de Trânsito enquanto praticava outra coisa, o desenvolvimento de CLIs em Go. Espero que seja útil se te serve para ao menos um dos casos.  