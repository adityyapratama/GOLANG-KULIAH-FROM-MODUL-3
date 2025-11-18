package utils

import "testing"


func TesHashPassword(t *testing.T) {
	password := "rahasia123"

	hash, err := HashPassword(password)

	if err != nil{
		t.Error("HashPassword error : %v",err)
	}
	if len(hash) == 0 {
		t.Error(("Hash tidak boleh kosong"))
	}
}

func TestCheckPassword( t *testing.T){
	password := "rahasia123"
	hash, _ := HashPassword(password)

	if !CheckPassword(password, hash){
		t.Error("seharusnya password valid")
	}

	if CheckPassword("salah123", hash){
		t.Error("seharusnya password tidak valid")
	}
}
