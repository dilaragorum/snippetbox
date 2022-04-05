Go projesine başlarken

`go mod init github.com/dilaragorum/snippetbox`
`go mod tidy`

# Web Applications Basic

**handler** (controller): They’re responsible for executing your application logic and for writing HTTP response headers and bodies.

**router**: (servemux): This stores a mapping between the URL patterns for your
application and the corresponding handlers.

| URL | Handler |
| /users | GetUsers |
| /user/1 | GetUser |

**Web server**: One of the great things about Go is that you can establish a web server
and listen for incoming requests as part of your application itself. (Nginx or Apache)

go run cmd/web/\*
