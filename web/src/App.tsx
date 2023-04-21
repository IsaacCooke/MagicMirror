import { useQuery, gql } from "@apollo/client";
import './App.scss'
import DisplayFlashcards from "./widgets/DisplayFlashcards";
import DisplayReminders from "./widgets/DisplayReminders";

function App() {
  return (
    <div className="App">
      <div className="box">
        <DisplayReminders />
      </div>
      <div className="box">
        <DisplayFlashcards />
      </div>
    </div>
  )
}

export default App;