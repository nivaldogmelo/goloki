# Build

```
docker build -f docker/Dockerfile -t goloki .
```

# Config

Acertar configurações do Loki: ir em `pkg/goloki/goloki.go` e modificar o
```go
	hook := lokirus.NewLokiHookWithOpts(
		"http://localhost:3100",
```

Acertar tempo de intervalo entre os logs: ir em `Dockerfile` e trocar o entrypoint

```
# O 1 significa para gerar log a cada 1s, pode colocar 0.500 para 500ms ou 0.001 para 1ms, o quanto quiser
CMD ["/goloki", "gen", "1"]
```


Aumentar a quantidade de campos no log: ir em `pkg/goloki/goloki.go` e modificar o
```go
		logger.WithField("fizz", "buzz").Warnln("warning")
```
Essa linha significa que vai gerar um `{"fizz":"buzz"}` no log, se quiser adicionar mais campos é só duplicar a linha


Mudar os campos de identificação da app: ir em `pkg/goloki/goloki.go` e modificar o
```go
		WithStaticLabels(lokirus.Labels{
			"app":         "goloki",
			"environment": "development",
		}).
```

pode só adicionar aí os campos que deseja
