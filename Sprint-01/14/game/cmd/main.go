package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"os"
	"strconv"

	application "github.com/DmitriySolopenkov/go-yandex-lyceum/Sprint-01/14/game/internal"
)

func main() {
	ctx := context.Background()
	// Exit завершает программу с заданным кодом
	os.Exit(mainWithExitCode(ctx))
}

type World struct {
	Height int
	Width  int
	Cells  [][]bool
}

func NewWorld(height, width int) (*World, error) {
	if height <= 0 || width <= 0 {
		return nil, errors.New("Размеры должны быть > 0")
	}
	cells := make([][]bool, height)
	for i := range cells {
		cells[i] = make([]bool, width) // создаём новый слайс в каждой строке
	}
	return &World{
		Height: height,
		Width:  width,
		Cells:  cells,
	}, nil
}

func run() error {
	args := os.Args[1:]
	if len(args) < 3 {
		return fmt.Errorf("Usage: main <rows> <columns> <fill_percentage>")
	}

	rows, err := strconv.Atoi(args[0])
	if err != nil {
		return fmt.Errorf("Invalid number of rows: %v", args[0])
	}

	columns, err := strconv.Atoi(args[1])
	if err != nil {
		return fmt.Errorf("Invalid number of columns: %v", args[1])
	}

	percentage, err := strconv.Atoi(args[2])
	if err != nil {
		return fmt.Errorf("Invalid fill percentage: %v", args[2])
	}

	file, err := os.Create("config.txt")
	if err != nil {
		return fmt.Errorf("Failed to create config file: %v", err)
	}
	defer file.Close()

	_, err = fmt.Fprintf(file, "%dx%d %d%%", rows, columns, percentage)
	if err != nil {
		return fmt.Errorf("Failed to write to config file: %v", err)
	}

	return nil
}

func mainWithExitCode(ctx context.Context) int {
	cfg := application.Config{
		Width:  10,
		Height: 10,
	}
	app := application.New(cfg)
	// Запускаем приложение
	if err := app.Run(ctx); err != nil {
		switch {
		case errors.Is(err, context.Canceled):
			log.Println("Processing cancelled.")
		default:
			log.Println("Application run error", err)
		}
		// Возвращаем значение, не равное нулю, чтобы обозначить ошибку
		return 1
	}
	// Выход без ошибок
	return 0
}
