### Variaveis de ambiente.
```bash
PORT (defalt value = "443")
```
### Gerar certificados.
```bash
$ openssl req -x509 -nodes -newkey rsa:2048 -keyout server.rsa.key -out server.rsa.crt -days 3650
```