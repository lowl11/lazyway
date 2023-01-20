# lazy-gateway
> Library for making gateway requests

## Example
```go
gateway := lazyway.Echo(server)

healthRouter := lazyway.Route("/health").SetHosts("http://host:port")

gateway.Route(healthRouter)
```