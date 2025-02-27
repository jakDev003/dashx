import React from 'react';
import {
  BrowserRouter as Router,
  Switch,
  Route,
  Link
} from "react-router-dom";
import Home from './components/Home';
import Configuration from './components/Configuration';
import Help from './components/Help';

function App() {
  return (
    <Router>
      <div>
        <nav>
          <div className="logo">
            <img src="/logo/logo.png" alt="Logo" />
            <span>{process.env.REACT_APP_LOGO_TEXT}</span>
          </div>
          <ul>
            <li>
              <Link to="/">Home</Link>
            </li>
            <li>
              <Link to="/configuration">Configuration</Link>
            </li>
            <li>
              <a href="README.md">Help</a>
            </li>
          </ul>
        </nav>

        <Switch>
          <Route path="/configuration">
            <Configuration />
          </Route>
          <Route path="/help">
            <Help />
          </Route>
          <Route path="/">
            <Home />
          </Route>
        </Switch>
      </div>
    </Router>
  );
}

export default App;