import React, { Component } from 'react';
import './App.css';

class Navbar extends Component {
  render() {
    return (
      <div className="nav">
        <ul>
        { this.props.data.map (
          (item, key) =>
            <li
              onClick={() => this.props.changeTab(key)}
              key={key}>
                <span>{item.card.source.name}</span>
            </li>
        )}
        </ul>
      </div>
    );
  }
}

export default Navbar;