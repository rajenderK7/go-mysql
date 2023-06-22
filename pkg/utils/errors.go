package utils

func CheckDBError(err error) {
	if err != nil {
		panic(err)
	}
}
