import { Dispatch } from "react"

interface alertProps {
    setAlertToFalse: () => void
}

const Alert = ({ setAlertToFalse }: alertProps) => {
    return (
        <div className="alert alert-warning alert-dismissible fade show" role="alert">
            You pressed the button!
            <button type="button" className="btn-close" data-bs-dismiss="alert" aria-label="Close" onClick={setAlertToFalse}></button>
        </div>

    )
}

export default Alert