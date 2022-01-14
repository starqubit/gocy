package common
/*
uuid
"github.com/satori/go.uuid"
*/ 

import "github.com/satori/go.uuid"

// uuidv1
func UuidV1() string{
	return uuid.NewV1().String()
}

// uuidv4
func UuidV4() string{
	return uuid.NewV4().String()
}
