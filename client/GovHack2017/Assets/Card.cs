using System;
using System.Collections;
using System.Collections.Generic;
using UnityEngine;
using UnityEngine.EventSystems;

[CreateAssetMenu]
public class CardSpriteSettings : ScriptableObject
{
    public Sprite lifeCard;
    public Sprite deathCard;

    public Sprite lifeCardSelected;
    public Sprite deathCardSelected;
}

[Serializable]
public class TraitEntry
{
    public string key;
    public string name;
    public bool death;
}

[Serializable]
public class PersonCard
{
    public string name;
    public TraitEntry[] traits;
}

[Serializable]
public class PersonCardState
{
    public PersonCard card;
    public bool dead;

    public int[] completed_traits;
    public int score;
}

[Serializable]
public class HandState
{
    public PersonCardState[] people;
    public ActionCardState[] actions;
}

[Serializable]
public class Player
{
    public string name;
    public HandState hand;
    public int score;
}

[Serializable]
public class State
{
    public enum Type
    {
        Lobby = 0,
        InGame = 1,
        GameOver = 2
    }

    public State.Type state;

    public Dictionary<int, Player> players;

    public int clock;
    public int whose_turn;
}

[Serializable]
public class ActionCard
{
    public string name;

    public TraitEntry trait;
}

[Serializable]
public class ActionCardState
{
    public ActionCard card;
    public bool played;
    public bool discarded;
}

[Serializable]
public class Response
{
    public int you;
    public State state;
}

public class Card : MonoBehaviour, IPointerClickHandler {

    public enum Type
    {
        Life,
        Death
    }

    public Card.Type type;

    public string text;

    public Sprite icon;

    public bool selected;

    void IPointerClickHandler.OnPointerClick(PointerEventData eventData)
    {
        selected = !selected;
    }
}
