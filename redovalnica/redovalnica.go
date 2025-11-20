package redovalnica

import "fmt"

//Predstavlja študenta z naslednjimi polji:
//   - Ime: Študentovo ime
//   - Priimek: Študentov priimek
//   - Ocene: Seznam ocen študenta
type Student struct {
	Ime     string
	Priimek string
	Ocene   []int
}

// DodajOceno doda oceno študentu z dano vpisno številko.
// Parametri:
//   - studenti: Mapa študentov, kjer so ključi vpisne številke
//   - vpisnaStevilka: Vpisna številka študenta, ki mu želimo dodati oceno
//   - ocena: Ocena, ki jo želimo dodati
//   - minOcena: Najnižja dovoljena vrednost ocene
//   - maxOcena: Najvišja dovoljena vrednost ocene
// Preveri veljavnost ocene in obstoj študenta.
// Če ocena ni v območju [minOcena, maxOcena], izpiše opozorilo.
// Če študent ne obstaja, izpiše ustrezno sporočilo.
func DodajOceno(studenti map[string]Student, vpisnaStevilka string, ocena int, minOcena int, maxOcena int) {
	if ocena < minOcena || ocena > maxOcena {
		fmt.Println("Ocena ni v ustreznem območju 0..10")
		return
	}
	student, ok := studenti[vpisnaStevilka]
	if !ok {
		fmt.Println("študent ne obstaja")
		return
	}
	student.Ocene = append(student.Ocene, ocena)
	studenti[vpisnaStevilka] = student
}

func povprecje(studenti map[string]Student, vpisnaStevilka string, minStOcen int) float64 {
	student, ok := studenti[vpisnaStevilka]
	if !ok {
		return -1.0
	}
	if len(student.Ocene) < minStOcen {
		return 0.0
	}
	return avg(student.Ocene)
}

func avg(num []int) float64 {
	sum := 0
	for _, v := range num {
		sum += v
	}
	return float64(sum) / float64(len(num))
}

func IzpisiKoncniUspeh(studenti map[string]Student, minStOcen int) {
	for k, v := range studenti {
		avgOcena := povprecje(studenti, k, minStOcen)
		var uspeh string
		switch {
		case avgOcena >= 9:
			uspeh = "Odličen študent!"
		case avgOcena >= 6:
			uspeh = "Povprečen študent"
		default:
			uspeh = "Neuspešen študent!"

		}
		fmt.Printf("%s %s: povpreča ocena %f -> %s \n", v.Ime, v.Priimek, avgOcena, uspeh)
	}
}

func IzpisRedovalnice(studenti map[string]Student) {
	for k, v := range studenti {

		fmt.Printf("%s - %s %s: ", k, v.Ime, v.Priimek)
		fmt.Println(v.Ocene)
	}
}
