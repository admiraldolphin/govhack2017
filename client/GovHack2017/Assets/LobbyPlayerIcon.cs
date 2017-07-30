using System.Collections;
using System.Collections.Generic;
using UnityEngine;
using UnityEngine.UI;

public class LobbyPlayerIcon : MonoBehaviour {

    private Image image
    {
        get
        {
            return GetComponent<Image>();
        }
    }

    private Text text
    {
        get
        {
            return GetComponentInChildren<Text>();
        }
    }

    public Sprite currentPlayer;
    public Sprite otherPlayer;

    public bool isCurrentPlayer
    {
        set
        {
            if (value)
            {
                image.sprite = currentPlayer;
            } else
            {
                image.sprite = otherPlayer;
            }
        }
    }

    public string playerName
    {
        set
        {
            text.text = value;
        }
    }
    
}
