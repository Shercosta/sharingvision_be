# How to run

> by Shercosta

1. This project is using go, so make sure to have go installed
2. run go mod tidy
3. feel free to modify the `.env.local` (add your database stuff in there, and yes do not forget to have a fresh database ready for this, and the port for the api to run in)
4. get the postman collection .json in the root, import it to postman
5. run `go run main.go`

if things go well, you will automatically run the migrations needed, and then you can try the requests in the postman collection.

_No Agentic AI was used, i dont even use Claude (or vibe code for that matter). All code was templated from my previous project._
