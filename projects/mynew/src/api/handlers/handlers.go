package handlers

import (
	createUserHandler "mynew/src/api/handlers/create-user"
	rootHandler "mynew/src/api/handlers/root"
)

var (
	HandlerCreateUser = createUserHandler.HandlerCreateUser
	Root              = rootHandler.Root
)
