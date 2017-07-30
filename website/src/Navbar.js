import React, { Component } from 'react';
import './App.css';
import InlineSVG from "react-inlinesvg";

class Navbar extends Component {
  render() {

    return (
      <div className="nav">
        <ul>
        { this.props.data.map (
          (item, key) =>
            <li
              onClick={() => this.props.changeTab(key)}
              key={key}
              className={this.props.isCurrent(key) ? 'current' : ''}
              style = {{
	              transform: 'rotate(' + (Math.random() * 8 - 4) + 'deg)'
              }}
            >
                <span>{item.card.source.name}</span>
	              {    console.log(item.card.traits[0].key) }
		            <InlineSVG src={ require('./images/dc/' + item.card.traits[0].key.replace('dc_', '') + '.svg') } />
            </li>
        )}
        </ul>
      </div>
    );
  }
}

export default Navbar;