package main

import (
	"context"
	"log"
	"os"

	"github.com/urfave/cli/v3"
	red "github.com/uselesnik/redovalnica/redovalnica"
)

func main() {
	cmd := &cli.Command{
		Name:  "redovalnica",
		Usage: "redovalnica ki hrani posodablja briše in izpisuje Ocene",
		Flags: []cli.Flag{
			&cli.IntFlag{
				Name:  "minOcena",
				Usage: "Minimalna vrednost Ocene ki jo hranimo v redovalnico",
				Value: 0,
			},
			&cli.IntFlag{
				Name:  "maxOcena",
				Usage: "Maksimalna vrednost Ocene ki jo hranimo v redovalnico",
				Value: 10,
			},
			&cli.IntFlag{
				Name:  "minStOcen",
				Usage: "Minimalno število ocen potrebnih za pozitivno zaključeno oceno",
				Value: 6,
			},
		},
		Action: func(ctx context.Context, c *cli.Command) error {
			minOcena := c.Int("minOcena")
			maxOcena := c.Int("maxOcena")
			minStOcen := c.Int("minStOcen")
			test(minOcena, maxOcena, minStOcen)
			return nil
		},
	}
	if err := cmd.Run(context.Background(), os.Args); err != nil {
		log.Fatal(err)
	}

}

func test(minOcena, maxOcena, minStOcen int) {
	studenti := make(map[string]red.Student)

	studenti["63210001"] = red.Student{Ime: "Ana", Priimek: "Novak", Ocene: []int{10, 9}}
	studenti["63210002"] = red.Student{Ime: "Boris", Priimek: "Kralj", Ocene: []int{6, 7}}
	studenti["63210003"] = red.Student{Ime: "Janez", Priimek: "Novak", Ocene: []int{4, 5, 3}}

	red.DodajOceno(studenti, "63210001", 8, minOcena, maxOcena)
	red.DodajOceno(studenti, "63210002", 5, minOcena, maxOcena)
	red.DodajOceno(studenti, "63210003", 5, minOcena, maxOcena)

	red.DodajOceno(studenti, "63210001", 15, minOcena, maxOcena)

	red.DodajOceno(studenti, "63210099", 8, minOcena, maxOcena)

	red.IzpisRedovalnice(studenti)

	studenti["63210004"] = red.Student{Ime: "Mojca", Priimek: "Potok", Ocene: []int{8, 9, 10}}

	red.IzpisiKoncniUspeh(studenti, minStOcen)

	red.DodajOceno(studenti, "63210001", 10, minOcena, maxOcena)
	red.DodajOceno(studenti, "63210001", 10, minOcena, maxOcena)
	red.DodajOceno(studenti, "63210001", 9, minOcena, maxOcena)

	red.DodajOceno(studenti, "63210003", 2, minOcena, maxOcena)
	red.DodajOceno(studenti, "63210003", 3, minOcena, maxOcena)

	red.DodajOceno(studenti, "63210004", 7, minOcena, maxOcena)
	red.DodajOceno(studenti, "63210004", 8, minOcena, maxOcena)
	red.DodajOceno(studenti, "63210004", 9, minOcena, maxOcena)

	red.IzpisRedovalnice(studenti)

	red.IzpisiKoncniUspeh(studenti, minStOcen)
}
