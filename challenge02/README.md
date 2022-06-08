
# Desafio 02 - GetSocialBubbles

@here Pessoas lindas, bom dia, continuando com os desafios de código aqui aqui temos um novo.

Desafio 02, esse será analisado na Sexta-feira 27/05/2022:

Dado essa estrutura de dados do Tipo Person:

```type Person {
    ID string
    Friends []string
}```

```
var people []Person

people = []Person{
    Person{ID: "0", Friends: []string{"4", "1"}},
    Person{ID: "1", Friends: []string{"2", "3"}},
    Person{ID: "2", Friends: []string{"1"}},
    Person{ID: "3", Friends: []string{"2", "4"}},
    Person{ID: "4", Friends: []string{"3"}},
}
```

Escreva um algoritmo que dado um ID, retorne três Arrays:
1 - Array dos amigos(chave Friends).
2 - Array dos amigos dos amigos(utilizando os valores no array 1, pegue os amigos de cada ID).
3 - Array dos amigos dos amigos dos amigos do dev (utilizando os valores no array 2, pegue os amigos de cada ID).

Não esqueça de remover duplicatas, então se no array 2, existir o ID "1" mais de uma vez, ele não deve ser repetido nesse array, porém, ele pode aparecer em outros arrays.

Exemplo:

Input: `GetSocialBubbles(people, "0")`
Output:
```
[1 3]
[2 3 4]
[1 2 4 3]
```

Vamos lembrar os detalhes:
1 - Eu vou passar um desafio de código, e na sexta-feira discutiremos as soluções de vocês.
2 - Não tem solução errada, ou certa, vamos analisar todas e ver onde cada uma se encaixa.
3 - Por enquanto vão ser uns desafios meio agnósticos de tecnologia/linguagem até que eu consiga organizar melhor.

Vamos as regras:
1 - Se você for utilizar recursos de tecnologias, ou Libs tipo .sort(), você tem que saber o que o .sort() faz.
2 - Pode fazer em qualquer linguagem.
3 - Na sexta nós vamos discutir porquê uma solução pode ser melhor que outra, lembrando que não tem resposta certa, tem situação melhor.
4 - Se vc solucionar, manda o código por aqui, de preferência em arquivo.
5 - Se o desafio disser que o input será de tal forma, acredite que será de tal forma. (
