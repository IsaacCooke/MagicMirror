import { useQuery, gql } from "@apollo/client";
import './App.scss'
import DisplayReminders from "./widgets/DisplayReminders";

function App() {
  return (
    <div className="App">
      <div className="box">
        <DisplayReminders />
      </div>
    </div>
  )
}

export default App;