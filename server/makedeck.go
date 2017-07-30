package main

import (
	"github.com/admiraldolphin/govhack2017/server/game"
	"github.com/admiraldolphin/govhack2017/server/load"
	"log"
)

var (
	prettyDeath = map[string]string{}
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

	// Make cards for traits.
	acs := make([]*game.ActionCard, 0, len(traits))
	for _, t := range traits {
		acs = append(acs, &game.ActionCard{
			Name:  t.Name,
			Trait: t,
		})
	}

	// Scan people to make cards & accumulate matching traits
	var pcs []*game.PersonCard
	for _, p := range ppl {
		pc := &game.PersonCard{
			Name:   p.Name,
			Source: p,
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

		if p.Birth.Year != "" {
			addTrait("le_birth." + p.Birth.Year[:3] + "0")
		}
		if p.Immigration.Year != "" {
			addTrait("le_immigration." + p.Immigration.Year[:3] + "0")
		}
		if p.Convict.Year != "" {
			addTrait("le_convict." + p.Convict.Year[:3] + "0")
		}
		if p.Bankruptcy.Year != "" {
			addTrait("le_bankruptcy." + p.Bankruptcy.Year[:3] + "0")
		}
		if p.Marriage.Year != "" {
			addTrait("le_marriage." + p.Marriage.Year[:3] + "0")
		}
		if p.Court.Year != "" {
			addTrait("le_court." + p.Court.Year[:3] + "0")
		}
		if p.HealthWelfare.Year != "" {
			addTrait("le_health_welfare." + p.HealthWelfare.Year[:3] + "0")
		}
		if p.Census.Year != "" {
			addTrait("le_census." + p.Census.Year[:3] + "0")
		}
	}

	// Normalise PeopleMatching values
	for _, t := range traits {
		t.PeopleMatching /= float64(len(pcs))
	}

	log.Printf("Generated %d people cards and %d action cards", len(pcs), len(acs))
	return &game.Hand{
		People:  pcs,
		Actions: acs,
	}
}
