package config

// Put in this list the methods that do not require authentication
var UnprotectedMethods = []string{
	"/api.gen.ServerInfoService/Info",
	"/api.gen.UserService/Login",
}
