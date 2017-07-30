using System;
using System.Collections;
using System.Collections.Generic;
using System.Linq;
using UnityEngine;
using UnityEngine.UI;

public class PersonInfoSheet : MonoBehaviour {

    public RectTransform traitContainer;

    public Trait traitPrefab;

    public Text nameText;

    Trait.Status StatusForTrait(PersonCardState state, int traitIndex)
    {
        var trait = state.card.traits[traitIndex];
        
        if (state.completed_traits != null && state.completed_traits.Contains(traitIndex))
        {
            return Trait.Status.Completed;
        }

        if (state.dead)
        {
            return Trait.Status.NotCompletable;
        }

        return Trait.Status.NotYetCompleted;

    }

    internal void ShowWithPerson(PersonCardState state)
    {
        nameText.text = state.card.name;

        foreach (Transform trait in traitContainer)
        {
            Destroy(trait.gameObject);
        }

        var deathTrait = state.card.traits.Where(t => t.death).OrderBy(t => t.name == "Misc").Take(1);
        var otherTraits = state.card.traits.Where(t => t.death == false).Take(2);

        var traits = state.card.traits;

        var allTraits = otherTraits.Concat(deathTrait);

        foreach (var trait in allTraits)
        { 
            var index = Array.IndexOf(traits, trait);

            var traitUI = Instantiate(traitPrefab);

            string name;
            if (trait.death)
            {
                name = "Cause of Death: " + trait.name;    
            } else
            {
                name = trait.name;
            }
            traitUI.traitDescription = name;
            traitUI.percent = trait.people_matching;

            traitUI.transform.SetParent(traitContainer, false);

            traitUI.status = StatusForTrait(state, index);
            
        }
        

        this.gameObject.SetActive(true);
    }
}
