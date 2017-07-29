import React, { Component } from 'react';
import './App.css';

class Data extends Component {
  render() {
    return (
      <div>
        <p>
          This is some {this.props.data}
        </p>
      </div>
    );
  }
}

export default Data;

