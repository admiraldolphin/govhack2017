using System;
using System.Collections;
using System.Collections.Generic;
using UnityEngine;

public class Game : MonoBehaviour {

    public State state { get; private set; }
    public int myPlayerNumber { get; private set; }

    public static Game instance;

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

        var oldState = state;
        this.state = response.state;

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
