# go-dependency-injection

Jon Bodner's [Learning Go: An Idiomatic Approach to Real-World Go Programming](https://github.com/learning-go-book-2e).

Chapter 7: Types, Methods, and Interfaces | Implicit Interfaces Make Dependency Injection Easier

---

Clone the repository.

```sh
git clone https://github.com/chrisbradleydev/go-dependency-injection.git .
```

Build Docker image.

```sh
docker build -t go-dependency-injection .
```

Run Docker container.

```sh
docker run -p 8080:8080 -v $(pwd):/app -v /app/tmp go-dependency-injection
```
