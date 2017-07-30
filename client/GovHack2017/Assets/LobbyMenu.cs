using System.Collections;
using System.Collections.Generic;
using UnityEngine;
using UnityEngine.SceneManagement;
using UnityEngine.UI;

public class LobbyMenu : MonoBehaviour {

    public Text lobbyStatusText;
    public Text readyLabel;

    public RectTransform lobbyContainer;

    private void Start()
    {
        Game.instance.StateUpdated += StateUpdated;
    }

    private void StateUpdated(State oldState, State newState)
    {
        var playerCount = newState.players.Count;

        var playerStatus = string.Format("{0} player{1}", playerCount, playerCount == 1 ? "" : "s");

        lobbyStatusText.text = playerStatus;

        readyLabel.gameObject.SetActive(playerCount > 1);
        
        if (newState.state == State.Type.InGame)
        {
            SceneManager.LoadScene("Game");
        }

        if (newState.state == State.Type.GameOver)
        {
            SceneManager.LoadScene("MainMenu");
        }
    }

    public void ExitLobby()
    {
        SceneManager.LoadScene("MainMenu");
    }

    private void OnDestroy()
    {
        Game.instance.StateUpdated -= StateUpdated;
    }



}
