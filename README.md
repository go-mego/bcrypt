# Bcrypt [![GoDoc](https://godoc.org/github.com/go-mego/brypt?status.svg)](https://godoc.org/github.com/go-mego/brypt)

Bcrypt 是用以取代傳統 MD5 和 SHA1 的演算法，用以演算並產生更安全的密碼。

# 索引

* [安裝方式](#安裝方式)
* [使用方式](#使用方式)
    * [演算設置](#演算設置)

# 安裝方式

打開終端機並且透過 `go get` 安裝此套件即可。

```bash
$ go get github.com/go-mego/bcrypt
```

# 使用方式

透過 `Encrypt` 可以透過 Bcrypt 加密一段文字。透過 `Compare` 可以比對加密結果跟未加密字串是否兩者相符，這可以用在比對密碼上。

```go
package main

import (
	"fmt"

	"github.com/go-mego/bcrypt"
)

func main() {
	// 透過 `Encrypt` 加密一段純文字串。
	hashed := bcrypt.Encrypt("myPassword")
	// 以 `Compare` 來確認輸入的純文字是否和加密的結果相符。
	if bcrypt.Compare(hashed, "myPassword") {
		fmt.Println("密碼相符！")
	}
}
```

## 演算設置

在 `Encrypt` 的第二個參數可以傳入 Bcrypt 的演算設置。

```go
bcrypt.Encrypt("myPassword", bcrypt.Option{
    // 演算的花費次數，越高則越安全但會消耗更多 CPU 資源。
    // 預設為 `10`（最低為 `4`、最高為 `31`）。
    Cost: 10,
})
```