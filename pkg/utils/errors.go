package utils

func CheckDBErr(err error) {
	if err != nil {
		panic(err)
	}
}

func CheckGeneralErr(err error) {
	if err != nil {
		panic(err)
	}
}
