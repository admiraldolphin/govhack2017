import React, { Component } from 'react';
import './App.css';
import InlineSVG from "react-inlinesvg";

class Navbar extends Component {
  render() {
    let deathtype = 'misc';
    let deathimage = require('./images/dc/misc.svg');

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
								<InlineSVG src={ deathimage } />
            </li>
        )}
        </ul>
      </div>
    );
  }
}

export default Navbar;