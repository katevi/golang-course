
# Домашнее задание

## I. Установка необходимого окружения

Для прохождения курса желательно поставить себе какой-нибудь Unix environment. Если основная ОС ноутбука -- Unix, то можно ничего не делать, если Windows -- см. раздел про установку WSL далее.

### Установка WSL 

1. Можно воспользоваться следующей инструкцией: https://learn.microsoft.com/en-us/windows/wsl/install .

### Установка VsCode

1. Нужно поставить VSCode https://code.visualstudio.com/download .
2. И подключить WSL в VSCode, установив WSL extension как тут: https://code.visualstudio.com/docs/remote/wsl-tutorial (эта инструкция включает в себя установку WSL, если WSL уже установлен в прошлом шаге, заново его ставить не нужно).

### Установка Go

1. Предполагается, что Unix cистема или WSL уже поставлены. Если WSL, то go надо устанавливать в WSL, а не на Windows (подключиться через VSCode к WSL и туда скачать архив с go).
2. Поставить Go по инструкции как тут: https://go.dev/doc/install (проверить, что `go version` команда выполняется успешно и показывает версию Go).
3. Поставить Go extension в VSCode https://code.visualstudio.com/docs/languages/go (скорее всего, VSCode будет сам предлагать, когда увидит проект на go).

## II. Задачи

1. Про конкатенацию слайса: https://leetcode.com/problems/concatenation-of-array/description/ 
2. Про пары чисел в слайсе: https://leetcode.com/problems/number-of-good-pairs/description/ 
3. Про консистентные строчки: https://leetcode.com/problems/count-the-number-of-consistent-strings/description/
4. Реализовать свой тип ошибки. Написать программу, в которой создается буферизованный канал ошибок `error`, положить в него созданную ошибку. Прочитать ошибку из канала, напечатать.

## III. Формат сдачи

- Самостоятельно проверить, что окружение (Go, VSCode, WSL) поставлено правильно, так как оно нужно для сдачи следующих домашних заданий.
- Задачи с leetcode можно решить прямо на leetcode (чтобы проверить, что все тесты прошли), и потом скопировать в https://go.dev/play/ (добавить main, чтобы Playground можно было запустить). Для каждой задачи ссылку на Go Playground с ее решением положить в таблицу.
