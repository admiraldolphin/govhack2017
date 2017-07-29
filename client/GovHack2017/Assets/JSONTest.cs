using System.Collections;
using System.Collections.Generic;
using UnityEngine;
using Newtonsoft.Json;

public class JSONTest : MonoBehaviour {

	// Use this for initialization
	void Start () {

        var text = @"{""you"":0,""state"":{""state"":0,""players"":{""0"":{""name"":""Player 0"",""hand"":null,""score"":0},""2"":{""name"":""Player 2"",""hand"":null,""score"":0}},""clock"":0,""whose_turn"":0}}";
        var state = JsonConvert.DeserializeObject<Response>(text);  //JsonUtility.FromJson<Response>(text);

	}
	
	
}
