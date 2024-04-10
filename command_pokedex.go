package main

import "fmt"

func commandPokedex (cfg	*config, args ...string) error {
	fmt.Println("Pokedex:")
	for name := range cfg.caughtPokemon {
		fmt.Println(name)
	}
	return nil
}