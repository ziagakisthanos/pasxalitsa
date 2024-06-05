package piscine

func UltimateDivMod(a, b *int) {
	div := *a / *b
	mod := *a % *b
	*a = div
	*b = mod
}
