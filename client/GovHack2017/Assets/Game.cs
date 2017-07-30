using System;
using System.Collections;
using System.Collections.Generic;
using System.Linq;
using UnityEngine;

public class Game : MonoBehaviour {

    public class TurnOutcome
    {
        public int[] peopleKilled;
        public int[] peopleTraitsUpdated;
    }

    public TurnOutcome ComputeTurnOutcome(State oldState, State newState)
    {
        
        if (new[] { oldState.players, newState.players }.SelectMany(p => p.Values).Any(p => p.hand == null))
        {
            return null;
        }

        var outcome = new TurnOutcome();


        var allOldPeople = oldState.players.SelectMany(p => p.Value.hand.people);
        var allNewPeople = newState.players.SelectMany(p => p.Value.hand.people);

        var allPeople = allOldPeople
            .Join(allNewPeople, p => p.card.ID, p => p.card.ID, (oldPerson, newPerson) => new { oldPerson, newPerson });

        var peopleKilled = allPeople
            .Where(x => x.oldPerson.dead == false && x.newPerson.dead == true)
            .Select(x => x.newPerson.card.ID)
            .ToArray();

        outcome.peopleKilled = peopleKilled;

        var peopleTraitsUpdated = allPeople
            .Where(x => x.newPerson.completed_traits.Length > x.oldPerson.completed_traits.Length)
            .Select(x => x.newPerson.card.ID)
            .ToArray();

        outcome.peopleTraitsUpdated = peopleTraitsUpdated;

        return outcome;

    }

    public State state { get; private set; }
    public int myPlayerNumber { get; private set; }

    public static Game instance;

    public TurnOutcome lastTurnOutcome;

    // people and actions
    public HandState myHand
    {
        get
        {
            return state.players[myPlayerNumber].hand;
        }
    }
    
    // Params = old, new
    public event System.Action<State, State> StateUpdated;

    private ServerConnection connection
    {
        get
        {
            return GetComponent<ServerConnection>();
        }
    }

    public bool isMyTurn {
        get
        {
            return state.whose_turn == myPlayerNumber;
        }
    }

    private void Awake()
    {
        if (instance == null)
        {
            instance = this;
            DontDestroyOnLoad(this);
            ServerConnection.MessageReceived += MessageReceived;
        } else
        {
            Destroy(this);
        }
    }
    
    private IEnumerator Start()
    {
        // wait a frame and then connect
        yield return new WaitForEndOfFrame();
        connection.Connect();
    }

    private void MessageReceived(string message)
    {
        var response = Newtonsoft.Json.JsonConvert.DeserializeObject<Response>(message);

        this.myPlayerNumber = response.you;
        
        var oldState = this.state;
        this.state = response.state;
        
        if (oldState != null)
        {
            lastTurnOutcome = ComputeTurnOutcome(oldState, this.state);
        }
        
        if (StateUpdated != null)
        {
            StateUpdated(oldState, this.state);
        }

    }

    public void StartGame()
    {
        connection.SendStartGame();
    }

    internal void PlayCard(int index)
    {
        connection.SendPlayCard(index);
    }

    internal void DiscardCard(int index)
    {
        connection.SendDiscardCard(index);
    }
}
