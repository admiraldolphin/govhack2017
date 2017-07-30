using System.Collections;
using System.Collections.Generic;
using UnityEngine;

public class Trait : MonoBehaviour {

    public string description;
    public float percent; // 0 - 1

    public UnityEngine.UI.Text descriptionLabel;
    public UnityEngine.UI.Text percentLabel;

    public void Start()
    {
        descriptionLabel.text = description;

        var percentString = string.Format("{0}%", (int)(percent * 100));
        percentLabel.text = percentString;

    }
}
