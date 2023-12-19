package main

func main() {
	// log.Println("Hello, World!111")

	// // Открываем файл
	// file, err := os.OpenFile("./log.txt", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	// if err != nil {
	// 	return
	// }
	// // Закрываем файл после выхода из main
	// defer file.Close()
	// // Конфигурируем логгер, чтобы он выводил лог в файл
	// log.SetOutput(file)

	// log.Println("Hello, World!222")
	// /*
	// 	"log.txt": строка с именем файла, который вы хотите открыть или создать.
	// 	os.O_CREATE|os.O_WRONLY|os.O_APPEND: комбинация флагов, которые определяют, как будет открыт файл:
	// 	os.O_CREATE: указывает, что файл должен быть создан, если его нет. Если файл существует, этот флаг ничего не делает.
	// 	os.O_WRONLY: файл будет открыт только для записи (write-only). Вы не сможете читать из этого файла внутри программы.
	// 	os.O_APPEND: данные будут добавляться в конец файла, а не перезаписывать его содержимое.
	// 	0644: восьмеричное число, которое даёт права доступа к файлу. 0644 означает, что файл будет доступен для чтения и записи владельцу файла, а остальным — только для чтения.
	// */

	// log.SetOutput(os.Stderr)
	// log.Println("[ERROR] Something went wrong!")

	// log := logrus.New()
	// log.Info("Hello, World!")

	// logger, _ := zap.NewProduction()
	// defer logger.Sync()

	// logger.Info("Hello, World!")

	// logger, err := zap.NewDevelopment()
	// if err != nil {
	// 	log.Fatalf("can't initialize zap logger: %v", err)
	// }
	// defer logger.Sync()

	// logger.Info("Hello, world!")

	// logger, err := zap.NewProduction()
	// if err != nil {
	// 	log.Fatalf("can't initialize zap logger: %v", err)
	// }
	// defer logger.Sync()

	// logger.Info("Hello, world!")
	// /*
	// 	Конфигурация	Печатает stack trace для	Печатает файл/номер строки	Печатает уровень сообщения в	Печатает время в формате
	// 	Development	Warn и выше	всегда	верхнем регистре	ISO8601 с миллисекундами
	// 	Production	Error/DPanic levels	всегда	нижнем регистре	epoch
	// */

	// // Создаём конфигурацию для логгера
	// logger, err := zap.NewProduction()
	// if err != nil {
	// 	log.Fatalf("can't initialize zap logger: %v", err)
	// }
	// defer logger.Sync()

	// // Добавляем поля
	// logger.With(zap.String("user", "JohnDoe"), zap.Int("age", 30)).
	// 	Info("User login")

	// // Другой пример с полями
	// logger.With(zap.String("city", "New York"), zap.Float64("temperature", 78.5)).
	// 	Warn("Weather update")

	// // Логируем, не добавляя поля
	// logger.Info("Simple log message")

	// logger, _ := zap.NewProduction()
	// defer logger.Sync() // flushes buffer, if any
	// // Вызываем SugaredLogger
	// sugar := logger.Sugar()
	// sugar.Infow("failed to fetch URL",
	// 	// Можем прокидывать любые типы и не указывать их
	// 	"url", "какой-то url",
	// 	"attempt", 3,
	// 	"backoff", time.Second,
	// )
	// sugar.Infof("Failed to fetch URL: %s", "какой-то url")

	// // По умолчанию текстовый формат лога
	// slog.Debug("Debug message") // 2023/09/12 17:29:42 INFO Debug message

	// // Но можно и JSON
	// logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	// logger.Info("Info message") // {"time":"2023-09-12T17:29:42.64891153+03:00","level":"INFO","msg":"Info message"}

	// logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	// logger.Info("finished",
	// 	slog.Group("req",
	// 		slog.String("method", r.Method),
	// 		slog.String("url", r.URL.String())),
	// 	slog.Int("status", http.StatusOK),
	// 	slog.Duration("duration", time.Second))
	// // {"time":"2009-11-10T23:00:00Z","level":"INFO","msg":"finished","req":{"method":"GET","url":"localhost"},"status":200,"duration":1000000000}

}
