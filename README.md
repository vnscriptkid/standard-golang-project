# canvas
A blank canvas for your web app, in Go.

This repository is used in the course [Build Cloud Apps in Go](https://www.golang.dk/courses/build-cloud-apps-in-go).

## QnA
- `go doc net/http.Server`
- Why need timeout values for `http.Server`?
- Why we need graceful shutdown? What happens if server is not gracefully shutdown?
- How load-balancer handles TLS cert?
- How load-balancer checks health endpoint?
- How come 2 package names inside one folder? `handlers` vs `handlers_test`
- Why structured logging instead of fmt log?