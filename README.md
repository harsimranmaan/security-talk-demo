# Demo for SQL injection and XSS
Made with sqlite3 and go 1.9
## Run server
```bash
go run main.go
```

### Browse to http://localhost:8500/ 
Play around with normal text inputs and use both secure and non-secure button to submit respective inputs.

## Try the following test input in insecure text box and hit insecure button

### Sql Injection
123';DELETE FROM users where uid =18;

## Try the following test input in insecure text box:
### XSS
<script>alert("You need to fix the security")</script>
