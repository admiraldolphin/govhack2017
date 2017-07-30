using System.Collections;
using System.Collections.Generic;
using UnityEngine;
using UnityEngine.UI;

[ExecuteInEditMode]
public class CardBackground : MonoBehaviour
{
    public Sprite lifeCard;
    public Sprite deathCard;

    public Sprite lifeCardSelected;
    public Sprite deathCardSelected;
   

    void Update()
    {
        
        Card card = GetComponentInParent<Card>();

        Sprite sprite;

        switch (card.type)
        {
            case Card.Type.Death:
                sprite = card.selected ? deathCardSelected : deathCard;
                break;

            case Card.Type.Life:
            default:

                sprite = card.selected ? lifeCardSelected : lifeCard;
                break;
        }

        GetComponent<Image>().sprite = sprite;
    }
}