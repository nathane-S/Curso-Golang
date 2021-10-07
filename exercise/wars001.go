package kata

//Sua função recebe dois argumentos:
// 1. idade atual do pai (anos)
// 2. idade atual de seu filho (anos)
//Calcule quantos anos atrás o pai tinha duas vezes mais velho que seu filho (ou em quantos anos ele terá duas vezes mais).

func TwiceAsOld(dadYearsOld, sonYearsOld int) int {
	x := dadYearsOld - sonYearsOld*2
	if x < 0 {

		return -x

	}
	return x
}
