import './App.css';

import React from 'react';
import MainHeader from './components/MainHeader';
import MainContent from './components/MainContent';


function App() {
  return (
    <div className="app">
      <div className="app-header">
        <MainHeader />
      </div>
      <hr className="app-hr"/>
      <div className="app-body">
        <MainContent />
      </div>
    </div>
  );
}

export default App;
