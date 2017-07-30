using System.Collections;
using System.Collections.Generic;
using UnityEngine;
using UnityEngine.SceneManagement;

public class LobbyMenu : MonoBehaviour {

    public LobbyPlayerIcon playerPrefab;

    public RectTransform lobbyContainer;

    private void Start()
    {
        Game.instance.StateUpdated += StateUpdated;
    }

    private void StateUpdated(State oldState, State newState)
    {
        foreach (Transform player in lobbyContainer)
        {
            Destroy(player.gameObject);
        }

        foreach (var player in newState.players)
        {
            var newItem = Instantiate(playerPrefab);

            newItem.playerName = player.Value.name;

            newItem.isCurrentPlayer = player.Key == Game.instance.myPlayerNumber;
            
            newItem.transform.SetParent(lobbyContainer, false);
        }
        
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
