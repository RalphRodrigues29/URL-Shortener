import React from 'react';

const URLList = ({ urls }) => {
  return (
    <div className="url-list">
      {urls.map((url) => (
        <div key={url.shortUrl} className="url-item">
          <a href={'http://localhost:8080/${url.shortUrl}'} target="_blank" rel="noopener noreferrer">
            {url.shortUrl}
          </a>
          <p>{url.longUrl}</p>
        </div>
      ))}
    </div>
  );
};

export default URLList;
