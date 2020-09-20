import React, { Component } from "react";
import "./App.css";

class App extends Component {
  state = {
    karma: 0
  };
  componentDidMount() {
    fetch("/api/karma", { credentials: "same-origin" })
      .then(response => response.json())
      .then(json => {
        this.setState({ karma: json.karma });
      })
      .catch(e => console.error(e));
  }
  addKarma = e => {
    e.preventDefault();
    fetch("/api/karma/increment", {
      method: "POST",
      credentials: "same-origin"
    })
      .then(response => response.json())
      .then(json => {
        this.setState({ karma: json.karma });
      })
      .catch(e => console.error(e));
  };
  removeKarma = e => {
    e.preventDefault();
    fetch("/api/karma/decrement", {
      method: "POST",
      credentials: "same-origin"
    })
      .then(response => response.json())
      .then(json => {
        this.setState({ karma: json.karma });
      })
      .catch(e => console.error(e));
  };
  render() {
    return (
      <div className="app">
        <h1>Karma: {this.state.karma}</h1>
        <button onClick={this.addKarma}>
          <img src="https:icon.now.sh/plus/64/fff" alt="Add" />
        </button>
        <button onClick={this.removeKarma}>
          <img src="https:icon.now.sh/minus/64/fff" alt="Remove" />
        </button>
      </div>
    );
  }
}

export default App;
