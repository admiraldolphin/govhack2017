import React, { Component } from 'react';
import './App.css';

class Data extends Component {

  render() {
    return (
      <div className="main">

      <div className="data">
        <h1>
          You killed {this.props.data.name}!
        </h1>

        {this.props.data.birth.birth_date &&
        <p>
          <a title="View reference" href={this.props.data.birth.permalink}>
            They were born on {this.props.data.birth.birth_date} at {this.props.data.birth.birth_place} to {this.props.data.birth.birth_mother} and {this.props.data.birth.birth_father}.
          </a>
        </p>
        }
        {this.props.data.immigration.from_country &&
        <p>
          <a title="View reference" href={this.props.data.immigration.permalink}>
            They immigrated from {this.props.data.immigration.from_country} in {this.props.data.immigration.year}.
          </a>
        </p>
        }
        {this.props.data.convict.convict_port &&
        <p>
          <a title="View reference" href={this.props.data.convict.permalink}>
            They were a convict from {this.props.data.convict.convict_port} and arrived on the ship {this.props.data.convict.convict_ship} in {this.props.data.convict.year}.
          </a>
        </p>
        }
        {this.props.data.bankruptcy.bankrupt_date &&
        <p>
          <a title="View reference" href={this.props.data.bankruptcy.permalink}>
            They were declared bankrupt on {this.props.data.bankruptcy.bankrupt_date}.
          </a>
        </p>
        }
        {this.props.data.marriage.marriage_date &&
        <p>
          <a title="View reference" href={this.props.data.marriage.permalink}>
            They married on {this.props.data.marriage.marriage_date} to {this.props.data.marriage.spouse_name}.
          </a>
        </p>
        }
        {this.props.data.court.trial_offence &&
        <p>
          <a title="View reference" href={this.props.data.court.permalink}>
           They went to court for {this.props.data.court.trial_offence} in {this.props.data.court.year}.
          </a>
        </p>
        }
        {this.props.data.census.year &&
        <p>
          <a title="View reference" href={this.props.data.census.permalink}>
            They were on the {this.props.data.census.year} census.
          </a>
        </p>
        }
        {this.props.data.inquest.death_verdict &&
        <p>
          <a title="View reference" href={this.props.data.inquest.permalink}>
            Died from {this.props.data.inquest.death_verdict.toLowerCase()} on {this.props.data.inquest.death_date}.
          </a>
        </p>
        }

      </div>
      </div>
    );
  }
}

export default Data;

