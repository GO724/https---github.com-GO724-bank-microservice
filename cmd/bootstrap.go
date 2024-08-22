package bootstrap

// Инициализация и запуск сервера

type Bootstrapper struct {
	// Вложенность как в пакетах
	conltroller
	database
}

func New() bootstrap {
	bootstrap

}
