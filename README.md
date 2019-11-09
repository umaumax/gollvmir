# gollvmir

* extract function infomation from llvm ir

## how to install
```
go get -u github.com/umaumax/gollvmif
```

## how to use
```
cat main.ll | gollvmir --format='csv'
gollvmir --format='tsv' main.ll
gollvmir --format='json' main.ll
```

## WARN
* add `-g` option to include function information
```
clang++ -g -emit-llvm main.cpp -c -o main.bc
clang++ -g -emit-llvm main.cpp -c -S -o main.ll
```


### FYI
* [LLVM IR and Go GopherAcademy]( https://blog.gopheracademy.com/advent-2018/llvm-ir-and-go/ )
