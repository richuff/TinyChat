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

	r := router.Router()
	r.Run("localhost:8080")
}
```

## License

ichu_vueapp is [MIT LICENSE](https://github.com/richu94/richu_vueapp/blob/master/LICENSE)