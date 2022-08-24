package piscine

func ConvertBase(nbr, baseFrom, baseTo string) string {
	decimaForm := AtoiBase(nbr, baseFrom)
	return PrintNbrBase(decimaForm, baseTo)
}
