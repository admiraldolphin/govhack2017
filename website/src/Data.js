import React, { Component } from 'react';
import './App.css';

class Data extends Component {
  render() {
    return (
      <div className="data">
        <h1>
          You killed {this.props.data.name}!
        </h1>
        {this.props.data.birth.birth_date &&
        <p>
          They were born on {this.props.data.birth.birth_date} at {this.props.data.birth.birth_place} to {this.props.data.birth.birth_mother} and {this.props.data.birth.birth_father} [<a href={this.props.data.birth.permalink}>Ref</a>]
        </p>
        }
        {this.props.data.immigration.from_country &&
        <p>
          They immigrated from {this.props.data.immigration.from_country} in {this.props.data.immigration.year} [<a href={this.props.data.immigration.permalink}>Ref</a>]
        </p>
        }
        {this.props.data.convict.convict_port &&
        <p>
          They were a convict from {this.props.data.convict.convict_port} and arrived on the ship {this.props.data.convict.convict_ship} in {this.props.data.convict.year} [<a href={this.props.data.convict.permalink}>Ref</a>]
        </p>
        }
        {this.props.data.bankruptcy.bankrupt_date &&
        <p>
          They were declared bankrupt on {this.props.data.bankruptcy.bankrupt_date} [<a href={this.props.data.bankruptcy.permalink}>Ref</a>]
        </p>
        }
        {this.props.data.marriage.marriage_date &&
        <p>
          They married on {this.props.data.marriage.marriage_date} in {this.props.data.marriage.marriage_place} to {this.props.data.marriage.spouse_name} [<a href={this.props.data.marriage.permalink}>Ref</a>]
        </p>
        }
        {this.props.data.court.trial_offence &&
        <p>
          They went to court for {this.props.data.court.trial_offence} in {this.props.data.court.year} [<a href={this.props.data.court.permalink}>Ref</a>]
        </p>
        }
        {this.props.data.census.year &&
        <p>
          They were on the {this.props.data.census.year} census [<a href={this.props.data.census.permalink}>Ref</a>]
        </p>
        }
        {this.props.data.inquest.death_verdict &&
        <p>
          They died from {this.props.data.inquest.death_verdict} on {this.props.data.inquest.death_date} [<a href={this.props.data.inquest.permalink}>Ref</a>]
        </p>
        }

      </div>
    );
  }
}

export default Data;

