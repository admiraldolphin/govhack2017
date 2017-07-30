using System.Collections;
using System.Collections.Generic;
using UnityEngine;
using UnityEngine.UI;
using UnityEngine.EventSystems;
using System;

public class PersonIcon : MonoBehaviour, IPointerClickHandler {

    public Sprite aliveSprite;
    public Sprite deadSprite;
    public Sprite recentlyDeadSprite;

    private PersonCardState _state;

    public PersonCardState state
    {
        get
        {
            return _state;
        }
        set
        {
            _state = value;

            var sprite = value.dead ? deadSprite : aliveSprite;

            icon.sprite = sprite;
        }
    }

    public Image icon
    {
        get
        {
            return GetComponentInChildren<Image>();
        }
    }

    void IPointerClickHandler.OnPointerClick(PointerEventData eventData)
    {
        FindObjectOfType<GameUI>().ShowPersonInfo(this.state);
    }

    public void JustBecameDead()
    {
        StartCoroutine(RunJustBecameDead());
    }


    IEnumerator RunJustBecameDead()
    {
        this.icon.sprite = recentlyDeadSprite;

        yield return new WaitForSeconds(1.0f);

        this.icon.sprite = deadSprite;
    }
}
