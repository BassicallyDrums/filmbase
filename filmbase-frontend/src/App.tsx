import { Dispatch, useState } from 'react'
import Alert from './componants/Alert'
import Button from './componants/Button'



function App() {
  const [alertActive, setAlertActive] = useState(false)

  const setAlertToFalse = () => {
    setAlertActive(false)
  }

  return (
    <div>
      {alertActive === true ? <Alert setAlertToFalse={setAlertToFalse} /> : <></>}
      <Button onClick={() => setAlertActive(true)}>Shites About to go down</Button>
    </div>
  )
}

export default App