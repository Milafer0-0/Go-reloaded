# Go Reloaded

## Description

**Go Reloaded** is a text processing tool written in Go. It reads an input text file, applies a series of transformations and formatting rules, and writes the processed text into an output file.

The program supports number conversion, text case modification, punctuation formatting, quotation formatting, and automatic correction of the indefinite article (`a` → `an`).

---

## Features

The program performs the following transformations:

### Number Conversion

* **Hexadecimal to Decimal**

  * `1E (hex)` → `30`

* **Binary to Decimal**

  * `10 (bin)` → `2`

---

### Text Case Conversion

* **Uppercase**

  * `hello (up)` → `HELLO`

* **Lowercase**

  * `HELLO (low)` → `hello`

* **Capitalize**

  * `hello (cap)` → `Hello`

The commands also support modifying multiple previous words:

* `(up, N)`
* `(low, N)`
* `(cap, N)`

Example:

```text
This is so exciting (up, 2)
```

becomes

```text
This is SO EXCITING
```

---

### Punctuation Formatting

The program correctly formats the following punctuation marks:

* `.`
* `,`
* `!`
* `?`
* `:`
* `;`

Rules:

* No space before punctuation.
* One space after punctuation.
* Consecutive punctuation groups such as `...` and `!?` are preserved.

Example:

```text
I was sitting over there ,and then BAMM !!
```

becomes

```text
I was sitting over there, and then BAMM!!
```

---

### Quotation Marks

Quotation marks are automatically attached to the quoted text.

Example:

```text
' awesome '
```

becomes

```text
'awesome'
```

Example with multiple words:

```text
' I am the most well-known homosexual in the world '
```

becomes

```text
'I am the most well-known homosexual in the world'
```

---

### Article Correction

The article **a** is automatically replaced with **an** when the following word begins with:

* a
* e
* i
* o
* u
* h

Example:

```text
A amazing rock
```

becomes

```text
An amazing rock
```

---

## Project Structure

```text
.
├── main.go
├── transform.go
├── sample.txt
├── result.txt
└── README.md
```

---

## Usage

Run the program using:

```bash
go run . sample.txt result.txt
```

The program:

1. Reads the contents of `sample.txt`.
2. Applies all supported transformations.
3. Writes the final result into `result.txt`.

---

## Example

### Input

```text
it (cap) was the best of times (up) , there was a amazing opportunity .
```

### Output

```text
It was the best of TIMES, there was an amazing opportunity.
```

---

## Requirements

* Go 1.20 or later
* Standard Go packages only

---

## Learning Objectives

This project demonstrates:

* File input/output
* String manipulation
* Number conversion
* Text parsing
* Error handling
* Working with slices
* Go standard library usage

# Go Reloaded

## Описание

**Go Reloaded** — это консольная утилита, написанная на языке **Go**, предназначенная для автоматической обработки и форматирования текста.

Программа считывает текст из входного файла, применяет к нему ряд преобразований и сохраняет результат в выходной файл.

---

# Возможности

Программа поддерживает следующие операции.

## Преобразование чисел

### Шестнадцатеричная система

Конструкция:

```text
<число> (hex)
```

заменяет шестнадцатеричное число на его десятичное представление.

**Пример**

```text
1E (hex)
```

↓

```text
30
```

---

### Двоичная система

Конструкция:

```text
<число> (bin)
```

заменяет двоичное число на десятичное.

**Пример**

```text
10 (bin)
```

↓

```text
2
```

---

# Изменение регистра

## Верхний регистр

Команда

```text
(up)
```

переводит предыдущее слово в верхний регистр.

**Пример**

```text
hello (up)
```

↓

```text
HELLO
```

---

## Нижний регистр

Команда

```text
(low)
```

переводит предыдущее слово в нижний регистр.

**Пример**

```text
HELLO (low)
```

↓

```text
hello
```

---

## Капитализация

Команда

```text
(cap)
```

делает первую букву слова заглавной.

**Пример**

```text
brooklyn (cap)
```

↓

```text
Brooklyn
```

---

## Обработка нескольких слов

Команды

```text
(up, N)
(low, N)
(cap, N)
```

применяются сразу к нескольким предыдущим словам.

**Пример**

```text
This is so exciting (up, 2)
```

↓

```text
This is SO EXCITING
```

---

# Форматирование пунктуации

Программа автоматически исправляет расположение следующих знаков препинания:

* `.`
* `,`
* `!`
* `?`
* `:`
* `;`

Правила:

* перед знаком препинания пробел отсутствует;
* после знака препинания ставится один пробел;
* группы знаков (`...`, `!?`, `!!` и т.д.) сохраняются.

**Пример**

```text
I was sitting over there ,and then BAMM !!
```

↓

```text
I was sitting over there, and then BAMM!!
```

---

# Обработка кавычек

Одинарные кавычки автоматически прижимаются к тексту внутри них.

**Пример**

```text
' awesome '
```

↓

```text
'awesome'
```

Если внутри находится несколько слов:

```text
' I am the most well-known homosexual in the world '
```

↓

```text
'I am the most well-known homosexual in the world'
```

---

# Исправление артиклей

Если после артикля **a** следует слово, начинающееся с:

* a
* e
* i
* o
* u
* h

то артикль автоматически заменяется на **an**.

**Пример**

```text
a amazing rock
```

↓

```text
an amazing rock
```

---

# Структура проекта

```text
.
├── main.go
├── transform.go
├── sample.txt
├── result.txt
└── README.md
```

---

# Использование

Запуск программы:

```bash
go run . sample.txt result.txt
```

После запуска программа:

1. Считывает содержимое файла `sample.txt`;
2. Выполняет все предусмотренные преобразования;
3. Записывает готовый результат в `result.txt`.

---

# Пример работы

## Входной файл

```text
it (cap) was the best of times (up) , there was a amazing opportunity .
```

## Выходной файл

```text
It was the best of TIMES, there was an amazing opportunity.
```

---

# Используемые технологии

* Go
* Стандартная библиотека Go (`fmt`, `os`, `strconv`, `strings`, `filepath`)

---

# Цель проекта

Проект позволяет закрепить навыки работы с:

* чтением и записью файлов;
* обработкой строк;
* преобразованием числовых систем счисления;
* работой со срезами;
* обработкой ошибок;
* парсингом текста;
* использованием стандартной библиотеки Go.

---
