import { Film } from "../App"

interface FilmCardProps {
  film: Film
}

interface Style {
  width: string
}

const FilmCardStyle: Style = {
  width: "18rem"
} 

function FilmCard({film}: FilmCardProps ) {
  return (
    <div className="card">
      <div className="card-body">
        <h5 className="card-title">{film.title}</h5>
        <h6 className="card-subtitle mb-2 text-body-secondary">Rating - {Math.round(film.vote_average * 10)/10}/10</h6>
        {/* <p className="card-text">{film.genre_ids}</p> */}
        <p className="card-text">{film.release_date}</p>
        <p className="card-text">{film.overview}</p>
      </div>
    </div>
  )
}

export default FilmCard
