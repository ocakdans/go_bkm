package others

import (
	"fmt"
	"strconv"
)

func main() {

	ferr := new(Ferrari)
	ferr.Brand = "Ferrai"
	ferr.Modal = "F50"
	ferr.Color = "Red"
	ferr.Speed = 500
	ferr.isSpecial = true
	ferr.Price = 1.4
	//fmt.Println(ferr.Information())

	merc := &Mercedes{
		Car: Car{
			Brand: "Merceds",
			Modal: "CLK",
			Color: "Red",
			Speed: 400,
			Price: 0.2,
		},
	}
	//fmt.Println("\n")
	//fmt.Println(merc.Information())
	CarExecute(ferr)
	CarExecute(merc)

}

type Car struct {
	Brand string
	Modal string
	Color string
	Speed int
	Price float64
}

type SpecialProduction struct {
	isSpecial bool
}

func CarExecute(c CarFace) {
	fmt.Println("\n")
	fmt.Println("Araç Bilgi: " + "\n" + c.Information())
	fmt.Println("\n")
	isRun := c.Run()
	if isRun {
		fmt.Println("Araç çalışıyor")
	} else {
		fmt.Println("Araç çalışMıyor")
	}
	isStop := c.Stop()
	if isStop {
		fmt.Println("Araç durdu")
	} else {
		fmt.Println("Araç duramadı...")
	}

}

// Ferrrai
type Ferrari struct {
	Car // composition
	SpecialProduction
}

func (Ferrari) Run() bool {
	return true
}

func (Ferrari) Stop() bool {
	return true
}

func (f Ferrari) Information() string {
	info := "\t" + f.Brand + " " + f.Modal +
		"\n" + "\t" + "Color: " + f.Color + "\n" + "\t" +
		"Speed: " + strconv.Itoa(f.Speed) + "\n" + "\t" + "Price: " + strconv.FormatFloat(f.Price, 'g', 0, 64)
	addSpecial := "Yes"
	if f.isSpecial {
		info += "\n" + "\t" + "Special: " + addSpecial
	}
	return info
}

// Lamborghini
type Lamborghini struct {
	Car // composition
	SpecialProduction
}

func (Lamborghini) Run() bool {
	return true
}

func (Lamborghini) Stop() bool {
	return true
}

func (f Lamborghini) Information() string {
	info := "\t" + f.Brand + " " + f.Modal +
		"\n" + "\t" + "Color: " + f.Color + "\n" + "\t" +
		"Speed: " + strconv.Itoa(f.Speed) + "\n" + "\t" + "Price: " + strconv.FormatFloat(f.Price, 'g', 0, 64)
	addSpecial := "Yes"
	if f.isSpecial {
		info += "\n" + "\t" + "Special: " + addSpecial
	}
	return info
}

// Mercedes
type Mercedes struct {
	Car // composition
	SpecialProduction
}

func (Mercedes) Run() bool {
	return true
}

func (Mercedes) Stop() bool {
	return true
}

func (f Mercedes) Information() string {
	info := "\t" + f.Brand + " " + f.Modal +
		"\n" + "\t" + "Color: " + f.Color + "\n" + "\t" +
		"Speed: " + strconv.Itoa(f.Speed) + "\n" + "\t" + "Price: " + strconv.FormatFloat(f.Price, 'g', 0, 64)
	return info
}

type CarFace interface {
	Run() bool
	Stop() bool
	Information() string
}
