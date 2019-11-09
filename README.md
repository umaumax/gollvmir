# gollvmir

* extract function infomation from llvm ir

## how to install
```
go get -u github.com/umaumax/gollvmir
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

* if you are using llvm `v8.0.0`
  * you may rewrite code from `github.com/llir/llvm` to `github.com/umaumax/llvm`
    * [umaumax/llvm at v8\.0\.0]( https://github.com/umaumax/llvm/tree/v8.0.0 )
      * [Fix DIFlagNonTrivial to DIFlagTrivial · umaumax/llvm@45ef7b9]( https://github.com/umaumax/llvm/commit/45ef7b9c888ab826ee83f93f92355c01dd773479 )

### FYI
* [LLVM IR and Go GopherAcademy]( https://blog.gopheracademy.com/advent-2018/llvm-ir-and-go/ )
