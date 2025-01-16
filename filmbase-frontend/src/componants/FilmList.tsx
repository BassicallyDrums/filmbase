import FilmCard from "./FilmCard";
import { Film } from "../App"

interface FilmListProps {
    filmsList: Film[]
    title: string
}

const FilmList = ({ filmsList, title }: FilmListProps) => {
    return (
        <>
            <h1>{title}</h1>
            <div className="container ">
                <div className="row">
                    {filmsList.map((film, index) => (
                        <div className="col-12 col-sm-6 col-md-4 col-lg-3 mb-4" key={index}>
                            <FilmCard film={film} />
                        </div>
                    ))}
                </div>
            </div>
        </>
    )
}

export default FilmList