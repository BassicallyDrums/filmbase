interface alertProps {
    setAlert: (arg0: boolean) => void
}

const Alert = ({ setAlert }: alertProps) => {
    return (
        <div className="alert alert-warning alert-dismissible fade show" role="alert">
            You pressed the button!
            <button type="button" className="btn-close" data-bs-dismiss="alert" aria-label="Close" onClick={() => setAlert(false)}></button>
        </div>

    )
}

export default Alert