using System.Collections;
using System.Collections.Generic;
using UnityEngine;
using UnityEngine.UI;

[ExecuteInEditMode]
public class CardBackground : MonoBehaviour
{
    public CardSpriteSettings settings;

    void Update()
    {
        Card card = GetComponentInParent<Card>();

        Sprite sprite;

        switch (card.type)
        {
            case Card.Type.Death:
                sprite = card.selected ? settings.deathCardSelected : settings.deathCard;
                break;

            case Card.Type.Life:
            default:

                sprite = card.selected ? settings.lifeCardSelected : settings.lifeCard;
                break;
        }

        GetComponent<Image>().sprite = sprite;
    }
}