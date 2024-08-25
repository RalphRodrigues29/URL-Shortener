import React, { useState } from 'react';
import URLForm from './components/URLForm';
import URLList from './components/URLList';
import { ToastContainer } from 'react-toastify';
import 'react-toastify/dist/ReactToastify.css';
import './App.css';

function App() {
  const [urls, setUrls] = useState([]);

  const handleAddUrl = (newUrl) => {
    setUrls([newUrl, ...urls]);
  };

  return (
    <div className="App">
      <h1>URL Shortener</h1>
      <URLForm onAdd={handleAddUrl} />
      <URLList urls={urls} />
      <ToastContainer />
    </div>
  );
}

export default App;