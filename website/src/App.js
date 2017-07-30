import React, { Component } from 'react';
import './App.css';
import Data from './Data';
import Navbar from './Navbar';

class App extends Component {

  constructor(props) {
    super(props);
    this.state = {
      data:
        [
      {
        "card": {
          "name": "Towers, John",
          "traits": [
            {
              "key": "dc_misc",
              "name": "Misc",
              "death": true,
              "people_matching": 1
            },
            {
              "key": "le_convict.1840",
              "name": "Convicted in 1840s",
              "death": false,
              "people_matching": 0.2269315673289183
            }
          ],
          "source": {
            "name": "Towers, John",
            "inquest": {
              "death_date": "31 Dec 1880",
              "death_verdict": "Natural causes",
              "death_causes": [
                "dc_misc"
              ],
              "permalink": "https://linctas.ent.sirsidynix.net.au/client/en_AU/all/search/detailnonmodal/ent:$002f$002fNAME_INDEXES$002f0$002fNAME_INDEXES:1359086/one"
            },
            "birth": {},
            "immigration": {},
            "convict": {
              "departure_date": "1 May 1846",
              "convict_port": "Portsmouth",
              "convict_ship": "Palmyra",
              "year": "1846",
              "permalink": "https://linctas.ent.sirsidynix.net.au/client/en_AU/all/search/detailnonmodal/ent:$002f$002fNAME_INDEXES$002f0$002fNAME_INDEXES:1441273/one"
            },
            "bankruptcy": {},
            "marriage": {},
            "court": {},
            "health-welfare": {},
            "census": {}
          }
        },
        "dead": false,
        "completed_traits": null,
        "score": 0
      },
      {
        "card": {
          "name": "Donnelly, John",
          "traits": [
            {
              "key": "dc_misc",
              "name": "Misc",
              "death": true,
              "people_matching": 1
            },
            {
              "key": "le_convict.1840",
              "name": "Convicted in 1840s",
              "death": false,
              "people_matching": 0.2269315673289183
            },
            {
              "key": "le_marriage.1880",
              "name": "Married in 1880s",
              "death": false,
              "people_matching": 0.09492273730684327
            }
          ],
          "source": {
            "name": "Donnelly, John",
            "inquest": {
              "death_date": "1 Jul 1907",
              "death_verdict": "Accidental death",
              "death_causes": [
                "dc_misc"
              ],
              "year": "1907",
              "permalink": "https://linctas.ent.sirsidynix.net.au/client/en_AU/all/search/detailnonmodal/ent:$002f$002fNAME_INDEXES$002f0$002fNAME_INDEXES:1355327/one"
            },
            "birth": {},
            "immigration": {},
            "convict": {
              "departure_date": "2 May 1842",
              "convict_port": "Dublin",
              "convict_ship": "Isabella Watson",
              "year": "1842",
              "permalink": "https://linctas.ent.sirsidynix.net.au/client/en_AU/all/search/detailnonmodal/ent:$002f$002fNAME_INDEXES$002f0$002fNAME_INDEXES:1387869/one"
            },
            "bankruptcy": {},
            "marriage": {
              "marriage_date": "31 May 1881",
              "spouse_name": "Kennedy, Mary",
              "year": "1881",
              "permalink": "https://linctas.ent.sirsidynix.net.au/client/en_AU/all/search/detailnonmodal/ent:$002f$002fNAME_INDEXES$002f0$002fNAME_INDEXES:890357/one"
            },
            "court": {},
            "health-welfare": {},
            "census": {}
          }
        },
        "dead": false,
        "completed_traits": null,
        "score": 0
      },
      {
        "card": {
          "name": "Riseley, Edward",
          "traits": [
            {
              "key": "dc_misc",
              "name": "Misc",
              "death": true,
              "people_matching": 1
            },
            {
              "key": "le_birth.1890",
              "name": "Born in 1890s",
              "death": false,
              "people_matching": 0.06181015452538632
            }
          ],
          "source": {
            "name": "Riseley, Edward",
            "inquest": {
              "death_date": "16 Sep 1908",
              "death_verdict": "Accidentally killed through a rope breaking whilst he was being hauled from a shaft at Ridgeway",
              "death_causes": [
                "dc_misc"
              ],
              "year": "1908",
              "permalink": "https://linctas.ent.sirsidynix.net.au/client/en_AU/all/search/detailnonmodal/ent:$002f$002fNAME_INDEXES$002f0$002fNAME_INDEXES:1355521/one"
            },
            "birth": {
              "birth_date": "21 Sep 1894",
              "birth_place": "Geeveston",
              "birth_mother": "Murrell, Mary Jane",
              "birth_father": "Riseley, John William",
              "year": "1894",
              "permalink": "https://linctas.ent.sirsidynix.net.au/client/en_AU/all/search/detailnonmodal/ent:$002f$002fNAME_INDEXES$002f0$002fNAME_INDEXES:1054801/one"
            },
            "immigration": {},
            "convict": {
              "year": ""
            },
            "bankruptcy": {},
            "marriage": {},
            "court": {},
            "health-welfare": {},
            "census": {}
          }
        },
        "dead": false,
        "completed_traits": null,
        "score": 0
      },
      {
        "card": {
          "name": "Perrin, Amy Maud",
          "traits": [
            {
              "key": "dc_misc",
              "name": "Misc",
              "death": true,
              "people_matching": 1
            },
            {
              "key": "le_birth.1870",
              "name": "Born in 1870s",
              "death": false,
              "people_matching": 0.06445916114790287
            }
          ],
          "source": {
            "name": "Perrin, Amy Maud",
            "inquest": {
              "death_date": "6 Jan 1877",
              "death_verdict": "Accidentally drowned in a tank",
              "death_causes": [
                "dc_misc"
              ],
              "year": "1877",
              "permalink": "https://linctas.ent.sirsidynix.net.au/client/en_AU/all/search/detailnonmodal/ent:$002f$002fNAME_INDEXES$002f0$002fNAME_INDEXES:1359565/one"
            },
            "birth": {
              "birth_date": "23 Feb 1873",
              "birth_place": "Launceston",
              "birth_mother": "Wilson, Henrietta Kate",
              "birth_father": "Perrin, Walter",
              "year": "1873",
              "permalink": "https://linctas.ent.sirsidynix.net.au/client/en_AU/all/search/detailnonmodal/ent:$002f$002fNAME_INDEXES$002f0$002fNAME_INDEXES:933883/one"
            },
            "immigration": {},
            "convict": {
              "year": ""
            },
            "bankruptcy": {},
            "marriage": {},
            "court": {},
            "health-welfare": {},
            "census": {}
          }
        },
        "dead": false,
        "completed_traits": null,
        "score": 0
      },
      {
        "card": {
          "name": "Lucas, Joseph",
          "traits": [
            {
              "key": "dc_misc",
              "name": "Misc",
              "death": true,
              "people_matching": 1
            },
            {
              "key": "le_convict.1830",
              "name": "Convicted in 1830s",
              "death": false,
              "people_matching": 0.1403973509933775
            },
            {
              "key": "le_marriage.1840",
              "name": "Married in 1840s",
              "death": false,
              "people_matching": 0.08079470198675497
            }
          ],
          "source": {
            "name": "Lucas, Joseph",
            "inquest": {
              "death_date": "12 Feb 1878",
              "death_verdict": "Chlorodyne overdose",
              "death_causes": [
                "dc_misc"
              ],
              "year": "1878",
              "permalink": "https://linctas.ent.sirsidynix.net.au/client/en_AU/all/search/detailnonmodal/ent:$002f$002fNAME_INDEXES$002f0$002fNAME_INDEXES:1359708/one"
            },
            "birth": {},
            "immigration": {},
            "convict": {
              "departure_date": "2 Jun 1831",
              "convict_port": "London",
              "convict_ship": "William Glen Anderson",
              "year": "1831",
              "permalink": "https://linctas.ent.sirsidynix.net.au/client/en_AU/all/search/detailnonmodal/ent:$002f$002fNAME_INDEXES$002f0$002fNAME_INDEXES:1412910/one"
            },
            "bankruptcy": {},
            "marriage": {
              "marriage_date": "26 Aug 1841",
              "spouse_name": "King, Mary",
              "year": "1841",
              "permalink": "https://linctas.ent.sirsidynix.net.au/client/en_AU/all/search/detailnonmodal/ent:$002f$002fNAME_INDEXES$002f0$002fNAME_INDEXES:829446/one"
            },
            "court": {},
            "health-welfare": {},
            "census": {}
          }
        },
        "dead": false,
        "completed_traits": null,
        "score": 0
      }
    ],
      focusId: 0
    };

  }

  componentDidMount() {
    this.getData()
    console.log(this.props.match.params.id)
  }

  changeTab(id) {
    this.setState({focusId: id})
  }

  getData() {
    fetch('http://35.197.178.221/statusz')
      .then((response) => response.json())
      .then((response) => response.players[this.props.match.params.id].hand.people)
      .then(response => {
          this.setState({data: response})
        }
      )
      .catch(error => console.log(error))
  }

  render() {

    return (
      <div className="App">
        <Navbar
          data={this.state.data}
          changeTab={this.changeTab.bind(this)}
        />
        <Data
          data={this.state.data[this.state.focusId].card.source}
        />
      </div>
    );
  }
}

export default App;

