using System.Collections;
using System.Collections.Generic;
using UnityEngine;
using UnityEngine.SceneManagement;

public class LobbyMenu : MonoBehaviour {

	public void ExitLobby()
    {
        SceneManager.LoadScene("MainMenu");
    }

    public void StartGame() 
    {

    }

}
