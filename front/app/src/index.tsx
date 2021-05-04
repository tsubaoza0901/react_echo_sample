import React from 'react';
import ReactDOM from 'react-dom';
import './index.css';
import Login from './components/Login';
import App from './App';
import { Route, BrowserRouter } from 'react-router-dom';
import reportWebVitals from './reportWebVitals';
import { CookiesProvider } from 'react-cookie';

const routing = (
<React.StrictMode>
  <BrowserRouter>
    <CookiesProvider>
      <Route exact path="/" component={Login} />
      <Route exact path="/profiles" component={App} />
    </CookiesProvider>
  </BrowserRouter>
</React.StrictMode>);

ReactDOM.render(routing, document.getElementById('root')
);

// If you want to start measuring performance in your app, pass a function
// to log results (for example: reportWebVitals(console.log))
// or send to an analytics endpoint. Learn more: https://bit.ly/CRA-vitals
reportWebVitals();
