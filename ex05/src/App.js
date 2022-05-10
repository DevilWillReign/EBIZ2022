import logo from './logo.svg';
import './App.css';
import { BrowserRouter, Route, Switch } from "react-router-dom"
import Basket from './components/Basket';
import Products from './components/Products';

function App() {
  return (
      <BrowserRouter>
        <Switch>
          <Route path='/basket'>
            <Basket />
          </Route>
          <Route path='/payment'>
            <Payments />
          </Route>
          <Route path='/products'>
            <Products />
          </Route>
          <Route path='/'>
            <div className="App">
              <header className="App-header">
                <img src={logo} className="App-logo" alt="logo" />
                <p>
                  Edit <code>src/App.js</code> and save to reload.
                </p>
                <a
                  className="App-link"
                  href="https://reactjs.org"
                  target="_blank"
                  rel="noopener noreferrer"
                >
                  Learn React
                </a>
              </header>
            </div>
          </Route>
        </Switch>
      </BrowserRouter>
  );
}

export default App;
