import React, { Component } from 'react';
import './App.css';
import Data from './Data';
import Navbar from './Navbar';

class App extends Component {

  constructor(props) {
    super(props);
    this.state = {
      data: [
        {"AGE":"78","CATKEY":"(Sirsi) 1352730","DEATH_DATE":"28 May 1913","FORMAT":"VIEW","FORMAT_LINCTAS":"Names Index","INDEX":"Inquests","INQUEST_DATE":"29 May 1913","LINC_TAS_AVAIL":"Online","NAME":"Ward, David George","NAME_FULL_DISPLAY":"Ward, David George","PERMA_LINK":"https:\/\/linctas.ent.sirsidynix.net.au\/client\/en_AU\/all\/search\/detailnonmodal\/ent:$002f$002fNAME_INDEXES$002f0$002fNAME_INDEXES:1352730\/one","PUBDATE":"1913","PUBDATE_RANGE":"1913","REFERENCE_URL":{"URL":"https:\/\/stors.tas.gov.au\/AGD20-1-14-20765","URL_TEXT":"AGD20\/1\/14\/20765"},"TASMANIAN":["Published in Tasmania","About Tasmania","By a Tasmanian"],"VERDICT":"Natural causes heart failure","YEAR":"1913"},
        {"AGE":"9","CATKEY":"(Sirsi) 1352726","DEATH_DATE":"24 Dec 1893","FORMAT":"VIEW","FORMAT_LINCTAS":"Names Index","INDEX":"Inquests","INQUEST_DATE":"27 Dec 1893","LINC_TAS_AVAIL":"Online","NAME":"Partridge Or Lovett, George Thomas","NAME_FULL_DISPLAY":"Partridge Or Lovett, George Thomas","PERMA_LINK":"https:\/\/linctas.ent.sirsidynix.net.au\/client\/en_AU\/all\/search\/detailnonmodal\/ent:$002f$002fNAME_INDEXES$002f0$002fNAME_INDEXES:1352726\/one","PUBDATE":"1893","PUBDATE_RANGE":"1893","REFERENCE_URL":[{"URL":"https:\/\/stors.tas.gov.au\/AGD20-1-7-8349","URL_TEXT":"AGD20\/1\/7\/8349"},{"URL":"https:\/\/stors.tas.gov.au\/POL709-1-25","URL_TEXT":"POL709\/1\/25 p.3 (1894)"},{"URL":"https:\/\/stors.tas.gov.au\/SC195-1-70-10333","URL_TEXT":"SC195\/1\/70 Inquest 10333"}],"SHIP_NATIVE_PLACE":"Born in Tasmania","TASMANIAN":["Published in Tasmania","About Tasmania","By a Tasmanian"],"VERDICT":"Accidentally drowned","YEAR":"1893"}
        ],
      focusId: 0
    };

  }

  componentDidMount() {
    // this.getData()
    console.log(this.props.match.params.id)
  }

  changeTab(id) {
    console.log(id)
    this.setState({focusId: id})
  }

  getData() {
    // testing to see if we can hit dataset, this will change once we hit json from server
    fetch('http://data.gov.au/dataset/5b5ced0f-0032-4178-a874-57b92d7b09d9/resource/980aa7af-45ce-49a8-a813-013d20c52011/download/inquests.json')
      .then((response) => response.json())
      .then((response) => Object.keys(response).map(function (key) { return response[key]; }) )
      .then(response => {
          this.setState({data: response[0]})
        }
      )
      .catch(error => console.log(error))
  }

  render() {

    return (
      <div className="App">
        <img src={require('./images/Logo_Transparent_Subtitle.png')} className="HeaderImage" />
        <Navbar
          data={this.state.data}
          changeTab={this.changeTab.bind(this)}
        />
        <Data
          data={this.state.data[this.state.focusId]}
        />
      </div>
    );
  }
}

export default App;

