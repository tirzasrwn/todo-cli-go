# todo-cli-go

## About

Todo app CLI is means for learning the language to the core. This is only using library for SQLite driver. Everything else is using standard library.

# Requirement

- Go

# Installation

```sh
go mod tidy
go build -o ./todo ./cmd/app
./todo --help
```

## Usage

- Show help message

```sh
$ ./todo --help

  usage: todo [options] command

  commands:
    readall   Show all tasks
    create    Create new task
    read      Show task by id
    update    Update content task by id
    delete    Delete task by id
    toggle    Toggle done task by id
    help      Show this help message

  options:
    -h, --help Show this help message

```

- Show all tasks

```sh
$ ./todo readall
--> 2023/07/21 08:03:55 readall
------------------------------------------------------------------------------------------------------------------------
|ID |CONTENT                       |DONE |CREATED_AT               |UPDATED_AT               |DONE_AT                  |
------------------------------------------------------------------------------------------------------------------------
|1  |continue read 1984            |true |2023-07-21T07:50:35+07:00|2023-07-21T08:03:54+07:00|2023-07-21T08:03:54+07:00|
|2  |watch anime                   |false|2023-07-21T07:51:02+07:00|2023-07-21T07:51:02+07:00|NULL                     |
|3  |fix todo bug sql query        |true |2023-07-21T07:51:34+07:00|2023-07-21T08:02:54+07:00|2023-07-21T08:02:54+07:00|
|4  |learn builder design pattern  |false|2023-07-21T07:52:39+07:00|2023-07-21T07:52:39+07:00|NULL                     |
------------------------------------------------------------------------------------------------------------------------
--> 2023/07/21 08:03:55 command readall success
```

- Show task by id

```sh
$ ./todo read 3
--> 2023/07/21 07:59:57 read
------------------------------------------------------------------------------------------------------------------------
|ID |CONTENT                       |DONE |CREATED_AT               |UPDATED_AT               |DONE_AT                  |
------------------------------------------------------------------------------------------------------------------------
|3  |fix todo bug                  |false|2023-07-21T07:51:34+07:00|2023-07-21T07:51:34+07:00|NULL                     |
------------------------------------------------------------------------------------------------------------------------
--> 2023/07/21 07:59:57 command read success
```

- Update content task by id

```sh
$ ./todo update 3 "fix todo bug sql query"
--> 2023/07/21 08:01:16 update
--> 2023/07/21 08:01:16 command update success
```

- Toggle done task by id

```sh
$ ./todo toggle 3
--> 2023/07/21 08:02:54 toggle
--> 2023/07/21 08:02:54 command toggle success
```

- Delete task by id

```sh
$ ./todo delete 5
--> 2023/07/21 08:05:36 delete
--> 2023/07/21 08:05:36 command delete success
```

- Show done task

```sh
$ ./todo done
--> 2023/07/21 10:38:30 done
------------------------------------------------------------------------------------------------------------------------
|ID |CONTENT                       |DONE |CREATED_AT               |UPDATED_AT               |DONE_AT                  |
------------------------------------------------------------------------------------------------------------------------
|1  |continue read 1984            |true |2023-07-21T07:50:35+07:00|2023-07-21T08:03:54+07:00|2023-07-21T08:03:54+07:00|
|3  |fix todo bug sql query        |true |2023-07-21T07:51:34+07:00|2023-07-21T08:02:54+07:00|2023-07-21T08:02:54+07:00|
------------------------------------------------------------------------------------------------------------------------
--> 2023/07/21 10:38:30 command done success
```

- Show undone task

```sh
$ ./todo undone
--> 2023/07/21 11:03:10 undone
------------------------------------------------------------------------------------------------------------------------
|ID |CONTENT                       |DONE |CREATED_AT               |UPDATED_AT               |DONE_AT                  |
------------------------------------------------------------------------------------------------------------------------
|2  |watch anime                   |false|2023-07-21T07:51:02+07:00|2023-07-21T07:51:02+07:00|NULL                     |
|4  |learn builder design pattern  |false|2023-07-21T07:52:39+07:00|2023-07-21T07:52:39+07:00|NULL                     |
------------------------------------------------------------------------------------------------------------------------
--> 2023/07/21 11:03:10 command undone success
```
