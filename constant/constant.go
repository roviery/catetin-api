package constant

import (
	"time"

	"github.com/golang-jwt/jwt/v4"
)

const (
	MongoDatabaseName      = "CatetinDB"
	CollectionUsers        = "users"
	CollectionDeadlines    = "deadlines"
	CollectionQuicknotes   = "quicknotes"
	CollectionFinances     = "finances"
	CollectionTransactions = "transactions"
	CollectionTodos        = "todos"
)

var JWTExpiredTime = time.Duration(1) * time.Hour
var JWTSecretKey = []byte("catetin-secret-key")
var JWTSigningMethod = jwt.SigningMethodHS256
