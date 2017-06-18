import React, { Component } from 'react';
import logo from './logo.svg';
import './App.css';
import { getConsentCCVersion } from './utils/Api';
class App extends Component {
  constructor() {
    super();
    console.log('App construction...')
    this.state = { ccversion: {} };
  }

 getVersion() {
    console.log('getVersion...')
    getConsentCCVersion().then((ccversion) => {
      this.setState( ccversion );
      console.log('chaincode version: '+ ccversion.version)
      console.log('chaincode version: '+ this.state.version)
      });
  }

  componentDidMount() {
    this.getVersion();
  }

  render() {

    const ccversion = this.state;

    return (
      <div className="App">
        <div className="App-header">
          <img src={logo} className="App-logo" alt="logo" />
          <h2>Consent chaincode: { ccversion.version }  </h2>
        </div>
        <p className="App-intro">
          To get started, edit <code>src/App.js</code> and save to reload.
        </p>
      </div>
    );
  }
}

export default App;

