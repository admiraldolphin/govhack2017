using System.Collections;
using System.Collections.Generic;
using UnityEngine;
using UnityEngine.UI;

[ExecuteInEditMode]
public class CardText : MonoBehaviour {

    private void Update()
    {
        var text = GetComponent<Text>();

        var card = GetComponentInParent<Card>();

        if (text.text != card.text)
        {
            text.text = card.text;
        }

    }
}
