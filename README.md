# Перевод [go.dev/tour](https://go.dev/tour/) на русский язык

Проект создан как:

- форк `golang/website`
- изменения в `_content/tour`
- изменения хранятся в ветке "ru"
- всё лишнее удалено

## Сборка

Сборка выполняется через Makefile.  
Готовые бинарники складываются в каталог `dist/`.

```bash
make
```

Готовые бинарники также автоматически собираются через GitHub Actions,
и их можно скачать со [страницы с релизами](https://github.com/sopov/golang-tour-ru/releases).

В релиз включены сборки для Linux, macOS и Windows, собранные из исходного кода через [GitHub Actions](./.github/workflows/release.yml).

## Запуск

```bash
./dist/tour-linux-amd64 -http 0.0.0.0:3999
```

или любой другой бинарь из `dist/` для вашей платформы.

### Список файлов

```bash
git diff --diff-filter=M --name-only master..ru # измененные
git diff --diff-filter=A --name-only master..ru # добавленные
git diff --diff-filter=D --name-only master..ru # удаленные
```
