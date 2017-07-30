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

        var traits = state.card.traits;

        for (int i = 0; i < traits.Length; i++)
        {
            var trait = traits[i];

            var traitUI = Instantiate(traitPrefab);
            traitUI.traitDescription = trait.name;
            traitUI.percent = trait.people_matching;

            traitUI.transform.SetParent(traitContainer, false);

            traitUI.status = StatusForTrait(state, i);
            
        }
        

        this.gameObject.SetActive(true);
    }
}
