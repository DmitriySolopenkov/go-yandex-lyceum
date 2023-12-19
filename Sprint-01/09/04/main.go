package main

type Vehicle interface {
	CalculateTravelTime(distance float64) float64
	NewVehicle() string
}

type Car struct {
	Type     string //тип транспорта
	Speed    float64
	FuelType string
}

func (c Car) CalculateTravelTime(distance float64) float64 {
	return distance / c.Speed
}

func (c Car) NewVehicle() string {
	return c.Type
}

type Motorcycle struct {
	Type     string
	Speed    float64
	FuelType string
}

func (m Motorcycle) CalculateTravelTime(distance float64) float64 {
	return distance / m.Speed
}

func (m Motorcycle) NewVehicle() string {
	return m.Type
}

func EstimateTravelTime(vehicles []Vehicle, distance float64) map[string]float64 {
	travelTimes := make(map[string]float64)

	for _, vehicle := range vehicles {
		time := vehicle.CalculateTravelTime(distance)
		switch vehicle.(type) {
		case Car:
			travelTimes["main.Car"] = time
		case Motorcycle:
			travelTimes["main.Motorcycle"] = time
		}
	}

	return travelTimes
}

// func EstimateTravelTime(vehicles []Vehicle, distance float64) map[string]float64 {
// 	outVehicles := map[string]float64{}

// 	for _, vehicle := range vehicles {

// 		outVehicles[vehicle.NewVehicle()] = vehicle.CalculateTravelTime(distance)
// 		// fmt.Println(outVehicles[vehicle.NewVehicle()])
// 	}

// 	return outVehicles
// }

// func main() {
// 	car := Car{Type: "Седан", Speed: 60.0, FuelType: "Бензин"}
// 	motorcycle := Motorcycle{Type: "Мотоцикл", Speed: 80.0}

// 	vehicles := []Vehicle{car, motorcycle}
// 	distance := 200.0

// 	travelTimes := EstimateTravelTime(vehicles, distance)

// 	fmt.Println(travelTimes["Motorcycle"])
// 	fmt.Println(travelTimes["Car"])
// 	// for vehicle, time := range travelTimes {
// 	// 	fmt.Printf("%s: %.2f hours\n", vehicle, time)
// 	// }
// }

/*
Управление транспортом
Вам нужно создать систему управления транспортными средствами,
	такими как автомобили и мотоциклы.

Каждое транспортное средство должно иметь метод для расчета времени
	в пути до определенного пункта назначения.

Создайте интерфейс Vehicle, который будет представлять транспортное средство
	и иметь метод CalculateTravelTime(distance float64) float64
	для расчета времени в пути.

Реализуйте две структуры, Car (автомобиль) и Motorcycle (мотоцикл),
	обе должны реализовывать интерфейс Vehicle и иметь соответствующие поля
	для хранения данных о транспортных средствах
	(например, скорость и тип транспортного средства).

Создайте функцию
	EstimateTravelTime(vehicles []Vehicle, distance float64) map[string]float64,
	которая принимает список транспортных средств
	и расстояние до пункта назначения, а затем возвращает карту (map),
	где ключи - это типы транспортных средств,
	а значения - время в пути для каждого транспортного средства.
	Используйте полиморфизм для вызова метода CalculateTravelTime()
	на каждом транспортном средстве независимо от его типа.
*/
