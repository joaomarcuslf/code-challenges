# Desafio 05

@here Pessoas lindas, bom dia, continuando com os desafios de código aqui aqui temos um novo.

Desafio 05, esse será analisado na Sexta-feira 17/06/2022.

Nesse desafio nós vamos trabalhar com árvores, mais precisamente árvore binária de busca. Resumindo muito, uma árvore binária de busca, é uma estrutura de dados que se organiza de tal forma que quando você precisa inserir um elemento, caso ele tenha um valor menor que o atual, ele será inserido para esquerda, caso contrário, será inserido para direita. Você pode entender melhor com as seguintes referências:

- https://www.youtube.com/watch?v=VmKkAQtnjsM&ab_channel=Programa%C3%A7%C3%A3oDin%C3%A2mica
- https://en.wikipedia.org/wiki/Binary_search_tree
- https://medium.com/arctouch/data-structure-101-part-4-57c19ec8b1d2

O seu algoritmo receberá um array de inteiros, e deverá retornar uma arvore binária de busca.

*Exemplo 01:*

``` // Definição
type BinarySearchTree struct {
  value int
  left BinarySearchTree
  right BinarySearchTree
}
```

```
Entrada: K = [ 2, 3, 1, 4, 6, 7]
Saída:
{
    "value": 2,
    "left": {
        "value": 1,
        "left": null,
        "right": null
    },
    "right": {
        "value": 3,
        "left": null,
        "right": {
            "value": 4,
            "left": null,
            "right": {
                "value": 6,
                "left": null,
                "right": {
                    "value": 7,
                    "left": null,
                    "right": null
                }
            }
        }
    }
}
```

Desafio bônus 01: implemente um algoritmo de busca na árvore binária de busca (retorne true para existente e false para não existente).
*Exemplo 02:*

```
Entrada: K = [ 2, 3, 1, 4, 6, 7]
Saída:
tree := FromArrayToBinarySearchTree(K)

SearchBinarySearchTree(tree, 1) // retorna true
SearchBinarySearchTree(tree, 9) // retorna false
```

Desafio bônus 02: implemente uma forma de balancear a árvore para que ela não fique um lado muito maior que o outro.

*Exemplo 03:*

```
Entrada: K = [ 2, 3, 1, 4, 6, 7]
Saída:
{
    "value": 4,
    "left": {
        "value": 2,
        "left": {
            "value": 1,
            "left": null,
            "right": null
        },
        "right": {
            "value": 3,
            "left": null,
            "right": null
        }
    },
    "right": {
        "value": 7,
        "left": {
            "value": 6,
            "left": null,
            "right": null
        },
        "right": null
    }
}
```

*Vamos lembrar os detalhes:*

1. Eu vou passar um desafio de código, e na sexta-feira discutiremos as soluções de vocês.
2. Não tem solução errada, ou certa, vamos analisar todas e ver onde cada uma se encaixa.
3. Por enquanto vão ser uns desafios meio agnósticos de tecnologia/linguagem até que eu consiga organizar melhor.

*Vamos as regras:*

1. Se você for utilizar recursos de tecnologias, ou Libs tipo .sort(), você tem que saber o que o .sort() faz.
2. Pode fazer em qualquer linguagem.
3. Na sexta nós vamos discutir porquê uma solução pode ser melhor que outra, lembrando que não tem resposta certa, tem situação melhor.
4. Se vc solucionar, manda o código por aqui, de preferência em arquivo.
5. Se o desafio disser que o input será de tal forma, acredite que será de tal forma.
