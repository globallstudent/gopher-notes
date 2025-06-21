package internal

type User struct {
	Email string
	Role  string
}

// In-memory user store for demo purposes
var users = map[string]User{
	"admin@example.com": {Email: "admin@example.com", Role: "admin"},
	"user@example.com":  {Email: "user@example.com", Role: "user"},
}

func GetUserByEmail(email string) (User, bool) {
	user, ok := users[email]
	return user, ok
}
