package health

import "go.uber.org/fx"

// Module -> agrupa todas as dependências do domínio "health"
var Module = fx.Module("health",
	fx.Provide(
		NewHandler, // ponte entre HTTP e o service
	),
)
