# ccwc

`ccwc` is a Go implementation of the Unix `wc` (word count) utility.  
It can count characters, words, lines, and bytes in a file.

## Features

- Supports counting **lines, words, characters, and bytes**.  
- Can process **multiple files** at once.  
- If no file is specified, reads from **standard input**.  
- Provides counts **separately or together**.  
- Lightweight, fast, and works **cross-platform**.

## Installation

Clone the repository and build the program with Go:

```bash
go build -o ccwc
```

## Usage

```bash
./ccwc -c file.txt   # count bytes
./ccwc -m file.txt   # count characters
./ccwc -w file.txt   # count words
./ccwc -l file.txt   # count lines
```

## Examples

```bash
./ccwc -c test.txt
1234 test.txt

./ccwc -m test.txt
1200 test.txt

./ccwc -l test.txt
56 test.txt

./ccwc -w test.txt
200 test.txt
```