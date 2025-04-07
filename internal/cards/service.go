package cards

func NewService() Service { // Uma fábrica que cria o service e retorna ele como interface
	return &serviceImpl{}
}

type serviceImpl struct{} // A implementação concreta desse service

type Service interface { // Define o que o service deve fazer (contrato)
	GetBySet(set string) []Card
}

// A Magia esta aqui, em GO ele consegue ver que usou o metodo Speak e a struct e interface se liga:
func (s *serviceImpl) GetBySet(set string) []Card {
	var filtered []Card
	for _, card := range MockCards {
		if card.Set == set {
			filtered = append(filtered, card)
		}
	}
	return filtered
}
