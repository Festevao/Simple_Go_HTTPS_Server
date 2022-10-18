### Variaveis de ambiente.
```bash
PORT (defalt value = "443")
```
### Gerar certificados.
```bash
$ cd TLS
$ openssl req -x509 -nodes -newkey rsa:2048 -keyout server.key -out server.crt -days 3650
```