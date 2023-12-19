package main

var knightIsAwake = false
var archerIsAwake = true
var prisonerIsAwake = false
var petDogIsPresent = false
// Output: false

// PodeAtaqueRápido só pode ser executada quando o cavaleiro está dormindo,
// pois ele fica vulnerável enquanto coloca a armadura.
func PodeAtaqueRápido(cavaleiroEstáAcordado bool) bool {
	return !cavaleiroEstáAcordado
}

// PodeEspionar pode ser executada se pelo menos um dos personagens estiver acordado.
func PodeEspionar(cavaleiroEstáAcordado, arqueiroEstáAcordado, prisioneiroEstáAcordado bool) bool {
	return cavaleiroEstáAcordado || arqueiroEstáAcordado || prisioneiroEstáAcordado
}

// PodeSinalizarPrisioneiro pode ser executada se o arqueiro estiver acordado e o prisioneiro estiver acordado.
func PodeSinalizarPrisioneiro(arqueiroEstáAcordado, prisioneiroEstáAcordado bool) bool {
	return arqueiroEstáAcordado && prisioneiroEstáAcordado
}

// PodeLiberarPrisioneiro pode ser executada se o prisioneiro estiver acordado e os outros 2 personagens estiverem dormindo
// ou se o cachorro de Annalyn estiver com ela e o arqueiro estiver dormindo.
func PodeLiberarPrisioneiro(cavaleiroEstáAcordado, arqueiroEstáAcordado, prisioneiroEstáAcordado, cachorroDeAnnalynEstáPresente bool) bool {
	if cachorroDeAnnalynEstáPresente {
		return !arqueiroEstáAcordado
	}
	return prisioneiroEstáAcordado && !cavaleiroEstáAcordado && !arqueiroEstáAcordado
}

