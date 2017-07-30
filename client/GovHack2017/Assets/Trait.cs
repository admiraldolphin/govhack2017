using System.Collections;
using System.Collections.Generic;
using UnityEngine;
using UnityEngine.UI;

public class Trait : MonoBehaviour {
    
    public UnityEngine.UI.Text descriptionLabel;
    public UnityEngine.UI.Text percentLabel;

    public Image statusIcon;

    public Sprite completedSprite;
    public Sprite notCompletableSprite;

    public enum Status
    {
        NotYetCompleted,
        Completed,
        NotCompletable
    }

    [HideInInspector]
    public Sprite traitIcon;

    public Status status
    {
        set
        {
            statusIcon.enabled = true;
            switch (value)
            {
                case Status.Completed:
                    statusIcon.sprite = completedSprite;
                    return;
                case Status.NotCompletable:
                    statusIcon.sprite = notCompletableSprite;
                    return;
                case Status.NotYetCompleted:
                    statusIcon.sprite = traitIcon;
                    return;
            }
        }
    }

    public string traitDescription {
        set
        {
            descriptionLabel.text = value;
        }
    }

    public float percent
    {
        set
        {
            var percentString = string.Format("{0}%", (int)(value * 100));
            percentLabel.text = percentString;
        }
    }
    
}
