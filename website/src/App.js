import React, { Component } from 'react';
import Data from './Data';
import './App.css';

class App extends Component {

  constructor(props) {
    super(props);
    this.state = {
      data: "some data"
    };
  }

  componentDidMount() {
    this.getData()
  }

  getData() {
    // testing to see if we can hit dataset, this will change once we hit json from server
    fetch('http://data.gov.au/dataset/5b5ced0f-0032-4178-a874-57b92d7b09d9/resource/980aa7af-45ce-49a8-a813-013d20c52011/download/inquests.json')
      .then((response) => response.json())
      .then((response) => Object.keys(response).map(function (key) { return response[key]; }) )
      .then(response => {
          this.setState({data: response[0].NAME})
        }
      )
      .catch(error => console.log(error))
  }

  render() {

    return (
      <div className="App">
        <Data
          data={this.state.data}
        />
      </div>
    );
  }
}

export default App;

