import (
    "net/http"
    "io/ioutil"
)

func llamarAPI() {
    resp, err := http.Get("https://api.example.com/datos")
    if err != nil {
        panic(err)
    }
    defer resp.Body.Close()

    body, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        panic(err)
    }
    fmt.Println(string(body)) // Imprime la respuesta de la API
}
