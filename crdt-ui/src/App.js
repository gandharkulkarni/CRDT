import './App.css';
import { useEffect, useState } from 'react';
import axios from 'axios';
function App() {
    const [data, setData] = useState('')
    useEffect(() => {
        const interval = setInterval(() => {
            axios.get('http://localhost:9998/state')
              .then(response => {
                  setData(response.data);
              })
              .catch(error => {
                  console.log(error); 
              });
          }, 2000); // 2000 milliseconds interval
          return() => clearInterval(interval)
      }); 
      const handleEdit = (e) =>{
        setData(e.target.value)
        axios.post('http://localhost:9998/state',{Data: e.target.value}, {
            headers: {
              'Content-Type': 'application/json'
            }
      })
        .then(response=>{
            console.log(response)
        }).catch(error=>{
            console.log(error)
        });
      }
  return (
    <div className="App">
      <header className="App-header">
      <textarea
      value={data}
      onChange={(e)=>{handleEdit(e)}}
      className={'editor dark'}
      placeholder="Start typing here..."></textarea>
      </header>
    </div>
  );
}

export default App;
