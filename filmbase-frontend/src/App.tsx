import axios from 'axios';
import { useEffect, useState } from 'react';
import FilmList from './componants/FilmList'
import './App.css'

export interface Film {
  id: number;
  title: string;
  vote_average: number;
  release_date: string;
  overview: string;
  genre_ids: number[];
}

function App() {

  const [filmsList, setFilmsList] = useState<Film[]>([]);
  const [loading, setLoading] = useState<boolean>(true)
  const [error, setError] = useState<Error>()



  useEffect(() => {
    axios
      .get("http://localhost:8080/api/films")
      .then((respose) => {
        setFilmsList(respose.data);
        setLoading(false)
      })
      .catch((err) => {
        console.error("Fetching Data Error: ", err);
        setLoading(false)
        setError(err)
      });
  }, []);

  return (
    <div>
      {loading ? <h1>Loading...</h1> :
      error != null ? <h1>Error getting data from server: {error.message}</h1>:
      <FilmList filmsList={filmsList} title='Films'/>}
    </div>
  )
}

export default App