//Intellias HW10
package main

import (
	"fmt"
	"math/rand"
	"time"
)

//cтворюємо контсанти по використанюю їжі тваринами
const (
	catsFoodPerKilo = 7.0
	dogsFoodPerKilo = 10.0 / 5.0
	cowsFoodPerKilo = 25.0
)

type cow struct {
	name   string
	weight float64
	milk   bool
	meat   bool
}

func (c cow) receiveFoodWeight() float64 {
	return c.weight * cowsFoodPerKilo
}

func (c cow) myName() string {
	return "корови " + c.name
}
func (c cow) receiveWeightOfAnimal() float64 {
	return c.weight
}

type cat struct {
	name       string
	weight     float64
	mouseCatch bool
}

func (c cat) receiveFoodWeight() float64 {
	return c.weight * catsFoodPerKilo
}

func (c cat) myName() string {
	return "кота " + c.name
}

func (c cat) receiveWeightOfAnimal() float64 {
	return c.weight
}

type dog struct {
	name     string
	weight   float64
	tailFlip bool
}

func (d dog) receiveFoodWeight() float64 {
	return d.weight * dogsFoodPerKilo
}
func (d dog) myName() string {
	return "собаки " + d.name
}

func (d dog) receiveWeightOfAnimal() float64 {
	return d.weight
}

type weightGetter interface {
	receiveFoodWeight() float64
	receiveWeightOfAnimal() float64
}

type nameGetter interface {
	myName() string
}

// виходячи з умови задачі можно було обійтись одним інтерфейсом, але якщо потенційно ми б розділили
//нажаль не вдалось перемогти сеттери(
type animals interface {
	nameGetter
	weightGetter
}

//makeSomeAnimal повертає екземпляр тварини(інтерфейсу) та текст помилки
func makeSomeAnimal(typeOfAnimal string, name string, weight float64) (animals, string) {
	var newAnimal animals
	if len(typeOfAnimal) == 0 || weight <= 0 {
		return newAnimal, "abnormal animal"
	}
	switch typeOfAnimal {
	case "dog":
		newAnimal = dog{tailFlip: false, weight: weight, name: name}
	case "cow":
		newAnimal = cow{milk: false, meat: true, weight: weight, name: name}
	case "cat":
		newAnimal = cat{mouseCatch: true, weight: weight, name: name}
	default:
		return newAnimal, "unknown animal"
	}
	return newAnimal, ""
}

//receiveInfo друкує інформацію про нашу ферму
func receiveInfo(farm *[]animals) {
	var total float64
	for _, t := range *farm {
		total += t.receiveFoodWeight()
		fmt.Printf("Вага корму на місяць для %s з власною вагою  %0.2f становить %.2f кг \n", t.myName(), t.receiveWeightOfAnimal(), t.receiveFoodWeight())
	}
	fmt.Println("---------")
	fmt.Printf("Загальна потреба в кормі для ферми становить: %.2f кг \n ", total)
}

func main() {
	//підготуємо тестові данні для заповнення, у випадку не рандомної ваги можна переробити на мапу
	rawCats := []string{"Мурзік", "Котик", "Мурчик"}
	rawDogs := []string{"Тотошка", "Пірат", "Рекс"}
	rawCows := []string{"Зорька", "Мілка"}

	// створюємо пустий слайс ферми
	farm := []animals{}

	//заповнюємо ферму котами
	for k, v := range rawCats {
		rand.Seed(time.Now().UnixNano())
		if newVal, err := makeSomeAnimal("cat", v, rand.Float64()*10.0+float64(k)); len(err) == 0 {
			farm = append(farm, newVal)
		}
	}
	//заповнюємо ферму собаками
	for k, v := range rawDogs {
		rand.Seed(time.Now().UnixNano())
		if newVal, err := makeSomeAnimal("dog", v, (rand.Float64()+float64(k))*10.0); len(err) == 0 {
			farm = append(farm, newVal)
		}
	}
	//заповнюємо ферму коровами
	for k, v := range rawCows {
		rand.Seed(time.Now().UnixNano())
		if newVal, err := makeSomeAnimal("cow", v, (rand.Float64()+float64(k))*100.0); len(err) == 0 {
			farm = append(farm, newVal)
		}
	}
	receiveInfo(&farm)
}
