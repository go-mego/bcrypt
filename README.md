# Bcrypt [![GoDoc](https://godoc.org/github.com/go-mego/bcrypt?status.svg)](https://godoc.org/github.com/go-mego/bcrypt) [![Coverage Status](https://coveralls.io/repos/github/go-mego/bcrypt/badge.svg?branch=master)](https://coveralls.io/github/go-mego/bcrypt?branch=master) [![Build Status](https://travis-ci.org/go-mego/bcrypt.svg?branch=master)](https://travis-ci.org/go-mego/bcrypt) [![Go Report Card](https://goreportcard.com/badge/github.com/go-mego/bcrypt)](https://goreportcard.com/report/github.com/go-mego/bcrypt)

Bcrypt 是用以取代傳統 MD5 和 SHA1 的演算法，用以演算並產生更安全的密碼。

# 索引

* [安裝方式](#安裝方式)
* [使用方式](#使用方式)
    * [演算配置](#演算配置)
	* [雜湊字串](#雜湊字串)
	* [比對雜湊](#比對雜湊)

# 安裝方式

打開終端機並且透過 `go get` 安裝此套件即可。

```bash
$ go get github.com/go-mego/bcrypt
```

# 使用方式

將 `bcrypt.New` 傳入 Mego 引擎的 `New` 就能夠作為全域中介軟體，在不同的路由中使用相關的雜湊演算法函式。

```go
package main

import (
	"github.com/go-mego/bcrypt"
)

func main() {
	m := mego.New()
	// 將 Bcrypt 作為全域中介軟體即能在不同路由中使用。
	m.Use(bcrypt.New())
	m.Run()
}
```

Bcrypt 也能夠僅用於單個路由。

```go
func main() {
	m := mego.New()
	// 也可以僅將 Bcrypt 傳入單個路由中使用。
	m.GET("/", bcrypt.New(), func(c *bcrypt.Crypt) {
		// ...
	})
	m.Run()
}
```

## 演算配置

在建立 Bcrypt 中介軟體的時候，可以傳入一個 `&bcrypt.Options` 來設置雜湊演算法。

```go
func main() {
	m := mego.New()
	m.Use(bcrypt.New(&bcrypt.Options{
		// 演算的花費次數，越高則越安全但會消耗更多 CPU 資源。
		// 預設為 `10`（最低為 `4`、最高為 `31`）。
		Cost: 10,
	}))
	m.Run()
}
```

## 雜湊字串

以 `Hash` 來透過 Bcrypt 雜湊一段純文字。

```go
func main() {
	m := mego.New()
	m.GET("/", bcrypt.New(), func(c *bcrypt.Crypt) string {
		// 透過 `Hash` 雜湊一段純文字串。
		return c.Hash("myPassword") // 結果：$2a$12$yEbegYKik2I4UdWpUvafsu...
	})
	m.Run()
}
```

## 比對雜湊

透過 `Compare` 比對雜湊結果跟未雜湊字串是否兩者相符，這可以用在比對密碼上。

```go
func main() {
	m := mego.New()
	m.GET("/", bcrypt.New(), func(c *bcrypt.Crypt) string {
		hashed := c.Hash("myPassword")
		// 以 `Compare` 來確認輸入的純文字是否和加密的結果相符。
		if c.Compare(hashed, "myPassword") {
			return "密碼相符！"
		} else {
			return "密碼不對啦！"
		}
	})
	m.Run()
}
```