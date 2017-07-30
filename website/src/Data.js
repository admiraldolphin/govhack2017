import React, { Component } from 'react';
import './App.css';

class Data extends Component {
  render() {
    return (
      <div className="main">

      <div className="data">
        <h1>
          You killed {this.props.data.name.split(', ').reverse().join(' ')}!
        </h1>

        {this.props.data.birth.birth_date &&
        <p>
          They were born on {this.props.data.birth.birth_date} at {this.props.data.birth.birth_place} to {this.props.data.birth.birth_mother} and {this.props.data.birth.birth_father}.
          <cite><a href={this.props.data.birth.permalink} title="View reference" target="_blank">ref</a></cite>
        </p>
        }
        {this.props.data.immigration.from_country &&
        <p>
          They immigrated from {this.props.data.immigration.from_country} in {this.props.data.immigration.year}.
          <cite><a href={this.props.data.immigration.permalink} title="View reference" target="_blank">ref</a></cite>
        </p>
        }
        {this.props.data.convict.convict_port &&
        <p>
          They were a convict from {this.props.data.convict.convict_port} and arrived on the ship {this.props.data.convict.convict_ship} in {this.props.data.convict.year}.
          <cite><a href={this.props.data.convict.permalink} title="View reference" target="_blank">ref</a></cite>
        </p>
        }
        {this.props.data.bankruptcy.bankrupt_date &&
        <p>
          They were declared bankrupt on {this.props.data.bankruptcy.bankrupt_date}.
          <cite><a href={this.props.data.bankruptcy.permalink} title="View reference" target="_blank">ref</a></cite>
        </p>
        }
        {this.props.data.marriage.marriage_date &&
        <p>
          They married on {this.props.data.marriage.marriage_date} to {this.props.data.marriage.spouse_name}.
          <cite><a href={this.props.data.convict.permalink} title="View reference" target="_blank">ref</a></cite>
        </p>
        }
        {this.props.data.court.trial_offence &&
        <p>
          They went to court for {this.props.data.court.trial_offence} in {this.props.data.court.year}.
          <cite><a href={this.props.data.marriage.permalink} title="View reference" target="_blank">ref</a></cite>
        </p>
        }
        {this.props.data.census.year &&
        <p>
          They were on the {this.props.data.census.year} census.
          <cite><a href={this.props.data.census.permalink} title="View reference" target="_blank">ref</a></cite>
        </p>
        }
        {this.props.data.inquest.death_verdict &&
        <p>
          {this.props.data.inquest.death_verdict} on {this.props.data.inquest.death_date}.
          <cite><a href={this.props.data.inquest.permalink} title="View reference" target="_blank">ref</a></cite>
        </p>
        }

      </div>
      </div>
    );
  }
}

export default Data;

