package main

import "testing"

func TestSoma(t *testing.T) {
	//arrange
	teste := soma(3, 2, 1)

	//act
	resultado := 6

	//assert
	if teste != resultado {
		t.Error("Valor esperado:", resultado, "Valor retornado:", teste)
	}
}

func TestSoma2(t *testing.T) {

	//arrange
	teste := soma(3, 2, 1)

	//act
	resultado := 7

	//assert
	if teste != resultado {
		t.Error("Valor esperado:", resultado, "Valor retornado:", teste)
	}
}

func TestMult(t *testing.T) {
	//arrange
	teste := multiplica(10, 10)

	//act
	resultado := 100

	//assert
	if teste != resultado {
		t.Error("Valor esperado:", resultado, "Valor retornado:", teste)
	}
}

func TestMult2(t *testing.T) {
	//arrange
	teste := multiplica(10, 10)

	//act
	resultado := 2560

	//assert
	if teste != resultado {
		t.Error("Valor esperado:", resultado, "Valor retornado:", teste)
	}
}
func TestSub(t *testing.T) {
	//arrange
	teste := subtrai(10, 5)

	//act
	resultado := -5

	//assert
	if teste != resultado {
		t.Error("Valor esperado:", resultado, "Valor retornado:", teste)
	}
}

func TestSub2(t *testing.T) {
	//arrange
	teste := subtrai(10, 10)

	//act
	resultado := 5

	//assert
	if teste != resultado {
		t.Error("Valor esperado:", resultado, "Valor retornado:", teste)
	}
}

func TestDiv(t *testing.T) {
	//arrange
	teste := divide(10)

	//act
	resultado := 10

	//assert
	if teste != resultado {
		t.Error("Valor esperado:", resultado, "Valor retornado:", teste)
	}
}

func TestDiv2(t *testing.T) {
	//arrange
	teste := divide(10)

	//act
	resultado := 5

	//assert
	if teste != resultado {
		t.Error("Valor esperado:", resultado, "Valor retornado:", teste)
	}
}
