import React, { Component } from 'react';
import './App.css';

class Navbar extends Component {
  render() {
    return (
      <div>
        { this.props.data.map (
          (item, key) =>
            <button
              onClick={() => this.props.changeTab(key)}
              key={key}>{item.NAME}
            </button>
        )}
      </div>
    );
  }
}

export default Navbar;