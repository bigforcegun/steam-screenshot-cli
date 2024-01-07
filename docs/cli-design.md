# CLI

```sh
steam-screenshot-cli list -u kekw -g kekw -d .
steam-screenshot-cli sync -u kekw -g kekw -d .
```

В этом случае -g меняет поведение тулы

Кмк это странно - флаг это же **параметр** выполнения, 
Он не должен менять логику программы?

Типа сохраняем игру или сохраняем пользотваеля?
Сохраняем в директорию или в красвую директорию

```sh

steam-screenshot-cli list game -u kekw -g kekw -d .
steam-screenshot-cli sync game -u kekw -g kekw -f -d
steam-screenshot-cli sync user -u kekw -g kekw -f -d
steam-screenshot-cli su -u kekw -g kekw -f -d
steam-screenshot-cli sg -u kekw -g kekw -f -d
steam-screenshot-cli lu -u kekw -g kekw -f -d
steam-screenshot-cli lg -u kekw -g kekw -f -d
ssc lg -u kekw -g kekw -f -d
````
