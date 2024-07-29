package slices

import "fmt"

func Slice() {

	names := make([]string, 3) //slice set etmenin yöntemi
	//fmt.Println(names)
	names[0] = "Selim"
	names[1] = "Feyza"
	names[2] = "Beyza"
	//names[3] = "Berk" //yine hata veriyor.
	names = append(names, "Berk") //yeni değer ekleme yöntemi
	fmt.Println(names)

	cities := []string{"Ankara", "İstanbul", "İzmir"} // slice set etmenin diğer yöntemi.
	fmt.Println(cities)

	copyOfCities := make([]string, len(cities))
	fmt.Println(copyOfCities)
	copy(copyOfCities, cities) // destination from source
	fmt.Println(copyOfCities)
	cities = append(cities, "Adana")
	fmt.Println(cities)       // [Ankara İstanbul İzmir Adana]
	fmt.Println(copyOfCities) //[Ankara İstanbul İzmir]

	fmt.Println(cities[1:3]) // [İstanbul İzmir]
	fmt.Println(cities[:2])  // [Ankara İstanbul]

	primes := [6]int{2, 3, 5, 7, 11, 13}
	var myslice []int = primes[1:4]
	// slice olunca [] arasında sayı belirtmek zorunda kalmadık.
	// slice veri barındırmaz, sadece arka planda bir dizinin bir kısmını biteler. point eder.
	// slice ı değiştirmek yani değerini değiştirmek, array in de değerini değiştirmek anlamına gelir.

	fmt.Println(myslice)

}
