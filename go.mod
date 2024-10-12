module github.com/yoyo1025/persona-api

go 1.19

require (
	github.com/lib/pq v1.10.9
	github.com/sashabaranov/go-openai v1.30.3
)

require (
	github.com/labstack/echo/v4 v4.12.0 // indirect
	github.com/labstack/gommon v0.4.2 // indirect
	github.com/mattn/go-colorable v0.1.13 // indirect
	github.com/mattn/go-isatty v0.0.20 // indirect
	github.com/valyala/bytebufferpool v1.0.0 // indirect
	github.com/valyala/fasttemplate v1.2.2 // indirect
	golang.org/x/crypto v0.22.0 // indirect
	golang.org/x/net v0.24.0 // indirect
	golang.org/x/sys v0.19.0 // indirect
	golang.org/x/text v0.14.0 // indirect
)

replace github.com/yoyo1025/persona-api => ./ // ローカル開発用
