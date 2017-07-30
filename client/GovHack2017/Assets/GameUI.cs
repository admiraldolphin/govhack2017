using System;
using System.Collections;
using System.Collections.Generic;
using UnityEngine;
using System.Linq;
using UnityEngine.UI;

public class GameUI : MonoBehaviour {

    public RectTransform cardContainer;
    public RectTransform personContainer;

    public Card cardPrefab;
    public PersonIcon personPrefab;

    public PersonInfoSheet personInfoSheet;

    public Text clock;
    public Text turnStatus; // "your turn" etc
    public GameObject gameOverUI;

    public TurnOutcomeUI turnOutcomeUI;

    public Text gameOverSummary;

    // Use this for initialization
    void Start () {
        Game.instance.StateUpdated += StateUpdated;

        StateUpdated(Game.instance.state, Game.instance.state);
	}

    private void StateUpdated(State oldState, State newState)
    {
        if (newState.state == State.Type.GameOver)
        {
            gameOverUI.SetActive(true);

            gameOverSummary.text = GenerateEndReport();

            return;
        }

        clock.text = newState.clock.ToString();

        if (Game.instance.isMyTurn)
        {
            turnStatus.text = "Your Turn";
        } else
        {
            turnStatus.text = "Waiting...";
        }

        // Remove all cards
        foreach (Transform card in cardContainer.transform)
        {
            Destroy(card.gameObject);
        }

        foreach (Transform person in personContainer.transform)
        {
            Destroy(person.gameObject);            
        }

        // Add current hand
        foreach (var card in Game.instance.myHand.actions)
        {
            var cardUI = Instantiate(cardPrefab);

            cardUI.transform.SetParent(cardContainer, false);

            cardUI.state = card;
        }

        foreach (var person in Game.instance.myHand.people)
        {
            var personUI = Instantiate(personPrefab);

            personUI.transform.SetParent(personContainer, false);

            personUI.state = person;

            // Did they recently die?
            if (Game.instance.lastTurnOutcome != null && Game.instance.lastTurnOutcome.peopleKilled.Contains(person.card.ID))
            {
                personUI.JustBecameDead();
            }
        }

        var myPeopleRecentlyDead = Game.instance.myHand.people
            .Where(p => Game.instance.lastTurnOutcome.peopleKilled.Contains(p.card.ID));

        if (Game.instance.lastTurnOutcome != null)
        {
            turnOutcomeUI.ShowTurnOutcome(Game.instance.lastTurnOutcome);
        }
        
       
    }

    private string GenerateEndReport()
    {

        var playerSummary = Game.instance.state.players.OrderByDescending(p => p.Value.score)
            .Select(p =>
            {
                if (p.Key == Game.instance.myPlayerNumber)
                {
                    return string.Format("{0} (you): {1} point{2}", p.Value.name, p.Value.score, p.Value.score == 1 ? "" : "s");
                }
                else
                {
                    return string.Format("{0}: {1} point{2}", p.Value.name, p.Value.score, p.Value.score == 1 ? "" : "s");
                }
            });
         return string.Join("\n", playerSummary.ToArray());
            

    }

    public void PlayCard()
    {
        if (Game.instance.isMyTurn == false)
            return;

        if (selectedCard == null)
            return;

        var index = Game.instance.myHand.IndexOfCard(selectedCard.state);

        Game.instance.PlayCard(index);
    }

    public void DiscardCard()
    {
        if (Game.instance.isMyTurn == false)
            return;

        if (selectedCard == null)
            return;

        var index = Game.instance.myHand.IndexOfCard(selectedCard.state);

        Game.instance.DiscardCard(index);
    }

    public void ShowPersonInfo(PersonCardState state)
    {
        personInfoSheet.ShowWithPerson(state);
    }

    public Card selectedCard
    {
        get
        {
            return cardContainer.GetComponentsInChildren<Card>().Where(c => c.selected).FirstOrDefault();
        }
    }

    internal void CardSelected(Card card)
    {
        if (Game.instance.isMyTurn == false)
        {
            return;
        }

        var cards = cardContainer.GetComponentsInChildren<Card>();

        foreach (var theCard in cards)
        {
            theCard.selected = theCard == card;
        }

        
    }

    public string gameResultsURL = "http://deathwho.herokuapp.com/{0}";

    public void ShowEndGameResults()
    {
        var url = string.Format(gameResultsURL, Game.instance.myPlayerNumber);

        Application.OpenURL(url);
    }
}
