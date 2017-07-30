package main

import (
	"github.com/admiraldolphin/govhack2017/server/game"
	"github.com/admiraldolphin/govhack2017/server/load"
	"log"
)

// CreateDeck churns a bunch of people into a card deck.
func CreateDeck(ct *load.Cards, ppl []*load.Person) game.Deck {
	// Precompute different possible traits
	traits := make(map[string]*game.Trait)
	for _, d := range ct.Death {
		traits[d] = &game.Trait{
			Key:   d,
			Name:  d, // TODO: better name
			Death: true,
		}
	}
	for _, le := range ct.LifeEvent {
		for _, dc := range ct.Decade {
			key := le + "." + dc
			traits[key] = &game.Trait{
				Key:   key,
				Name:  key, // TODO: better name
				Death: false,
			}
		}
	}

	// Scan people to make cards & find matching traits
	var pcs []*game.PersonCard
	for _, p := range ppl {
		pc := &game.PersonCard{
			Name: p.Name,
		}
		pcs = append(pcs, pc)

		addTrait := func(key string) {
			t := traits[key]
			if t == nil {
				log.Printf("Couldn't find trait %q", key)
				return
			}
			t.PeopleMatching += 1.0
			pc.Traits = append(pc.Traits, t)
		}

		for _, d := range p.Inquest.DeathCauses {
			addTrait(d)
		}

		// "le_birth", Birth
		if p.Birth.Year != "" {
			addTrait("le_birth." + p.Birth.Year[:3] + "0")
		}

		// "le_immigration", Immigration
		// "le_convict", Convict
		// "le_bankruptcy", Bankruptcy
		// "le_marriage", Marriage
		// "le_court", Court
		// "le_health_welfare", HealthWelfare
		// "le_census", Census
	}

	// Normalise PeopleMatching
	for _, t := range traits {
		t.PeopleMatching /= float64(len(pcs))
	}
	return nil
}
