import React, { useState } from 'react';
import axios from 'axios';

function Configuration() {
  const [connectionString, setConnectionString] = useState('');

  const handleSubmit = (event) => {
    event.preventDefault();
    axios.post(`${process.env.REACT_APP_API_URL}/change-db`, { connectionString })
      .then(response => {
        alert('Database connection string updated successfully');
      })
      .catch(error => {
        console.error('Error updating database connection string:', error);
      });
  };

  return (
    <div>
      <h1>Configuration</h1>
      <form onSubmit={handleSubmit}>
        <label>
          Database Connection String:
          <input
            type="text"
            value={connectionString}
            onChange={(e) => setConnectionString(e.target.value)}
          />
        </label>
        <button type="submit">Update</button>
      </form>
    </div>
  );
}

export default Configuration;