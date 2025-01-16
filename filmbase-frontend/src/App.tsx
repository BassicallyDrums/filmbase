import { useState } from 'react'
import Alert from './componants/Alert'
import Button from './componants/Button'



function App() {
  const [alertActive, setAlertActive] = useState(false)

  return (
    <div>
      {alertActive === true ? <Alert setAlert={setAlertActive} /> : <></>}
      <Button onClick={() => setAlertActive(true)}>Shites About to go down</Button>
    </div>
  )
}

export default App