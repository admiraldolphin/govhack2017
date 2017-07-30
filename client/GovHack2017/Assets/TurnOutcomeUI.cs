using System.Collections;
using System.Collections.Generic;
using UnityEngine;
using UnityEngine.UI;

public class TurnOutcomeUI : MonoBehaviour {

    public Card cardUI;

    public Text outcomeLabel;

    private CanvasGroup group
    {
        get
        {
            return GetComponent<CanvasGroup>();
        }
    }

    IEnumerator Fade(float time)
    {
        var timeRemaining = time;

        while (timeRemaining > 0)
        {
            timeRemaining -= Time.deltaTime;

            var t = 1.0f - timeRemaining / time;

            var alpha = Mathf.Lerp(1, 0, t);

            group.alpha = alpha;

            yield return null;
        }

        this.gameObject.SetActive(false);
    }

    public void ShowTurnOutcome(Game.TurnOutcome outcome)
    {
        if (outcome.peopleKilled.Length > 0)
        {
            var noun = outcome.peopleKilled.Length == 1 ? "person" : "people";
            outcomeLabel.text = string.Format("{0} {1} killed", outcome.peopleKilled.Length, noun);
        }
        else if(outcome.peopleTraitsUpdated.Length > 0)
        {
            var noun = outcome.peopleTraitsUpdated.Length == 1 ? "person's" : "peoples'";

            outcomeLabel.text = string.Format("{0} life events matched!", outcome.peopleTraitsUpdated.Length);
        } else
        {
            this.gameObject.SetActive(false);

            return;
        }

        this.gameObject.SetActive(true);

        group.alpha = 1;

        StartCoroutine(Fade(fadeTime));

    }

    public float fadeTime = 2.0f;
}
