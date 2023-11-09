# print-value

![GitHub watchers](https://img.shields.io/github/watchers/XdpCs/print-value?style=social)
![GitHub stars](https://img.shields.io/github/stars/XdpCs/print-value?style=social)
![GitHub forks](https://img.shields.io/github/forks/XdpCs/print-value?style=social)
![GitHub last commit](https://img.shields.io/github/last-commit/XdpCs/print-value?style=flat-square)
![GitHub repo size](https://img.shields.io/github/repo-size/XdpCs/print-value?style=flat-square)
![GitHub license](https://img.shields.io/github/license/XdpCs/print-value?style=flat-square)

Print Golang's struct ptr's value

## Why I write this tool

When I use fmt.Print-related functions to output a struct pointer, it prints the address. At this point, you can choose
to implement the String() method, but I'm too lazy, so I wrote this tool.

## install

`go get`

```shell
go get -u github.com/XdpCs/print-value@master
```

`go mod`

```shell
require github.com/XdpCs/print-value latest
```

