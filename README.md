# DST Manager

## 文件结构

- cmd
    - 里面就一个`main.go`
- internal
    - 一些会被`main.go`使用的代码
    - 泛用性不高只能在这个app里用的话写这
- pkg
    - 一些会被`main.go`使用的代码
    - 泛用性高的写这
- configs
    - app的设定存放在这
- testdata
    - 测试用文件
    - 里面可以放Steam文件, 饥荒服务端, 游戏存档文件夹
    - 注: 代码的测试放在各个代码同一个文件夹里, 文件名为`xxx_test.go`
- docs
    - 文档
