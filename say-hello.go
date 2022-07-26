package gosayhello

func SayHello(firstname, lastname string) string {
	return "This " + "is " + firstname + lastname
	// kalau major upgrade seperti ini, ada baiknya nama module di gomod nya diubah misal ditambahkan v2 diakhir nama modul, walaupun git repo nya sama
}
