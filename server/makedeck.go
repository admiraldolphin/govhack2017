package main

import (
	"log"
	"strings"

	"github.com/admiraldolphin/govhack2017/server/game"
	"github.com/admiraldolphin/govhack2017/server/load"
)

var prettyLifeEvents = map[string]string{
	"le_bankruptcy":     "Bankrupt",
	"le_birth":          "Born",
	"le_census":         "Census",
	"le_convict":        "Convicted",
	"le_court":          "Court",
	"le_health_welfare": "Health/Welfare",
	"le_immigration":    "Immigrated",
	"le_marriage":       "Married",
}

// CreateDeck churns a bunch of people into a card deck.
func CreateDeck(ct *load.Cards, ppl []*load.Person) game.Deck {
	// Precompute different possible traits
	traits := make(map[string]*game.Trait)
	for _, d := range ct.Death {
		// Eliminate dc_misc
		if d == "dc_misc" {
			continue
		}
		traits[d] = &game.Trait{
			Key:   d,
			Name:  strings.Title(strings.TrimPrefix(d, "dc_")),
			Death: true,
		}
	}
	for _, le := range ct.LifeEvent {
		for _, dc := range ct.Decade {
			key := le + "." + dc
			traits[key] = &game.Trait{
				Key:   key,
				Name:  prettyLifeEvents[le] + " in " + dc + "s",
				Death: false,
			}
		}
	}

	// Scan people to make cards & accumulate matching traits
	var pcs []*game.PersonCard
	for _, p := range ppl {
		// Eliminate dc_misc
		if len(p.Inquest.DeathCauses) == 1 && p.Inquest.DeathCauses[0] == "dc_misc" {
			continue
		}

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
			// Eliminate dc_misc
			if d == "dc_misc" {
				continue
			}
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

	// Make cards for traits (but only that match someone).
	acs := make([]*game.ActionCard, 0, len(traits))
	for _, t := range traits {
		if t.PeopleMatching < 1 {
			continue
		}

		card := &game.ActionCard{
			Name:  t.Name,
			Trait: t,
		}
		acs = append(acs, card)
		// Add 2 of each death card
		if t.Death {
			acs = append(acs, card)
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
