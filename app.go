package main

import (
    "github.com/gocraft/web"
    _ "fmt"
    "net/http"
    "github.com/ahmadmuzakki29/helloapi/context"
)

func main() {
    ctx := context.Context{}
    router := web.New(ctx).
        Middleware(web.LoggerMiddleware).
        Middleware(web.ShowErrorsMiddleware)
    
    router.Get("/v1/products", ctx.ProductList)
    router.Get("/v1/productDetail/")

    http.ListenAndServe("localhost:3000", router)   // Start the server!
}
