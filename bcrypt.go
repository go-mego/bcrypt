package bcrypt

import (
	"github.com/go-mego/mego"
	"golang.org/x/crypto/bcrypt"
)

// New 會建立一個 Bcrypt 雜湊演算法模組可供安插於 Mego 引擎中。
func New(opts ...*Options) mego.HandlerFunc {
	var o *Options
	if len(opts) > 0 {
		o = opts[0]
	} else {
		o = &Options{
			Cost: bcrypt.DefaultCost,
		}
	}
	return func(c *mego.Context) {
		c.Map(&Crypt{
			options: o,
		})
	}
}

// Options 是雜湊演算的選項。
type Options struct {
	// Cost 是演算的花費次數，越高則越安全但會消耗更多 CPU 資源。
	// 預設為 `10`（最低為 `4`、最高為 `31`）。
	Cost int
}

// Crypt 是雜湊演算的主要模組。
type Crypt struct {
	// options 是演算的選項。
	options *Options
}

// Hash 能透過 Bcrypt 雜湊一段純文字並供之後以 `Compare` 比對。
func (c *Crypt) Hash(source string) string {
	result, err := bcrypt.GenerateFromPassword([]byte(source), c.options.Cost)
	if err != nil {
		panic(err)
	}
	return string(result)
}

// Compare 可以比對雜湊結果跟未雜湊字串是否兩者相符，這可以用在比對密碼上。
func (c *Crypt) Compare(hashed string, plain string) bool {
	return bcrypt.CompareHashAndPassword([]byte(hashed), []byte(plain)) == nil
}
