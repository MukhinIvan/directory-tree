# Утилита tree.

Утилита tree выводит дерево каталогов и файлов (если указана опция `-f`).

Список папок-файлов отсортирован по алфавиту.

Папка `testdata` создана специально для тестов.

```
go run main.go testdata -f
├───file1.txt (25b)
├───folder1
│       ├───file2.txt (25b)
│       └───file3.txt (empty)
└───folder2
        ├───file4.txt (20b)
        └───folder3
                ├───folder4
                │       └───file6.txt (34b)
                └───folder5
                        └───file5.txt (25b)
```

```
go run main.go testdata
├───folder1
└───folder2
        └───folder3
                ├───folder4
                └───folder5
```

Запустить тесты: `go test -v`.

```
$ go test -v
=== RUN   TestTreeWithFiles
--- PASS: TestTreeWithFiles (0.00s)
=== RUN   TestTreeWithoutFiles
--- PASS: TestTreeWithoutFiles (0.00s)
PASS
ok      github.com/MukhinIvan/directory-tree    0.039s
```
