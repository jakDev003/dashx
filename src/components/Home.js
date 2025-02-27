import React, { useEffect, useState } from 'react';
import axios from 'axios';

function Home() {
  const [records, setRecords] = useState([]);

  useEffect(() => {
    axios.get(`${process.env.REACT_APP_API_URL}/records`)
      .then(response => {
        setRecords(response.data);
      })
      .catch(error => {
        console.error('Error fetching records:', error);
      });
  }, []);

  return (
    <div>
      <h1>Home</h1>
      <div className="tiles">
        {records.map(record => (
          <div key={record.guid} className="tile">
            <h1>{record.data.title}</h1>
            <div className="body">
              {/* Dynamic content based on record data */}
            </div>
            <div className="footer">
              {process.env.REACT_APP_LOGO_TEXT}
            </div>
          </div>
        ))}
      </div>
    </div>
  );
}

export default Home;