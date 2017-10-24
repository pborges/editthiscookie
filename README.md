#Load cookies from EditThisCookie
Export cookies from the chrome plugin EditThisCookie to be loaded into a HttpClient

```
u, err := url.Parse("https://www.reddit.com")
if err != nil {
    panic(err)
}

f, err := os.Open("cookie.txt")
if err != nil {
    panic(err)
}

err = editthiscookie.Load(http.DefaultClient, f, u)
if err != nil {
    panic(err)
}
```