import React from 'react';
import ReactDOM from 'react-dom';
import './index.css';
import App from './App';
import reportWebVitals from './reportWebVitals';

import {createStore} from "@reduxjs/toolkit"
import reducer from './Components/redux-module/reducer-handlers';
import {Provider} from "react-redux"
import {PersistGate} from 'redux-persist/lib/integration/react'
import {store, persistor} from './Components/redux-module/redux-store'
import LoadingView from './Components/LoadingView'


//const store = createStore(reducer)
ReactDOM.render(
  <Provider store={store}> 
    <PersistGate loading={<LoadingView/>} persistor={persistor}>
      <App/>
    </PersistGate>
    
  </Provider>,
  document.getElementById('root')
);

// If you want to start measuring performance in your app, pass a function
// to log results (for example: reportWebVitals(console.log))
// or send to an analytics endpoint. Learn more: https://bit.ly/CRA-vitals
reportWebVitals();
