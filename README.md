# proxy
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

## Установка на сервер
1. для компиляции выполнить скрипт `build`
2. на сервер в директорию `/etc/systemd/system` добавить файл `proxy.service`. В файле заменить 
[USER_NAME] и [USEP_GROUP] на реквизиты не привилегированного пользователя
3. исполняемый файл разместить в директории `/usr/local/bin`
4. выполнить
```go
systemctl enable proxy
systemctl start proxy
systemctl status proxy
```
