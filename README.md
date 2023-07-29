# maproulette-go
go interface to maproulette api

early work

[docs](docs.md)

Example usage 

```go
apiKey := os.Getenv("MAPROULETTE_API_KEY")

if apiKey == "" {
    log.Fatal("Environment variable MAPROULETTE_API_KEY not set")
}
mr := maproulette.NewMapRouletteClient(&maproulette.MapRouletteClientOptions{
    APIKey: apiKey,
})

challenge, err := mr.GetChallenge(20202)
```