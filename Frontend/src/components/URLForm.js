import React, { useState } from 'react';
import axios from 'axios';
import { toast } from 'react-toastify';
import 'react-toastify/dist/ReactToastify.css';

const URLForm = ({ onAdd }) => {
  const [url, setUrl] = useState('');

  const handleSubmit = async (e) => {
    e.preventDefault();
    try {
      const response = await axios({
        url: 'http://localhost:8080/shorten',
        method: 'post',
        data: JSON.stringify({longUrl:url}),
        headers: {
          'content-type': 'application/json',
          'accept': 'application/json',
        }
      })
      setUrl(JSON.stringify(response.data));
      toast.success('URL shortened successfully!');
    } catch (error) {
      toast.error('Failed to shorten URL');
    }
  };

  return (
    <form onSubmit={handleSubmit}>
      <input
        type="url"
        value={url}
        onChange={(e) => setUrl(e.target.value)}
        placeholder="Enter URL"
        required
      />
      <button type="submit">Shorten</button>
    </form>
  );
};

export default URLForm;
