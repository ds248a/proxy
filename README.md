# socks5
Персональный (однопользовательский) [Socks5](http://en.wikipedia.org/wiki/SOCKS) прокси сервер.

Реализована аутентификация по логину и паролю.

На данный момент атрибуты авторизации, задаваемые следующей структурой, `жёстко` прописывать в коде.
```go
type params struct {
	User            string
	Password        string
	Port            string
	Addr            string
	AllowedDestFqdn string
}
```