package cards

import "go.uber.org/fx"

// Module -> agrupa todas as dependências do domínio "cards"
var Module = fx.Module("cards",
	fx.Provide(
		NewService, // lógica de negócio
		NewHandler, // ponte entre HTTP e o service
	),
)
