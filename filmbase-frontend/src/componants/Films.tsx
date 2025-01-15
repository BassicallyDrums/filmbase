import { useEffect, useState } from "react";
import axios from "axios";

interface FilmType {
  id: number;
  title: string;
  rating: number;
}

function Film() {
  const [films, setFilms] = useState<FilmType[]>([]);

  useEffect(() => {
    axios
      .get("http://localhost:8080/api/films")
      .then((respose) => {
        setFilms(respose.data);
      })
      .catch((err) => {
        console.error("Fetching Data Error: ", err);
      });
  }, []);

  return (
    <>
    <h1>Films List</h1>
    <ul>
        {films.map(film => (
            <li key={film.id}>{film.title} - {film.rating}</li>
        ))}
    </ul>
    </>
  )
}

export default Film
