import React, { Component } from 'react';
import './App.css';

class Data extends Component {
  render() {
    return (
      <div>
        <p>
          You killed {this.props.data.NAME}!
        </p>
        {this.props.data.DEATH_DATE &&
          <p>
            Date of DEATH: {this.props.data.DEATH_DATE}
          </p>
        }
        {this.props.data.AGE &&
        <p>
          Aged {this.props.data.AGE}
        </p>
        }
        {this.props.data.VERDICT &&
        <p>
          {this.props.data.VERDICT}
        </p>
        }
      </div>
    );
  }
}

export default Data;

