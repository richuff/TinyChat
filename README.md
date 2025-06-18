# RcAdmin


![GitHub language count](https://img.shields.io/github/languages/count/richuff/TinyChat)      ![GitHub commit activity](https://img.shields.io/github/commit-activity/w/richuff/TinyChat)

## Introduce

This a project written by gin and gorm

## Installation

> you need go(1.22.3 )
>
> go mod tidy

## Quick Start

> go build

``` go
// main.go
func main() {
	utils.InitConfig()
	utils.InitMysql()
	utils.InitRedis()

	r := router.Router()
	err := r.Run("localhost:8080")
	if err != nil {
		return
	}
}
```

## License

Tiny-Chat is [MIT LICENSE](https://github.com/richuff/TinyChat/blob/master/LICENSE)