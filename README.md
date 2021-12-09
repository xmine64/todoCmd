# TODO CMD APP! 🧙‍♂️

## Table of contents
* [General info](#general-info)
* [Update](#update)
* [Requirements](#requirements)
* [set-up](#setup)
* [usage](#usage)

## General info
todo-cmd is an app you can add your tasks , edit or delete them
it's all in your terminal !

## Update
fix some bugs

delete some ascII art 

## Requirements 

install Golang on your os

###windows:
download Golang latest version on https://go.dev

###Linux:
open your terminal and :
 
```sudo yum install golang```

## setup
clone the project:

```git clone https://github.com/dozheiny/todoCmd```

change directory to the project:

```cd  todoCmd```

build program :

```go build .```

run the program to build database and table and .log

```./todocmd```

you should see this message :

```
d888888P                dP  
   88                   88 
   88    .d8888b. .d888b88 .d8888b. 
   88    88'  '88 88'  '88 88'  '88
   88    88.  .88 88.  .88 88.  .88
   dP    '88888P' '88888P8 '88888P' 
   ooooooooooooooooooooooooooooooo // written by dozheiny

CONSOLE: Create database successfully
CONSOLE: Create Table successfully
CONSOLE: .log created!
```

it seems everything is right! 🦄

## Usage

type ```./todocmd -help``` to see help message
```
d888888P                dP
   88                   88
   88    .d8888b. .d888b88 .d8888b.
   88    88'  '88 88'  '88 88'  '88
   88    88.  .88 88.  .88 88.  .88
   dP    '88888P' '88888P8 '88888P'
   ooooooooooooooooooooooooooooooo // written by dozheiny

PS C:\Users\dozhe\go\src\TodoCmd> go run main.go -help
Usage of C:\Users\dozhe\AppData\Local\Temp\go-build1797785965\b001\exe\main.exe:
  -id int
        ID is your task id
  -object string
        object problaly is your task name (default "if you seen this, it means flag is not working")
  -status string
        status will decide what you want to do  (default "hi")
```

###ADD 👨‍💻:

for add task:

```
./todocmd -status add -object [your task]
```

example:
```
./todocmd -status add -object "Write readme.md"
```

### show 🤖
for see your tasks :
```
./todocmd -status show
```

### replace 🐲

for replace your tasks :
```
./todocmd -status replace -object [your task u want replace] -id [your task ID]
```

example

```
./todocmd -status replace -object "clean desktop" -id 3
```

### delete 🍵

for delete your tasks by id
```
./todocmd -status delete -id [id your task you want to delete]
```
example :
```
./todocmd -status delete -id 3
```

#### Hope enjoy well ! 🐋