using System.Collections;
using System.Collections.Generic;
using UnityEngine;
using UnityEngine.UI;

[ExecuteInEditMode]
public class CardIcon : MonoBehaviour {

    private void Update()
    {
        var image = GetComponent<Image>();

        var card = GetComponentInParent<Card>();

        if (image.sprite !=  card.icon)
        {
            image.sprite = card.icon;
        }

    }

}
